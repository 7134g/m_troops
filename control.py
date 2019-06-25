import importlib
from threading import Thread
import time
import os
import traceback
import gc

import config
import entrance

HOST_IP = str(os.environ.get('HOST_IP', ''))

def run(logger,task_name):
    main_error_count = 0
    max_main_error_count = 3
    while True:
        try:
            importlib.reload(entrance)
            entrance.main(logger,task_name)
            main_error_count = 0
        except Exception as e:
            main_error_count += 1
            if main_error_count >= max_main_error_count:
                main_error_count = max_main_error_count
                time.sleep(10)
            traceback.print_exc()
        gc.collect()


if __name__ == '__main__':
    print('输入模块名称：')
    task_name = input()

    # print('输入需要执行多少条线程数：')
    # config.thread_count = int(input())

    thread_count = config.thread_count
    threads = []

    for num in range(thread_count):
        logger = ''.join([HOST_IP,'_',str(num)])
        t = Thread(target=run, args=(logger,task_name))
        threads.append(t)

    for num in range(thread_count):
        threads[num].start()

    for num in range(thread_count):
        threads[num].join()
