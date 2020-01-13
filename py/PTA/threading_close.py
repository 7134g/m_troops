import threading
import time, datetime


class Job(threading.Thread):

    def __init__(self, *args, **kwargs):
        super(Job, self).__init__(*args, **kwargs)
        self.__running = threading.Event()      # 用于停止线程的标识
        self.__running.set()      # 设置为True

        print(f'__init__ ==》 __running {str(self.__running.is_set())}')

    def run(self):
        while True :
            if self.__running.is_set() == False:
                print("线程阻塞中")
            self.__running.wait()      # 为True时立即返回, 为False时阻塞直到内部的标识位为True后返回
            print(f'working: {datetime.datetime.now()}')
            time.sleep(0.3)

    def stop(self):
        # self.__flag.set()       # 将线程从暂停状态恢复, 如何已经暂停的话
        self.__running.clear()        # 停止线程
        print(f"stop __running {str(self.__running.is_set())}")

    def restart(self):
        self.__running.set()    # 恢复线程
        print(f"restart __running {str(self.__running.is_set())}")


if __name__ == '__main__':
    a = Job()
    print('线程启动')
    a.start()

    time.sleep(1)  # 给予执行时间

    a.stop()
    print('停止工作')
    time.sleep(5)

    # a.join()  # 尝试等待, 发现线程已不工作
    a.restart()

    time.sleep(2)  # 给予执行时间
    a.stop()

