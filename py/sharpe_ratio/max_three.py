import json

import pandas as pd


class Sharpe:
    name = ""
    ranking = "" # 排名
    day = ""
    week = ""
    month = ""

    def __init__(self, d):
        self.__dict__ = d

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True)


def dump_file(data, path):
    with open(path, "w", encoding="utf-8") as f:
        json.dump(data, f)

def load_file(path):
    with open(path, "r", encoding="utf-8") as f:
        # data = json.load(f, object_hook=Sharpe)
        data = json.load(f)
        for d in data:
            json_data_set.add(json.loads(d, object_hook=Sharpe))
        # print()


def sharpe_bast():
    df = pd.read_csv("data/股票型基金_夏普率排名_2019-01-01_2020-12-19.csv", sep=',', dtype=str)
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
    for name, value in table_list:
        json_data_set.add(value.toJSON())
    dump_file(list(json_data_set), "data/shape.json")
    # json_data = json.dumps(list(json_data_set))
    print()

if __name__ == '__main__':
    table_count = {}
    json_data_set = set()
    # sharpe_name()
    load_file("data/shape.json")