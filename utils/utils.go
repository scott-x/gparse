/*
* @Author: scottxiong
* @Date:   2019-10-11 15:17:52
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-10-12 17:50:35
 */
package utils

import (
	"fmt"
	"github.com/scott-x/gutils/fs"
	"strings"
)

func GetContent(path string) string {
	content, err := fs.ReadFile1(path)
	if err != nil {
		panic(err)
	}
	return content
}

//return -1 表示没找到
func GetIndex(str, substr string) int {
	return strings.Index(str, substr)
}

//return -1 表示没找到
func GetLastIndex(str, substr string) int {
	return strings.LastIndex(str, substr)
}

//get all substr's index
func GetAllIndex(str, substr string) []int {
	var result []int
	arr := strings.Split(str, substr)
	if len(arr) == 0 {
		return nil
	}
	var origin int = len(arr[0])
	var position int = origin
	result = append(result, position)
	for k, v := range arr {
		if k == 0 || k == len(arr)-1 {
			continue
		}
		position += len(v) + len(substr)
		result = append(result, position)
	}
	return result
}

func Tab(n int) string {
	return strings.Repeat(" ", n)
}

func FirstLetterToUpperOfOneWord(word string) string {
	return strings.ToUpper(word[0:1]) + strings.ToLower(word[1:])
}

//"I am scott" => [0] : I
func GetWord(str string, index int) string {
	str = strings.Trim(str, " ")
	return strings.Split(str, " ")[index]
}

func GetAllMatchedPosition(str, c string) [][]int {
	var a, b string
	a = c
	switch c {
	case "(":
		b = ")"
	case "{":
		b = "}"
	case "<":
		b = ">"
	case "<<":
		b = ">>"
	case "[":
		b = "]"
	}
	p1 := GetAllIndex(str, a)
	p2 := GetAllIndex(str, b)
	var arr []int
	var result [][]int
	arr = append(arr, p1...)
	arr = append(arr, p2...)

	//冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for k := 0; i < len(arr)-1; k++ {
			var min int
			min = arr[0]
			if min > arr[i] {
				min = arr[i]
			}
		}
	}
	fmt.Println(arr)
	// for i := 0; i < len(p1); i++ {
	// 	var a1, a2 int
	// 	for j := 0; j < 2; j++ {
	// 		if j == 0 {
	// 			a1 = arr[0][i]
	// 		}
	// 		if j == 1 {
	// 			a2 = arr[1][i]
	// 		}
	// 	}
	// 	result = append(result, []int{a1, a2})
	// }
	return result
}

func GetNextMatchedPosition(str string, c rune) string {
	var a, b rune
	a = c
	switch c {
	case 40:
		b = 41
	case 90:
		b = 93
	case 123:
		b = 125

	}
	// fmt.Println(a, b)
	sum := 0
	var p1, p2 int
	for k, v := range []rune(str) {
		if v == a {
			if sum == 0 {
				p1 = k
			}
			sum += 1
		}
		if v == b {
			sum -= 1
			if sum == 0 {
				p2 = k
				fmt.Println(p1, p2)
				break
			}
		}

	}
	return str[p1+1 : p2]
}

//"aas*adsas*dasas"
func GetAllSubStrings(content, regexp string) [][]int {
	if !strings.Contains(regexp, "*") {
		return nil
	}
	var arr [][]int
	args := strings.Split(regexp, "*")
	for _, i := range args {
		fmt.Println(i)
		arr = append(arr, GetAllIndex(content, i))
	}
	return arr
}
