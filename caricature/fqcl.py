import requests
import os
import shutil
from concurrent.futures import ProcessPoolExecutor,ThreadPoolExecutor
from threading import currentThread
import traceback
from PIL import Image


def save_image(input_name):
    im = Image.open(input_name)
    if im.mode=="RGBA":
        im.load()  # required for png.split()
        background = Image.new("RGB", im.size, (255, 255, 255))
        background.paste(im, mask=im.split()[3])  # 3 is the alpha channel
        im = background
    im.save(input_name.replace('.webp', '.jpg'),'JPEG')
    os.remove(input_name)

def get_task():
    url = 'https://www.manhuatai.com/api/getComicInfoBody?product_id=2&productname=mht&platformname=pc&comic_id=9680'

    response = requests.get(url, verify=False)
    response_json = response.json()
    tasks = response_json['data']['comic_chapter']

    # pool = ThreadPoolExecutor(max_workers=2)
    # for task in tasks:
    #     pool.submit(deal_task,task)
    # pool.shutdown()

    with ProcessPoolExecutor(max_workers=2) as pool:
        for index,task in enumerate(tasks[::-1]):
            pool.submit(deal_task, index, task)



def deal_task(index, task):
    # for index,task in enumerate(tasks[::-1]):
    print(currentThread())
    page_count = task['end_num']
    page_name = task['chapter_name']
    path = "../download/fqcl/{}".format(page_name)
    if os.path.exists(path):
        shutil.rmtree(path)
    os.makedirs(path)
    for count in range(page_count):
        img_url = "https://mhpic.manhualang.com/comic/F/%E9%A3%8E%E8%B5%B7%E8%8B%8D%E5%B2%9A%E6%9D%A1%E6%BC%AB%E7%89%88/{}%E8%AF%9DGQ/{}.jpg-mht.middle.webp".format(index+1,count+1)
        response_img = requests.get(img_url)

        img_path = path + '/{}_{}.webp'.format(page_name,count)
        with open(img_path, 'wb') as f:
            f.write(response_img.content)
            f.close()
        save_image(img_path)

    print(page_name, '已经完成')



if __name__ == '__main__':
    get_task()

