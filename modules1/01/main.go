package main

import "fmt"

/**
编写一个小程序

给定一个字符串数组
["I","am","stupid","and","weak"]

用 for 循环遍历该数组并修改为
["I","am","smart","and","strong"]
*/
func main () {

	var arr = [5]string{"I","am","stupid","and","weak"}
	for i := 0;i < len(arr); i++ {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		}
		if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}
	fmt.Println(arr)
}
