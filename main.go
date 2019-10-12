/*
* @Author: scottxiong
* @Date:   2019-10-11 15:11:50
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-10-12 17:32:33
 */
package main

import (
	"fmt"
	"github.com/scott-x/gparse/utils"
)

func main() {
	conetent := utils.GetContent("temp.sql")
	var arr [][]int = utils.GetAllMatchedPosition(conetent, "(")
	for _, v := range arr {
		var f, e int
		for k, x := range v {
			if k == 0 {
				f = x
			}
			if k == 1 {
				e = x
			}

		}
		fmt.Printf("item=>%v\n", conetent[f:e+1])
	}

}
