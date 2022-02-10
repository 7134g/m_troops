# pip install pymupdf cnocr pdfminer.six pyinstaller
# pyinstaller -D pdf.spec
from sys import stdout
from io import TextIOWrapper
from warnings import filterwarnings

filterwarnings("ignore")
stdout = TextIOWrapper(stdout.buffer, encoding='utf-8')

from pdfminer.high_level import *
from os import path as os_path
from os import remove as os_remove
from cnocr import CnOcr
from re import sub as re_sub
import fitz

def pdf_to_img(pdfPath):
    img_paths = []
    pdf_doc = fitz.open(pdfPath)
    for pg in range(pdf_doc.pageCount):
        page = pdf_doc[pg]
        rotate = int(0)
        zoom_x = 1.3  # (1.33333333-->1056x816)   (2-->1584x1224)
        zoom_y = 1.3
        mat = fitz.Matrix(zoom_x, zoom_y).prerotate(rotate)
        pix = page.get_pixmap(matrix=mat, alpha=False)

        temp_dir, _ = os_path.split(pdfPath)
        name, ext = os_path.splitext(pdfPath)
        image_path = os_path.join(temp_dir, "{}_img_{}.png".format(name, (pg + 1)))
        pix.save(image_path)
        img_paths.append(image_path)
    pdf_doc.close()
    return img_paths

def extract_img(path: str):
    ocr = CnOcr()
    res = ocr.ocr(path)
    lines = []
    for obj in res:
        line = "".join(obj[0])
        lines.append(line)
    data = "\n".join(lines)
    os_remove(path)
    return data


def extract_pdf(pdf_path):
    text = extract_text(pdf_path)
    text = re_sub("\n+", "\n", text)

    return text


def main():
    pdf_path = sys.argv[1]
    text = extract_pdf(pdf_path)
    if len(text) < 10:
        # pdf 是扫描件，都为图片
        paths = pdf_to_img(pdf_path)
        text: str = ""
        for p in paths:
            text += extract_img(p)

        print(text)
    else:
        print(text)


if __name__ == '__main__':
    import datetime

    n = datetime.datetime.now()
    main()
    print(datetime.datetime.now() - n)
