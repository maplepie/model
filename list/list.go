// 一维数组类型
// 用于创建、打印、值获取
package list

import (
	"encoding/json"
	"errors"
)

type List []int

// 创建空一维数组
func Empty() *List {
	return &List{}
}

// 创建带值的一维数组
func New(arr []int) *List {
	return Empty().Set(arr)
}

// 通过字符串解析获取二维数组
func Parse(b []byte) (*List, error) {
	m := Empty()
	err := m.Parse(b)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// 给一维数组设置值
func (l *List) Set(arr []int) *List {
	length := len(arr)
	if length == 0 {
		return &List{}
	}
	res := List(arr)
	return &res
}

// 获取一维数组的值
func (l *List) GetValue(pos int) int {
	_l := *l
	value := []int(_l)
	return value[pos]
}

// 通过字符串解析获取一维数组
func (l *List) Parse(b []byte) error {
	err := json.Unmarshal(b, l)
	if err != nil {
		return errors.New("convert list failed")
	}
	return nil
}

// 打印一维数组
func (l *List) String() string {
	_l := *l
	b, _ := json.Marshal(_l)
	return string(b)
}
