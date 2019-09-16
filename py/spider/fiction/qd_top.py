#起点网站原创风云榜内容爬取并保存到Excel表格
import requests
from bs4 import BeautifulSoup
import xlsxwriter

url = "https://www.qidian.com/rank/yuepiao?style=1"
#添加头部信息
header = {
      "user-agent":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.221 Safari/537.36 SE 2.X MetaSr 1.0"
  }

#在E盘创建Excel文件
workbook = xlsxwriter.Workbook(r'../download/qd/test.xlsx')
worksheet = workbook.add_worksheet()

#设置文件第一行的信息
data_list = ["小说名","作者名","小说类型","小说状态","小说更新","最新章节地址","简介内容"]
row,col = 0,0
for data in data_list:
  worksheet.write(row, col, data)
  col +=1

row,col = 1,0

#总共25页，从第一页开始爬取
for i in range(1,26):
  page_url = url + "&page=" + str(i)
  res = requests.get(page_url, headers=header)
  soup = BeautifulSoup(res.text, "html.parser")

  for book in soup.select(".book-mid-info"):
   #小说名
    name = book.select("a")[0].text
   #作者名
    author = book.select("a")[1].text
   #小说类型
    _type = book.select("a")[2].text
   #小说状态
    state = book.select("span")[0].text
   #小说更新
    update = book.select("a")[3].text
  #获取时间
    time = book.select("span")[1].text
  #最新一章地址
    href = book.select("a")[3]["href"]
  #简介内容
    content = book.select(".intro")[0].text
    
    expenses = [name, author, _type, state, update, href, time, content]
    for item in expenses:
      worksheet.write(row, col, item)
      col +=1
    row += 1
    col = 0

workbook.close()
