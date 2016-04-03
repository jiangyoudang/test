package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"strings"
	escapHTML "html"
	"runtime"
)

const (
	url_base = `http://www.sbkk8.cn`
	worker_num = 16
)

type chapter struct {
	id         int
	title, url string
}

func (chap *chapter) Save(book []string, done chan string) {
	book[chap.id] = getContent(chap.url)
	done <- "done!"

}

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

func getBookTitle(html string) string {

	title_re := `<div class="[^>]*?Title">\s*<h1[^>]*>([^<]*)</h1>`
	title_compile := regexp.MustCompile(title_re)
	title := title_compile.FindStringSubmatch(html)[1]
	return strings.TrimSpace(title)
}

func getChapters(html string) []*chapter {
	chapters_re := `(?s)<ul class="leftList[^>]*">(.*?)</ul>`
	chapters_compile := regexp.MustCompile(chapters_re)
	list := chapters_compile.FindStringSubmatch(html)[1]
	list_re := `(?s)<li> <a [^>]*?href="([^"]+)"[^>]*>([^<]*)</a> </li>`
	list_compile := regexp.MustCompile(list_re)
	res := list_compile.FindAllStringSubmatch(list, -1)
	chapters := make([]*chapter, len(res))

	for i, res_match := range res {
		title := res_match[2]
		url := url_base + res_match[1]
		chapters[i] = &chapter{
			id: i,
			title: title, url: url}
	}

	return chapters
}

func cleanParaText(text string) string {
	re := regexp.MustCompile(`<[^<>]*>[^<>]*</[^<>]*>`)
	return re.ReplaceAllString(text, "")
}

func getContent(url string) string {
	// get decoded content
	html := getHTMLDecoded(url)
	content_re := regexp.MustCompile(`(?s)<div [^>]*?id="f_article"[^>]*>(.*)<div [^>]*?class="mingzhuPage">`)
	contents := content_re.FindStringSubmatch(html)[1]
	paragraph_re := regexp.MustCompile(`<p>(.*?)</p>`)
	para_texts := paragraph_re.FindAllStringSubmatch(contents, -1)
	para_list := make([]string, len(para_texts))
	for i, text := range para_texts {
		_text := cleanParaText(text[1])
		para_list[i] = strings.TrimSpace(_text)
	}
	return escapHTML.UnescapeString(strings.Join(para_list, "\n"))
}

func worker(chap *chapter, book []string, done chan string) {
	//fmt.Println(pprof.Lookup("goroutine").Count())
	chap.Save(book, done)
}

func download(url string) (string, string) {

	book_content := getHTMLDecoded(url)
	title := getBookTitle(book_content)
	chapters := getChapters(book_content)
	book := make([]string, len(chapters))
	done := make(chan string)

	for _, chap := range chapters {
		go worker(chap, book, done)
	}

	for i := 0; i < len(chapters); i++ {
		<-done
	}

	book_text := strings.Join(book, "\n\n")
	ioutil.WriteFile(title, []byte(book_text), 0664)
	return  book_text, title
}

func main() {

	/*

	// set Go cpu number
	runtime.GOMAXPROCS(runtime.NumCPU())

	//url := "http://www.sbkk8.cn/lizhishu/langchaozhidian/"
	url := "http://www.sbkk8.cn/wangluo/chenxiangwan/"
	html := getHTMLDecoded(url)

	//fmt.Printf("%s", html)

	title := getTitle(html)
	fmt.Printf("%s", title)

	getChapters(html)

*/

	runtime.GOMAXPROCS(runtime.NumCPU())

	url := `http://www.sbkk8.cn/mingzhu/gudaicn/yijingshuji/qimendunjiamijidaquan/`
	_, title := download(url)
	fmt.Println(title)



}
