import requests
import bs4
import openpyxl
import re


def open_url(url):
    headers = {'User-Agent':'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36'}
    html = requests.get(url, headers = headers)
    return html

def find_data(html):
     content = bs4.BeautifulSoup(html.text,'html.parser')
     paragraph = content.find(id="Cnt-Main-Article-QQ")
     target = paragraph.find_all('p',style="TEXT-INDENT: 2em")
     for each in target:
         print(target)
     target = iter(target)

     data = []

     for each in target:
         if each.text.isnumeric():
             data.append([
                re.search(r'\[(.+)\]',next(target).text).group(1),
                re.search(r'\d.*',next(target).text).group(0),
                re.search(r'\d.*',next(target).text).group(0),
                re.search(r'\d.*',next(target).text).group(0)
                 ])
     return data
             
def save_excel(data):
    wb = openpyxl.Workbook()
    wb.guess_types = True
    ws = wb.active
    ws.append(['城市','平均房价','平均工资','房价工资比'])
    for each in data:
        ws.append(each)

    ws.column_dimensions['B'].width = 15
    ws.column_dimensions['C'].width = 15
    ws.freeze_panes = 'A2'
    
    wb.save('2017年中国主要城市房价工资比排行榜.xlsx')
    

def main():
    url = 'http://news.house.qq.com/a/20170702/003985.htm'
    html = open_url(url)

    data = find_data(html)

    save_excel(data)

if __name__ == '__main__':
    main()
    

