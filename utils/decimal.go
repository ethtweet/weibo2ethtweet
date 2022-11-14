package utils

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

var Decimal = &decimalUtils{}

// BigNumber Utils
type decimalUtils struct{}

// ParseDecimal , 任意数据类型转为 Decimal
func (s *decimalUtils) ParseDecimal(num interface{}) (decimal.Decimal, error) {
	var err error
	var amount decimal.Decimal

	switch v := num.(type) {
	case *big.Int:
		amount = decimal.NewFromBigInt(v, 0)
	case string:
		// 16进制
		if strings.HasPrefix(v, "0x") {
			n := new(big.Int)
			n.SetString(v[2:], 16)
			amount = decimal.NewFromInt(n.Int64())
		} else {
			amount, err = decimal.NewFromString(v)
		}
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromInt(v)
	case int:
		amount = decimal.NewFromInt(int64(v))
	case uint64:
		amount = decimal.NewFromInt(int64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	default:
		err = errors.New(fmt.Sprintf("type not found: %T", v))
	}

	return amount, err
}

// MustParseDecimal , 任意数据类型转为 Decimal
func (s *decimalUtils) MustParseDecimal(num interface{}) decimal.Decimal {
	parseDecimal, err := s.ParseDecimal(num)
	if err != nil {
		panic(err)
	}

	return parseDecimal
}

// ParseUnits 十进制数字转WEI ParseUnits('10.0', 3) => 10000
func (s *decimalUtils) ParseUnits(num interface{}, decimals uint8) (*big.Int, error) {
	amount, err := s.ParseDecimal(num)
	if err != nil {
		return nil, err
	}
	mul := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(decimals)))
	result := amount.Mul(mul)

	return result.BigInt(), nil
}

// FormatUnits WEI转十进制数字 FormatUnits('100.0', 3) => 0.1
func (s *decimalUtils) FormatUnits(num interface{}, decimals uint8) (decimal.Decimal, error) {
	amount, err := s.ParseDecimal(num)
	if err != nil {
		return decimal.Zero, err
	}
	mul := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(decimals)))
	// 除法时精度+1，防止结果被四舍五入
	result := amount.DivRound(mul, int32(decimals+1))
	// 最后返回时再取正确精度
	return result.RoundFloor(int32(decimals)), nil
}

// ParseEther 十进制数字转 Ether
func (s *decimalUtils) ParseEther(num interface{}) (*big.Int, error) {
	return s.ParseUnits(num, 18)
}

// FormatEther Ether 转十进制数字
func (s *decimalUtils) FormatEther(num interface{}) (decimal.Decimal, error) {
	return s.FormatUnits(num, 18)
}

func (s *decimalUtils) BuildCursor(number uint64, index uint) decimal.Decimal {
	// 游标 = 区块高度 + 事件索引 * 0.00001, 如果一个区块事件超过 10000 个就会出BUG
	idx := decimal.NewFromInt(int64(index)).Mul(decimal.NewFromFloat(0.00001))
	// 小数点分割
	idxs := strings.Split(idx.StringFixed(5), ".")
	idxs[0] = strconv.FormatUint(number, 10)
	res, _ := decimal.NewFromString(strings.Join(idxs, "."))

	return res
}
