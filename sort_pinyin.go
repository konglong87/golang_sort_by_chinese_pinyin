package main

import (
	"fmt"
	"bytes"
	"sort"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"golang.org/x/text/transform"
)

type ByPinyin []string

func (s ByPinyin) Len() int      { return len(s) }
func (s ByPinyin) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPinyin) Less(i, j int) bool {
	a, _ := UTF82GBK(s[i])
	b, _ := UTF82GBK(s[j])
	bLen := len(b)
	for idx, chr := range a {
		if idx > bLen-1 {
			return false
		}
		if chr != b[idx] {
			return chr < b[idx]
		}
	}
	return true
}

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

func main() {
	fmt.Println("排序开始=======>")
	b := []string{"最后","哈", "呼", "嚯", "ha", ",","恐龙","的看的","刘","张三","曾哥","12","da","55","---"}

	sort.Strings(b)
	//output: [, ha 呼 哈 嚯]
	fmt.Println("Default sort: ", b)

	sort.Sort(ByPinyin(b))
	//output: [, ha 哈 呼 嚯]
	fmt.Println("By Pinyin sort: ", b)
}
/*
排序开始=======>
Default sort:  [, --- 12 55 da ha 刘 呼 哈 嚯 张三 恐龙 曾哥 最后 的看的]
By Pinyin sort:  [, --- 12 55 da ha 的看的 哈 呼 恐龙 刘 曾哥 张三 最后 嚯]*/
