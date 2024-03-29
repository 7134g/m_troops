import logging
import threading
import time
from py.common.log_modele.Logmanager import get_log

mutex = threading.Lock()
Log = get_log()
Log.setLevel(logging.DEBUG)
# 打工人无工作时候休息时间
WORKER_SLEEP = 5
# 检查工厂状态
CHECK_FACTORY_STATUS = 10
# 打工人数量（并发）
CONCURRENT = 80
# 连续沉睡最大值
SLEEP_CONT_MAX = 10


class Task:
    def __init__(self, fun, *args, **kw):
        self.fun = fun
        self.args = args
        self.kw = kw
        self.name = self._get_task_class_name()

    def run(self):
        if self.args:
            if self.kw:
                self.fun(*self.args, **self.kw)
            else:
                self.fun(*self.args)
        else:
            self.fun()

    def _get_task_class_name(self):
        c = self.fun.__qualname__
        x = c.split(".")
        return x[0]


class Factory:
    def __new__(cls, *args, **kw):
        if not hasattr(cls, '_instance'):
            org = super(Factory, cls)
            cls._instance = org.__new__(cls)
            cls._instance.tasks = []
            cls._instance.workers = []
            cls._instance._active_working = {}
        return cls._instance

    def __init__(self):
        self._max_workers = CONCURRENT
        self.factory_status = True

    # 活动中的计数
    def minus_active(self, name):
        with mutex:
            self._active_working[name] -= 1

    # 活动中的计数
    def plus_active(self, name):
        if name not in self._active_working.keys():
            self._active_working[name] = 1
        self._active_working[name] += 1

    # 生成打工人
    def _create_worker(self, count=0):
        if count:
            worker_count = count
        else:
            worker_count = self._max_workers

        for wid in range(worker_count):
            w = Worker(len(self.workers) + 1, self)
            self.workers.append(w)

    # 开始打工
    def _woring(self):
        for w in self.workers:
            if w.wf == "wating":
                w.start()

    # 获取任务
    def get_task(self) -> (Task, None):
        with mutex:
            try:
                return self.tasks.pop()
            except IndexError:
                return None

    # 添加任务
    def add(self, fun, *args, **kwargs):
        with mutex:
            t = Task(fun, *args, **kwargs)
            active_name = t.name
            if active_name not in self._active_working.keys():
                self._active_working[active_name] = 0
            self.plus_active(active_name)
            self.tasks.append(t)
            return active_name

    # 获取某项任务活动中数量
    def get_active(self, name):
        with mutex:
            return self._active_working[name]

    # 等待某项任务完成
    def wait(self, names):
        for name in names:
            Log.info(f"{name} 等待本次执行完毕, 任务数: {self._active_working[name]}")
            while self.get_active(name) > 0:
                time.sleep(WORKER_SLEEP)
                # if "tester" == name:
                #     pass
            Log.info(f"{name} 结束等待")

    def transfer(self, concurrent):
        with mutex:
            if concurrent < len(self.workers):
                self._max_workers = concurrent
                fired_target, self.workers = self.workers[concurrent:], self.workers[:concurrent]
                for w in fired_target:
                    w.fired()
            else:
                self._create_worker(concurrent - CONCURRENT)
                self._woring()

    # 解雇所有打工人
    def stop(self):
        while self.tasks:
            time.sleep(CHECK_FACTORY_STATUS)

        if self.factory_status:
            self.factory_status = False
            for w in self.workers:
                w.fired()

    # 检查工人是不是全部解雇了
    def shutdown(self):
        for w in self.workers:
            while w.wf != "fired":
                Log.debug(f"正在赶打工人{w.wid}号离开")
                time.sleep(CHECK_FACTORY_STATUS)
            Log.debug(f"打工人{w.wid}号离开")
        Log.info("正常关闭工厂")
        return True

    def start(self):
        Log.debug("线程池启动，招聘打工人")
        self._create_worker()
        Log.debug("打工人准备就绪")
        self._woring()
        # while self.factory_status:
        #     time.sleep(CHECK_FACTORY_STATUS)

        # # 关闭池
        # return self.shutdown()


class Worker(threading.Thread):

    def __init__(self, wid: int, factory: Factory, *args, **kwargs):
        super(Worker, self).__init__(*args, **kwargs)
        self.__running = threading.Event()  # 用于停止线程的标识
        self.__running.set()  # 设置为True
        self.factory = factory
        self.wid = wid
        self.status = True
        self.wf = "wating"
        self.sleep_count = 0

    def run(self):
        self.wf = "doing"
        while self.factory.factory_status and self.status:
            if not self.__running.is_set():
                Log.info("咋瓦鲁多")
            # 获取任务
            t = self.factory.get_task()
            if not t:
                self._sleep(self.wid)
                continue
            try:
                self.sleep_count = 0
                t.run()
                # print(f"打工人{self.wid}号 ，剩余工作 {self.factory._active_working}")
            except Exception as e:
                Log.error(f"打工人工作异常: {e}")
            finally:
                self.factory.minus_active(t.name)
            self.__running.wait()  # 为True时立即返回, 为False时阻塞直到内部的标识位为True后返回
        Log.debug(f"打工人{self.wid}你被炒鱿鱼了")
        self.wf = "fired"

    def fired(self):
        self.status = False

    def _sleep(self, wid):
        self.sleep_count += 1
        Log.debug(f"打工人{wid}号，沉睡第{self.sleep_count}次")
        time.sleep(WORKER_SLEEP)

    def pause(self):
        # 暂停线程，调用wait时会发生阻塞，直到再次调用restart
        self.__running.clear()
        # print(f"stop __running {str(self.__running.is_set())}")

    def restart(self):
        self.__running.set()  # 恢复线程
        # print(f"restart __running {str(self.__running.is_set())}")
