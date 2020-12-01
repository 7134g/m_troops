import xlrd
from datetime import datetime
from xlrd import xldate_as_tuple
import pymysql
import os
import traceback


def run():
    def get_excel_data(file):
        workbook = xlrd.open_workbook(file)
        sheet = workbook.sheets()[0]  # 读取第一个sheet
        nrows = sheet.nrows  # 行数
        first_row = sheet.row_values(0)
        new_name = ['surnames', 'gender', 'country', 'passport', 'born', 'passport_time', 'issue']
        for index,value in enumerate(first_row):
            first_row[index] = new_name[index]

        # second_row_values = sheet.row_values(1)  # 第二行数据
        num = 1
        for row_num in range(1, nrows):
            row_values = sheet.row_values(row_num)
            dictT = {}
            for i in range(len(row_values)):
                ctype = sheet.cell(num, i).ctype
                # ctype : 0 empty,1 string, 2 number, 3 date, 4 boolean, 5 error
                cell = sheet.cell_value(num, i)
                if ctype == 2 and cell % 1 == 0.0:  # ctype为2且为浮点
                    cell = int(cell)  # 浮点转成整型
                    cell = str(cell)  # 转成整型后再转成字符串，如果想要整型就去掉该行
                # 这个代表表格中是日期类型
                elif ctype == 3:
                    date = datetime(*xldate_as_tuple(cell, 0))
                    cell = date.strftime('%Y/%m/%d')
                elif ctype == 4:
                    cell = True if cell == 1 else False
                cell = repr(cell)
                dictT.update({first_row[i]:cell})
            print(dictT)
            """
            ('LAO/WANG','男','中国','EG888888','1971/01/04','2029/06/14','中国')
            """
            cursor.execute("INSERT INTO aaaaa(surnames, gender, country, passport, born, passport_time, issue) values({0},{1},{2},{3},{4},{5},{6})".format(
                dictT["surnames"],
                dictT["gender"],
                dictT["country"],
                dictT["passport"],
                dictT["born"],
                dictT["passport_time"],
                dictT["issue"],
            ))

            db.commit()
            num = num + 1

    db = pymysql.connect(host='localhost', user='root', password='fxj123', port=3306, db='userinfo', charset='utf8')
    cursor = db.cursor()
    sql = "INSERT INTO aaaaa(surnames, gender, country, passport, born, passport_time, issue) values(%s, %s, %s, %s, %s, %s, %s)"
    files = os.listdir('../data')

    for file in ['1.xlsx']:

        try:
            file_name = '../data/'+file
            print(file_name)
            get_excel_data(file_name)
        except:
            traceback.print_exc()
            db.rollback()
        finally:
            cursor.close()
            db.close()


if __name__ == '__main__':
    run()