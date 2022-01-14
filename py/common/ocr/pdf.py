# -*- coding: utf-8 -*-
# pip install pdfminer.six -i https://pypi.doubanio.com/simple

import io
from pdfminer.high_level import *

sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

def return_txt():

    name = sys.argv[1]

    text = extract_text(name)
    print(text)



if __name__ == '__main__':
    return_txt()