import json
import os

import pandas as pd


path_root = 'data/sharpe/'


class Sharpe:
    name = ""
    ranking = "" # 排名
    day = ""
    week = ""
    month = ""

    def __init__(self):
        self.__dict__ = {}

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True)


def dump_file(data, path):
    with open(path, "w", encoding="utf-8") as f:
        json.dump(data, f)

def load_file(path, json_data_set):
    with open(path, "r", encoding="utf-8") as f:
        # data = json.load(f, object_hook=Sharpe)
        data = json.load(f)
        for d in data:
            json_data_set.add(json.loads(d, object_hook=Sharpe))
        # print()

def sharpe_top_5(datas):
    for data in datas:
        print(data[1].toJSON())
        # print(f"name: {data[1].name}, ranking: {data[1].ranking}")

def sharpe_bast(kw, path):
    if not os.path.exists(path_root):
        os.mkdir(path_root)

    table_count = {}
    json_data_set = set()
    df = pd.read_csv(path, sep=',', dtype=str)
    day = df["fund_rank_by_sr_daily"].to_list()
    day_v = df["sr_daily"].to_list()
    week = df["fund_rank_by_sr_weekly"].to_list()
    week_v = df["sr_weekly"].to_list()
    month = df["fund_rank_by_sr_monthly"].to_list()
    month_v = df["sr_monthly"].to_list()

    for i, v in enumerate(day):
        s = Sharpe()
        s.name = v
        s.ranking = i
        s.day = day_v[i]
        table_count[v] = s

    for i, v in enumerate(week):
        if v in table_count:
            s = table_count[v]
            s.ranking +=  i
            s.week = week_v[i]
            table_count[v] = s
        else:
            s = Sharpe()
            s.name = v
            s.ranking = i
            s.week = week_v[i]
            table_count[v] = s

    for i, v in enumerate(month):
        if v in table_count:
            s = table_count[v]
            s.ranking += i
            s.month = month_v[i]
            table_count[v] = s
        else:
            s = Sharpe()
            s.name = v
            s.ranking = i
            s.month = month_v[i]
            table_count[v] = s

    table_list =  sorted(table_count.items(), key=lambda x: x[1].ranking)
    # print(path)
    # pprint(table_list[:5])  # 打印前五
    # print("==========================")
    for name, value in table_list:
        json_data_set.add(value.toJSON())
    print(path)
    sharpe_top_5(table_list[:5])
    # pprint(table_list[:5])  # 打印前五
    print("==========================")
    dump_file(list(json_data_set), f"{path_root}{kw}_shape.json")
    # json_data = json.dumps(list(json_data_set))
    print()


def start():
    tab_key_word = [
        {"chose_type": "股票型", "time": "2019-01-01_2020-12-19"},
        {"chose_type": "指数型", "time": "2019-01-01_2020-12-19"},
        {"chose_type": "混合型", "time": "2019-01-01_2020-12-19"},
        {"chose_type": "债券型", "time": "2019-01-01_2020-12-19"},
    ]

    # "data/股票型基金_夏普率排名_2019-01-01_2020-12-19.csv"

    for kw in tab_key_word:
        sharpe_bast(kw['chose_type'], f"data/{kw['chose_type']}基金_夏普率排名_{kw['time']}.csv")




if __name__ == '__main__':
    start()