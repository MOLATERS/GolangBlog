package model

import (
	"database/sql/driver"
	"strings"
)

type Array []string

//从数据库读取变量之后进行处理，从而获得Go类型的变量

func (a *Array) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s),"|")
	*a = ss
	return nil
}

//Value 将数据保存到数据库里面，对数据进行处理，获得数据库支持的类型

func (a Array) Value()(driver.Value,error){
	str := strings.Join(a,"|")
	return str,nil
}