// 二维数组类型
// 用于创建、打印、值获取
package matrix

import (
	"encoding/json"
	"errors"
)

type Matrix [][]int

// 创建空二维数组
func Empty() *Matrix {
	return &Matrix{}
}

// 创建带值的二维数组
func New(arr [][]int) *Matrix {
	return Empty().Set(arr)
}

// 通过字符串解析获取二维数组
func Parse(b []byte) (*Matrix, error) {
	m := Empty()
	err := m.Parse(b)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// 给二维数组设置值
func (m *Matrix) Set(arr [][]int) *Matrix {
	l := len(arr)
	if l == 0 {
		return &Matrix{}
	}
	res := Matrix{}
	for _, v := range arr {
		res = append(res, v)
	}
	return &res
}

// 获取二维数组的值
func (m *Matrix) GetValue(row, col int) int {
	_m := *m
	value := [][]int(_m)
	return value[row][col]
}

// 通过字符串解析获取二维数组
func (m *Matrix) Parse(b []byte) error {
	err := json.Unmarshal(b, m)
	if err != nil {
		return errors.New("convert matrix failed")
	}
	return nil
}

// 打印二维数组
func (m *Matrix) String() string {
	_m := *m
	b, _ := json.Marshal(_m)
	return string(b)
}
