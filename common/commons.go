package common

import (
	"strconv"
	"strings"
	"math/rand"
	"fmt"
)

//随机获取的数据
var charArray =  []string{"1", "2", "3", "4", "5", "6", "7", "8","9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "R", "S", "T", "V", "W", "X", "Y"}
///权数
var WEIGHTVALUE =  []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}


//获取vin
func GetRandvin(vins string)string  {

	defer func() {
		re:=recover()
		if re!=nil {
			fmt.Println("GetRandvin 获取vin出错",re)
		}
	}()


	var vin=make([]string,0)
	for i:=0;i<17 ;i++  {
		if i==8 {
			vin=append(vin,"0" )
		}else {
			rm:=RandInt64(1,29)
			l:=charArray[rm]
			vin=append(vin,l )
		}
	}

	s:=strings.Split(vins,"")
	for c:=0;c<len(s) ;c++  {
		vin[c]=s[c]
	}
	i:=0
	th:
	c:= getCheck(vin)
	if c>9 {
		vin[17]=strconv.Itoa(i)
		i++
		goto th

	}
	vin[8]=strconv.Itoa(c)
	return strings.Join(vin,"")
}
//获取验证位数
func getCheck(list []string)int  {
	m:=GetAtoNum()
	count:=0
	for i:=0;i<len(list) ;i++  {
		if i!=8 {
			dy:=m[list[i]]
			count+=WEIGHTVALUE[i]*dy
		}
	}
	return count%11
}

//字母数字对应数字
func GetAtoNum() map[string] int {

	var m = make(map[string]int, 0)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	m["D"] = 4
	m["E"] = 5
	m["F"] = 6
	m["G"] = 7
	m["H"] = 8
	m["J"] = 1
	m["K"] = 2
	m["L"] = 3
	m["M"] = 4
	m["N"] = 5
	m["P"] = 7
	m["R"] = 9
	m["S"] = 2
	m["T"] = 3
	m["U"] = 4
	m["V"] = 5
	m["W"] = 6
	m["X"] = 7
	m["Y"] = 8
	m["Z"] = 9
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9
	m["0"] = 0
	return m
}


//随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}