# 二分法查找
def binary_search(iter, key):
    iter = list(iter.keys())
    low = 0
    high = len(iter) - 1
    time = 0
    while low < high:
        time += 1
        mid = int((low + high) / 2)
        if key < iter[mid]:
            high = mid - 1
        elif key > iter[mid]:
            low = mid + 1
        else:
            # 打印折半的次数
            print("times: %s" % time)
            return mid
    print("times: %s" % time)
    return False