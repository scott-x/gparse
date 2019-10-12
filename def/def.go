/*
* @Author: scottxiong
* @Date:   2019-10-11 21:41:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-10-11 21:44:06
 */
package def

type Table struct {
	Name string
	Fields
}
type Fields struct {
	Key   string
	Value string
}
type Tables []Table
