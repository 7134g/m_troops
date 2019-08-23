import time

# 计时器
def timer(func):
    def wrapper(*args,**kwargs):
        t1 = time.time()
        v = func(*args,**kwargs)
        t2 = time.time()
        z_time = t2-t1
        print("本次操作耗费{}秒".format(z_time))
        return v
    return wrapper