package util

import (
	"fmt"
	"math/big"
	"strconv"
)

func ToBigInt(s string) (*big.Int, error) {

	// 创建一个新的 big.Int 对象
	num := new(big.Int)

	// 将字符串转换为 big.Int
	_, ok := num.SetString(s, 10) // 第二个参数 10 表示字符串是十进制
	if !ok {
		return nil, fmt.Errorf("Data conversion failed :%v ", s)
	}
	return num, nil
}

// CalculatePercentageChange 计算增量百分比（可能为正或负）
func CalculatePercentageChange(oldStr, newStr string) (float64, error) {
	// 将字符串转换为浮点数
	oldValue, err := strconv.ParseFloat(oldStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid old value: %v", err)
	}

	newValue, err := strconv.ParseFloat(newStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid new value: %v", err)
	}

	// 避免旧值为零导致除零错误
	if oldValue == 0 {
		return 0, fmt.Errorf("old value cannot be zero")
	}

	// 计算增量百分比
	change := ((newValue - oldValue) / oldValue) * 100
	return change, nil
}
