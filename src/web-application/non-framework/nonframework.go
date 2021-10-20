package nonframework

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
	Web application 檔案結構
	.
	├── main.go
	├── public
	│   └── db.png
	└── templates
		├── layout.html
		├── nav.html
		└── test.html

	- 特色
		- 內建 HTTP server
		- 有提供前端模板語言

*/

type nonFramework struct{}

var Service = new(nonFramework)

func (n nonFramework) Run() {
	// 實作 routing 機制
	mux := http.NewServeMux()

	// 建立 靜態資料 路徑
	files := http.FileServer(http.Dir("./webApplication/public"))
	// 設定 路由為 /static/ 會對應至 /webApplication/public/ 底下的 substrees
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// 指定 function routing
	mux.HandleFunc("/", test)
	mux.HandleFunc("/set_cookie", setCookie)
	mux.HandleFunc("/get_cookie", getCookie)
	mux.HandleFunc("/nueip", toGoogle)
	mux.HandleFunc("/handle", handlerRequest)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// MyData ...
type MyData struct {
	Title string
	Nav   string
	Data  interface{}
}

// func generateHTML(w http.ResponseWriter, data interface{}, files ...string) {
// 	var tmp []string
// 	for _, f := range files {
// 		tmp = append(tmp, fmt.Sprintf("webApplication/templates/%s.html", f))
// 	}

// 	// 確認模板路徑與檔案是否正確，是的話取得該模板物件
// 	tmpl := template.Must(template.ParseFiles(tmp...))
// 	// 執行模板，並將 data 帶入其中，"layout"被定義在版型內，指定要從哪個區塊開始
// 	tmpl.ExecuteTemplate(w, "layout", data)
// }

func test(w http.ResponseWriter, r *http.Request) {
	data := &MyData{
		Title: "測試",
		Nav:   "test",
	}

	data.Data = struct {
		TestString   string
		SimpleString string
		TestStruct   struct{ A, B string }
		TestArray    []string
		TestMap      map[string]string
		Num1, Num2   int
		EmptyArray   []string
		ZeroInt      int
	}{
		`O'Reilly: How are <i>you</i>?`,
		"中文測試",
		struct{ A, B string }{"foo", "boo"},
		[]string{"Hello", "World", "Test"},
		map[string]string{"A": "B", "abc": "DEF"},
		10,
		101,
		[]string{},
		0,
	}

	tmpl, err := template.ParseFiles("webApplication/templates/layout.html", "webApplication/templates/nav.html", "webApplication/templates/test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	w.WriteHeader(http.StatusOK)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	allCookie := r.Cookies()
	firstCookie, err := r.Cookie("first_cookie")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Fprintln(w, "first cookie not found")
		} else {
			fmt.Fprintln(w, "get first cookie failure")
		}
	} else {
		fmt.Fprintln(w, "first cookie:", firstCookie)
	}

	fmt.Fprintln(w, allCookie)
}

func toGoogle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.google.com/")
	w.WriteHeader(http.StatusFound)
}

func handlerRequest(w http.ResponseWriter, r *http.Request) {

	// 取得 HTTP Verbs
	requestMethod := r.Method

	// Go 有內建多種取 request 值的方式
	form := r.Form
	postForm := r.PostForm
	multipartForm := r.MultipartForm
	// formValue := r.FormValue
	// PostFormValue := r.PostFormValue

	fmt.Fprintln(w, requestMethod)
	fmt.Fprintln(w, form)
	fmt.Fprintln(w, postForm)
	fmt.Fprintln(w, multipartForm)

	// // 若有其他 Header 改動，皆必須在 WriterHeader 之前；Cookie 也是，因為包在 header 裡
	// w.Header().set("location", "https://www.google.com")
	// http.SetCookie(w, &c1)

	// 修改 HTTP status code，必須放在 Header 更改的最後
	w.WriteHeader(http.StatusOK)
}
