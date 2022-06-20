package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//页面的结构体
type Page struct {
	Title string
	Body  []byte
}

//加载页面的函数
//传入文件名
//这个函数会读取文件名.txt的内容
//并以之前page的格式返回
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

//用来处理请求的hander
func myhander(w http.ResponseWriter, r *http.Request) {
	//截取url中/view/后面的内容
	//作为文件名
	title := r.URL.Path[len("/view/"):]

	//加载对应文件名的内容
	p, _ := loadPage(title)

	//用reponsewriter返回
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

//edithander
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

//savehander
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	//找到query中body的值
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", myhander)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//不难看出 go原生的http库就已经十分完备好用了
//如果上述代码用java或者c实现
//在不引用任何第三方库的情况下
//少说也要几百行代码
//这也不能解释为什么不少大厂的服务器开发都转go了
