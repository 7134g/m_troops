# pip install pymupdf
# pip install cnocr


import sys
from io import TextIOWrapper

sys.stdout = TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

import os
import fitz
from cnocr import CnOcr
import warnings

warnings.filterwarnings("ignore")

ocr = CnOcr(model_name='densenet_lite_136-fc')


def pyMuPDF_fitz(pdfPath):
    img_paths = []
    pdf_doc = fitz.open(pdfPath)
    for pg in range(pdf_doc.pageCount):
        page = pdf_doc[pg]
        rotate = int(0)
        # 此处若是不做设置，默认图片大小为：792X612, dpi=72 我扫描的文件是200dpi
        # 每个尺寸的缩放系数为1.3，这将为我们生成分辨率提高2.6的图像。
        zoom_x = 1.3  # (1.33333333-->1056x816)   (2-->1584x1224)
        zoom_y = 1.3
        mat = fitz.Matrix(zoom_x, zoom_y).prerotate(rotate)
        pix = page.get_pixmap(matrix=mat, alpha=False)

        temp_dir, _ = os.path.split(pdfPath)
        name, ext = os.path.splitext(pdfPath)
        image_path = os.path.join(temp_dir, "{}_img_{}.png".format(name, (pg + 1)))
        pix.save(image_path)  # 将图片写入指定的文件夹内
        img_paths.append(image_path)

    return img_paths


def ocr_extract(path: str):
    res = ocr.ocr(path)
    lines = []
    for obj in res:
        line = "".join(obj[0])
        lines.append(line)
    data = "\n".join(lines)
    os.remove(path)
    return data


if __name__ == "__main__":
    # pdf_path = "2.pdf"
    pdf_path = sys.argv[1]
    paths = pyMuPDF_fitz(pdf_path)
    content: str = ""
    for p in paths:
        content += ocr_extract(p)

    print(content)
