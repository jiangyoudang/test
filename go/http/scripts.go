package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"strings"
	"runtime"
)

const url_base string = "http://www.sbkk8.cn"

func getHTML(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		panic("Fail to get HTML")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	return body

}

func decodeChinese(text []byte) []byte {

	decoder := simplifiedchinese.GBK.NewDecoder()
	decoded_text, err := decoder.Bytes(text)

	if err != nil {
		log.Fatalln(err)
	}

	return decoded_text

}

func getHTMLDecoded(url string) string {
	text := getHTML(url)
	return string(decodeChinese(text))
}

func getTitle(html string) string {

	title_re := `<div class="mingzhuTitle">\s*<h1[^>]*>([^<]*)</h1>`
	title_compile := regexp.MustCompile(title_re)
	title := title_compile.FindStringSubmatch(html)[1]
	return strings.TrimSpace(title)
}

func getChapters(html string) {
	chapters_re := `(?s)<ul class="leftList">(.*?)</ul>`
	chapters_compile := regexp.MustCompile(chapters_re)
	list := chapters_compile.FindStringSubmatch(html)[1]
	fmt.Println(list)
	list_re := `(?s)<li> <a [^>]*?href="([^"]+)"[^>]*>([^<]*)</a> </li>`
	list_compile := regexp.MustCompile(list_re)
	res := list_compile.FindAllStringSubmatch(list, -1)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i][1], res[i][2])
	}

}

func download(urls []string) {

}

func main() {

	// set Go cpu number
	runtime.GOMAXPROCS(runtime.NumCPU())

	//url := "http://www.sbkk8.cn/lizhishu/langchaozhidian/"
	url := "http://www.sbkk8.cn/wangluo/chenxiangwan/"
	html := getHTMLDecoded(url)

	//fmt.Printf("%s", html)

	title := getTitle(html)
	fmt.Printf("%s", title)

	getChapters(html)

}
