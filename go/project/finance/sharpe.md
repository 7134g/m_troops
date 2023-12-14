#### 变量说明
- fr 基金的预期收益率
- nr 无风险利率 (1年期国债收益率)
- er 平均每日收益
- drs 每日收益
```
				(drs[0] - er) + (drs[1] - er) + ...  
sd 基金的标准差 = ————————————————————————————————————  
							len(drs)  
```
- A 投资组合获得的超额收益 = fr - nr(即承担了风险而获取的超过无风险利率的水平)
- sharpe 每承担单位风险而获得的超额收益的水平(夏普率) = A / sd


#### 例子

	fr = 15% （投资组合预期年回报是15%）
	nr = 3% （无风险年收益水平是3%）
	sd = 6% （标准差是6%）

	A = fr - nr = 12% （超出无风险投资的回报12%）
	sharpe = A / sd = 2 （夏普比率就是2）
	
	代表投资者风险每增长1%，换来的是2%的超额收益(单位：年)



#### new

fundBase
|  基金名字  |  代码  |  发行日期  |  基金经理管理时间  |  更新时间  |  创建时间  |
|  ----  | ----  |  ----  |  ----  |  ----  |  ----  |

fundDaily
|  基金名字  |  代码  |  涨跌幅度  |  更新时间  |  创建时间  |
|  ----  | ----  |  ----  |  ----  |  ----  |  ----  |



采用的算法：
- 每日報酬(%)=(今天資產淨值-昨天資產淨值)/昨天資產淨值
- 夏普率= [(每日報酬率平均值- 無風險利率) / (每日報酬的標準差)]x (252平方根)


> 一年一共252天交易日，其中252平方根是因為一年大約有252天交易日，意思是將波動數值從每日調整成年

1. fund表，以某一基金，每一天的涨跌数据为一条sql
2. 查询fund表，过滤掉所有少于252条数据的数据
3. 过滤管理基金时间少于1年的基金经理


```go
package main

import (
	"fmt"
	"math"
	"time"
)

type fundDaily struct {
	Code       string
	Value      float64
	CreateDate time.Time
	UpdateDate time.Time
}

func main() {
	fds := make([]fundDaily, 0)  // sqlite查询出来的
	var avg, sum, std, sigma, sharp float64
	for _, fd := range fds {
		sum += fd.Value
	}
	avg = sum / float64(len(fds)) // 平均值
	nr := 3.0  //无风险利率

	for _, fd := range fds {
		sigma += math.Pow((fd.Value - avg), 2)
	}
	std = math.Sqrt(sigma / float64(len(fds)))  // 标准差

	sharp = ((avg - nr) / std) * math.Sqrt(252)  // 夏普率
	fmt.Println(sharp)
}

```
