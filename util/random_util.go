package util

import (
	"math/rand"
	"time"
)

var (
	randomSeed = rand.New(rand.NewSource(time.Now().UnixNano())) // 分配的随机数
)

// RandomRange 范围随机 包下不包上
func RandomRange(min int, max int) int {
	return randomSeed.Intn(max-min) + min
}

// RandomDistribution 随机分城多少份 整数
func RandomDistribution(total int, num int) []uint {
	leftMoney, leftNum := total, num
	moneyList := make([]uint, num)
	for leftMoney > 0 {

		if leftNum == 1 {
			moneyList[num-1] = uint(leftMoney)
			break
		}

		if leftNum == leftMoney {
			for i := 0; i < leftNum; i++ {
				moneyList[i] = 1
			}
			break
		}

		rMoney := int(2 * float64(leftMoney) / float64(leftNum))
		randM := randomSeed.Intn(rMoney)
		moneyList[num-leftNum] = uint(randM)
		leftMoney -= randM
		leftNum--
	}
	return moneyList
}

// IntReflect 映射数据 把x从 in_min到in_max到数据范围映射到out_min到out_max区间内
func IntReflect(x int, in_min int, in_max int, out_min int, out_max int) int {
	return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min
}
