import logging
import os
import sys
# from SpiderExcept import *


def get_log(name="log"):
    if sys.platform == "win32" or sys.platform == 'darwin':
        path = os.path.join(os.path.dirname(os.path.realpath(__file__)))
        if not os.path.exists(path):
            os.mkdir(path)
        path = os.path.join(path, "{}.log".format(name))
    else:
        logs_path = '/apps/data/logs/'
        path = os.path.join(logs_path, "{}.log".format(name))
        if not os.path.exists(logs_path):
            os.makedirs(logs_path)

    # logging 对象
    logger = logging.getLogger(__name__)
    logger.setLevel(level=logging.INFO)
    formatter = logging.Formatter('%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s - %(message)s')

    # 日志
    handler = logging.FileHandler(path, encoding="utf-8")
    handler.setLevel(logging.INFO)
    handler.setFormatter(formatter)

    # 窗口
    windown_handler = logging.StreamHandler()
    windown_handler.setFormatter(formatter)

    logger.addHandler(windown_handler)
    logger.addHandler(handler)

    return logger


def write_log(name="log"):
    def wrapper(func):
        def get_log(*args, **kwargs):
            path = os.path.join(Setting.PROJECTFILE, "log", "{}.txt".format(name))

            logger = logging.getLogger(__name__)
            logger.setLevel(level=logging.INFO)
            handler = logging.FileHandler(path, encoding="utf-8")
            handler.setLevel(logging.INFO)
            formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
            handler.setFormatter(formatter)
            logger.addHandler(handler)
            try:
                result = func(*args, **kwargs)
                logger.info(result)
            except Exception as e:
                logger.error(e)
        return get_log
    return wrapper


@write_log("1")
def test():
    print("ok")
    raise Exception("2222")


if __name__ == '__main__':
    log = get_log()
    log.error("fffff")
    log.info("ttttt")
