package main

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
)

var debug bool = false

func main() {
	// args中传入止损比率、利润亏损比率、止盈比率、赢率
	args := os.Args[1:]
	invest_loss, _ := strconv.ParseFloat(args[0], 32)
	profit_loss, _ := strconv.ParseFloat(args[1], 32)
	profit_stop, _ := strconv.ParseFloat(args[2], 32)
	win_ratio, _ := strconv.ParseFloat(args[3], 32)
	monitor_num, _ := strconv.Atoi(args[4])
	win := 0
	total_success := 0
	total_loss := 0
	for i := 0; i < monitor_num; i++ {
		bool, success, loss := invest(invest_loss, profit_loss, profit_stop, win_ratio)
		if bool {
			win++
		}
		total_success += success
		total_loss += loss
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println(win, monitor_num, total_success + total_loss)
	fmt.Println("赢率", float64(win)/float64(monitor_num))
	fmt.Println("平均操作次数", (total_success + total_loss)/monitor_num)
	fmt.Println("平均胜率", float64(total_success)/float64(total_success + total_loss))
}

func invest(invest_loss, profit_loss, profit_stop, win_ratio float64) (win bool, success int, loss int) {
	capital := 9500.0
	profit := 500.0
	capital += profit
	id := 0
	rand.Seed(time.Now().Unix())
	if debug {
		fmt.Println("投资次数	本金	利润	总本金	总利润")
	}
	for {
		id++
		// 本次投资金额为总利润中可以亏损的金额除以止损率,比如总利润500,可以亏50%,则本次投资可以亏250
		// 如果止损率是10%,那么本次可以投资的金额是2500。当然计算得来的本次投资金额不能超过实际可投资的本金。
		current_invest_capital := profit * profit_loss / invest_loss
		if current_invest_capital > capital {
			current_invest_capital = capital
		}
		current_profit := 0.0
		loss_num := rand.Float64()
		// 如果随机数比赢率大表示本次投资失败,以止损结束,比如随机数为0.9,而赢率为0.8,那么止损
		if loss_num > win_ratio {
			current_profit = -current_invest_capital * invest_loss
			loss++
		} else {
			// 否则达到止盈比率就获利了结
			current_profit = current_invest_capital * profit_stop
			success++
		}
		capital += current_profit
		profit += current_profit
		if debug {
			fmt.Printf("%d	%.2f	%.2f	%.2f	%.2f\n", id, current_invest_capital, current_profit, capital, profit)
		}
		if capital > 20000 {
			win = true
			if debug {
				fmt.Println("成功", success)
				fmt.Println("失败", loss)
			}
			break
		}
		if profit < 1 {
			win = false
			if debug {
				fmt.Println("成功", success)
				fmt.Println("失败", loss)
			}
			break
		}
	}
	return
}

