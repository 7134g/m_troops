import time
import os
import logging
import traceback


# 计时器
def timer(func):
    def wrapper(*args, **kwargs):
        t1 = time.time()
        v = func(*args, **kwargs)
        t2 = time.time()
        z_time = t2-t1
        print("本次操作耗费{}秒".format(z_time))
        return v
    return wrapper


# 日志编写器
def write_log(name):
    def wrapper(func):
        def get_log(*args, **kwargs):
            logger = logging.getLogger(__name__)
            logger.setLevel(level=logging.INFO)
            path = os.path.split(os.path.realpath(__file__))[0]
            handler = logging.FileHandler(path + '\\log\\' + name + ".txt", encoding="utf-8")
            handler.setLevel(logging.INFO)
            formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
            handler.setFormatter(formatter)
            logger.addHandler(handler)
            try:
                result = func(*args, **kwargs)
            except:
                logger.error(traceback.format_exc())
            logger.info(result)
        return get_log
    return wrapper


# 日志编写器2
class LoggerManage:
    def __init__(self, name):
        self.name = name

    def __call__(self, func):
        def wrapper(*args, **kwargs):
            logger = logging.getLogger(__name__)
            logger.setLevel(level=logging.INFO)
            path = os.path.split(os.path.realpath(__file__))[0]
            handler = logging.FileHandler(path + '\\log\\' + self.name + ".txt", encoding="utf-8")
            handler.setLevel(logging.INFO)
            formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
            handler.setFormatter(formatter)
            logger.addHandler(handler)
            try:
                result = func(*args, **kwargs)
            except:
                logger.error(traceback.format_exc())
            logger.info(result)
        return wrapper


if __name__ == '__main__':
    @LoggerManage("error")
    def test():
        print("do something ...")
        return "gogogo_finish"

    test()
