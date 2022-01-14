from cnocr import CnOcr

ocr = CnOcr()

def ocr_extract(path: str):
    res = ocr.ocr(path)
    lines = []
    for obj in res:
        line = "".join(obj[0])
        lines.append(line)
    content = "\n".join(lines)
    return content


if __name__ == '__main__':
    ocr_extract('./images/images_1.png')
    ocr_extract('./images/images_2.png')


