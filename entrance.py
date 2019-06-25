import importlib
import traceback
import gc
import config



def main(logger,task_name):
    try:
        task_x = task_name.split('.')
        task_fl = task_x[0]
        task_nz = task_x[1]

        if task_fl in config.MODELENAME.keys():
            pass
        else:
            raise Exception('目录名错误')

        if task_nz in config.MODELENAME[task_fl].keys():
            print(''.join(['模块名称输入正确，开始执行:',task_name]))
        else:
            raise Exception('爬虫名错误')

        module_name = importlib.import_module('.', task_name)
        # 更新内存中的爬虫文件
        module_name = importlib.reload(module_name)
        # 开始执行
        module_name.main()

    except:
        print(logger,'，发生错误执行失败')
        traceback.print_exc()

    # 释放内存
    gc.collect()

if __name__ == '__main__':
    logger = '测试操作'
    task_name = 'mh.dpcq'
    main(logger,task_name)