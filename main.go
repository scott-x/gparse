/*
* @Author: scottxiong
* @Date:   2019-10-11 15:11:50
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-10-11 17:56:12
 */
package main

import (
	"fmt"
	"github.com/scott-x/gparse/utils"
)

func main() {
	fmt.Println("hello world")
	c := utils.GetContent("a.txt")
	fmt.Println(utils.GetAllIndex(c, "wo"))
	fmt.Println(utils.GetWord("hellO i love", 1))
}
