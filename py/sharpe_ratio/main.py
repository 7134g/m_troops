import fund_select_main
import max_three
import fund_sharpe_ratio

def all():
    fund_select_main.new_earnings_sharpe()
    max_three.sharpe_start() # 显示排名前三


def single():
    sharpe_r_list = []
    fund_code = "002190"
    begin_date = "2019-01-01"
    chose_type = "hh"
    fund_sharpe_ratio.task(sharpe_r_list, fund_code, begin_date, chose_type)
    print(sharpe_r_list)


if __name__ == '__main__':
    all()
    # single()