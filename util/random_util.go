package util

import (
	"math/rand"
	"time"
)

var (
	random_seed = rand.New(rand.NewSource(time.Now().UnixNano())) // 分配的随机数
)

// 范围随机 包下不包上
func RandomRange(min int, max int) int {
	return random_seed.Intn(max-min) + min
}

// 随机分城多少份 整数
func RandomDistribution(total int, num int) []uint {
	left_money, left_num := total, num
	money_list := make([]uint, num)
	for left_money > 0 {

		if left_num == 1 {
			money_list[num-1] = uint(left_money)
			break
		}

		if left_num == left_money {
			for i := 0; i < left_num; i++ {
				money_list[i] = 1
			}
			break
		}

		rMoney := int(2 * float64(left_money) / float64(left_num))
		rand_m := random_seed.Intn(rMoney)
		money_list[num-left_num] = uint(rand_m)
		left_money -= rand_m
		left_num--
	}
	return money_list
}
