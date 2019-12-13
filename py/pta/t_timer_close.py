from threading import Thread, enumerate
import importlib
import time
import datetime
from process import Setting
import traceback
import random
import gc


class Console:
    def __init__(self):
        self._running = True

    def terminate(self):
        self._running = False

    def run(self, modele_name):
        while True:
            start_time = datetime.datetime.now()
            Setting.LOGGER.info("程序启动 此时时间：{}".format(start_time))

            try:
                module_name = importlib.import_module('.', modele_name)
                task = importlib.reload(module_name)
                task.Setting = Setting
                task.main()
                del task
                gc.collect()

            except Exception:
                traceback.print_exc()
                Setting.LOGGER.error("程序意外停止了, 睡眠100秒: {}".format(traceback.format_exc()))
                time.sleep(100)
                gc.collect()

            finally:
                end_time = datetime.datetime.now()
                Setting.LOGGER.info("程序结束 运行时间: ({}) - ({})".format(start_time, end_time))
                time.sleep(random.randint(1800, 3600))


def main():
    status = True
    while True:
        if status:
            t_list:list = []
            ctrl = Console()
            t1 = Thread(target=ctrl.run, args=("media.qyyjt.spider",))
            t2 = Thread(target=ctrl.run, args=("xxx.xxx.xxx",))
            t_list.append(t1)
            t_list.append(t2)
            for i, t in enumerate(t_list):
                try:
                    t.start()
                except:
                    Setting.LOGGER.err("启动失败，t{}出问题了。。。".format(i))

            status = False

        if datetime.datetime.now() > datetime.datetime.now().replace(hour=23, minute=0, second=0) or \
                datetime.datetime.now() < datetime.datetime.now().replace(hour=7, minute=0, second=0):
            ctrl.terminate()
            print("此时线程数: {}".format(len(enumerate())))
            status = True
            gc.collect()

        time.sleep(600)


if __name__ == '__main__':
    # "media.qyyjt.spider", "media.baidu.baidu"
    main()



