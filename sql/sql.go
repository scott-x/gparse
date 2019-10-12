/*
* @Author: scottxiong
* @Date:   2019-10-11 21:41:19
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-10-11 23:19:49
 */
package sql

import (
	"github.com/scott-x/gparse/def"
	"github.com/scott-x/gparse/utils"
	"strings"
)

func GetTablesInfo(path string) string {
	table := &def.Table{}
	// tables := &def.Tbales{}
	content := utils.GetContent(path)
	arr := []string{"CREATE TABLE", "(", ")"}
	index1 := strings.Index(content, arr[0])
	index2 := utils.GetAllIndex()
	for {
		if index2 < index1 {

		}
	}
	return "hello"
}
