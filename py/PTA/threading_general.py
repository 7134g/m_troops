import random
import threading
import time

from concurrent.futures import wait
from concurrent.futures.thread import ThreadPoolExecutor


# 线程池, 非阻塞
class MThreadPool:
    def __init__(self):
        self.db = []

    def start(self, params, params2):
        pass

    def run(self):
        count = 0  # 计数, 控制线程池的任务数量
        max_workers = 16# 最多同时操作16个
        with ThreadPoolExecutor(max_workers=max_workers) as thread_pool:
            tasks = []

            while True:
                for each in self.db:
                    params = each
                    params2 = each
                    futrue = thread_pool.submit(self.start, params, params2)

                    count += 1
                    tasks.append(futrue)
                    if count % 100 == 0:
                        print(f"此时操作了 {count}", flush=True)
                        wait(tasks)  # 等待任务完成
                        tasks = []

# 直接开新线程, 会阻塞
class MThread:
    def start(self, params, params2):
        time.sleep(params)
        print(threading.current_thread(), params, params2)

    def run(self):
        ts = []

        # 构建线程
        temp = list(range(5))
        random.shuffle(temp)
        for i in temp:
            params, params2 = i, i + 100
            ts.append(threading.Thread(target=self.start, args=(params, params2)))

        # 线程启动
        for t in ts:
            t.start()

        # 等待所有完成
        for t in ts:
            t.join()


if __name__ == '__main__':
    q = MThread()
    q.run()