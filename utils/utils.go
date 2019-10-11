/*
* @Author: scottxiong
* @Date:   2019-10-11 15:17:52
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-10-11 17:58:38
 */
package utils

import (
	_ "fmt"
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
		if k == 0 {
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
