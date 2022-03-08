/**
* @Author: Chao
* @Date: 2022/3/4 13:24
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "100[leetcode]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	ss := make([]string, 0)

	//数字字符串,为了连续位数字
	dig := ""
	for _, i := range s {
		//判断是否为数字
		if i >= '0' && i <= '9' {
			//拼接
			dig = dig+string(i)
			continue
		}

		//如果数字字符串不为空,整体写入栈,并重置
		if dig != "" {
			ss = append(ss, dig)
			dig = ""
		}

		//判断是否为]
		if i == ']' {
			//构建临时字符串
			str := ""
			//弹出字符,直到[
			for len(ss) > 0 && ss[len(ss)-1] != "[" {
				str = ss[len(ss)-1] + str
				ss = ss[:len(ss)-1]
			}
			//弹出[
			ss = ss[:len(ss)-1]

			//判断是否还有数据,有数据的话需要判断下一个弹出的是不是数字
			if len(ss) > 0 {
				//弹出一个
				n := ss[len(ss)-1]
				//判断是否为数字
				num, ok := isDigital(n)
				if ok {
					//构建单串,循环拼接
					tmp := str
					for i := 1; i < num; i++ {
						str += tmp
					}
				}
				//弹出数字
				ss = ss[:len(ss)-1]
			}
			//将拼接好的字符串写入栈
			ss = append(ss, str)
			continue
		}

		//[或者字母直接入栈
		ss = append(ss, string(i))
	}
	//将栈元素组合为字符串返回
	return strings.Join(ss, "")
}

func isDigital(s string) (int, bool) {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return atoi, true
}