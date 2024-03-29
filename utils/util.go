package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"time"
)

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// 将传入的时间戳转换为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

/**
 * 将文章详情的内容，转换成HTML语句
 */
func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	// 获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/*
		对document进行查询，选择器和css的语法一样
		第一个参数：i是查询到的第几个元素
		第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
