# -*- coding: utf-8 -*-
from io import StringIO

from pdfminer import pslexer
from pdfminer.pdfparser import PDFParser, PDFDocument
from pdfminer.pdfinterp import PDFResourceManager, PDFTextExtractionNotAllowed, PDFPageInterpreter
from pdfminer.converter import TextConverter
from pdfminer.layout import LAParams, LTTextBox
from pdfminer.converter import PDFPageAggregator


class ScrapyPDFParser(PDFParser):
    # 覆盖超类的初始化方法
    def __init__(self, data):
        if isinstance(data, bytes):
            data = data.decode('latin-1')
        self.data = data
        self.lex = pslexer.lexer.clone()
        self.lex.input(data)

        self.context = []
        self.curtype = None
        self.curstack = []
        self.results = []
        self.doc = None
        self.fallback = False


def process_pdf(rsrcmgr, device, data: bytes, pagenos=None, maxpages=0, password='',
                caching=True, check_extractable=True):
    # 创建一个与文件对象相关联的PDF解析器对象。
    parser = ScrapyPDFParser(data)
    # 创建一个存储文档结构的PDF文档对象。
    doc = PDFDocument(caching=caching)
    # 连接解析器和文档对象
    parser.set_document(doc)
    doc.set_parser(parser)

    # 为初始化提供文档密码。
    # #(如果没有设置密码，则给出一个空字符串。)
    doc.initialize(password)
    # 检查文档是否允许文本提取。如果不是,中止。
    if check_extractable and not doc.is_extractable:
        raise PDFTextExtractionNotAllowed('Text extraction is not allowed: %r' % data)
    # 创建一个PDF解释器对象。
    interpreter = PDFPageInterpreter(rsrcmgr, device)
    # 处理文档中包含的每个页面.
    for (pageno, page) in enumerate(doc.get_pages()):
        if pagenos and (pageno not in pagenos): continue
        interpreter.process_page(page)
        if maxpages and maxpages <= pageno + 1: break

# 文本
def readPDF(pdfFile: bytes) -> str:
    rsrcmgr = PDFResourceManager()  # 共享资源的存储库
    retstr = StringIO()  # 内存中读写str
    laparams = LAParams()  # 创建一个聚合器
    device = TextConverter(rsrcmgr, retstr, laparams=laparams)

    process_pdf(rsrcmgr, device, pdfFile)
    device.close()

    content = retstr.getvalue()
    retstr.close()
    return content


# 分行文本
def parse_pdf(body):
    pdf_data_list = []

    # 用文件对象来创建一个pdf文档分析器
    praser = ScrapyPDFParser(body)
    # 创建一个PDF文档
    doc = PDFDocument()
    # 连接分析器 与文档对象
    praser.set_document(doc)
    doc.set_parser(praser)

    # 提供初始化密码
    # 如果没有密码 就创建一个空的字符串
    doc.initialize()

    # 检测文档是否提供txt转换，不提供就忽略
    if not doc.is_extractable:
        raise PDFTextExtractionNotAllowed
    else:
        # 创建PDf 资源管理器 来管理共享资源
        rsrcmgr = PDFResourceManager()
        # 创建一个PDF设备对象
        laparams = LAParams()
        device = PDFPageAggregator(rsrcmgr, laparams=laparams)
        # 创建一个PDF解释器对象
        interpreter = PDFPageInterpreter(rsrcmgr, device)

        # 循环遍历列表，每次处理一个page的内容
        for page in doc.get_pages():
            interpreter.process_page(page)
            # 接受该页面的LTPage对象
            layout = device.get_result()
            # 这里layout是一个LTPage对象，里面存放着这个 page 解析出的各种对象
            # 包括 LTTextBox, LTFigure, LTImage, LTTextBoxHorizontal 等
            for x in layout:
                if isinstance(x, LTTextBox):
                    content = x.get_text().strip().split('\n')
                    pdf_data_list.append(content)
                    # print(content)
                    # print('---------------')
    return pdf_data_list


if __name__ == '__main__':
    import requests

    # 例子
    headers = {
        "Connection": "keep-alive",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
    }
    url = 'http://www.innocom.gov.cn/gxjsqyrdw/dfbhdbdf/201912/5df07f19de434b619cf75cede36d20d2/files/ad9943bad4aa40babd869933ba092ec9.pdf'
    # proxy = "39.108.167.148:7777"
    # proxies = { "http": f"http://{proxy}", "https": f"http://{proxy}", }
    # text = requests.get(url, verify=False, headers=headers, proxies=proxies).text

    response = requests.get(url, verify=False, headers=headers)

    outputString = readPDF(response.content)
    print(outputString)
