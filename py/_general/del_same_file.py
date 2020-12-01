# /usr/bin/env python
# -*- coding:utf-8 -*-
# 删除重复文件
# 运行的代码文件要放到删除重复的文件或图片所包含的目录中
import os
import hashlib
from pprint import pprint


def md5sum(filename):
    with open(filename, 'rb') as f:
        md5 = hashlib.md5()
        while True:
            fb = f.read(8096)
            if not fb:
                break
            md5.update(fb)
        return md5.hexdigest()


def delfile():
    filedir = list(os.walk(os.getcwd()))

    # 生成所有文件路径
    paths = []
    for each in filedir:
        for filename in each[2]:
            if ".idea" in each[0]:
                continue
            path = os.path.join(each[0], filename)
            paths.append(path)

    print('去重前有', len(paths))

    all_md5 = []
    remove_file = []
    # 获取所有md5
    for path in paths:
        md5_value = md5sum(path)
        if md5_value in all_md5:
            remove_file.append(path)
            os.remove(path)
        else:
            all_md5.append(md5_value)

    print('去重前后', len(all_md5))
    print("被删除的文件")
    pprint(remove_file)
    input("按任意键关闭")


if __name__ == '__main__':
    delfile()


