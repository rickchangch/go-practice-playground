package gorilla

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
)

type gorillaInstance struct{}

var Service = new(gorillaInstance)

func (g gorillaInstance) Run() {

	// hashKey is required, used to authenticate values using HMAC.
	// blockKey is optional, used to encrypt values.
	secureC = securecookie.New([]byte(hashKey), []byte(blockKey))

	// 初始化 Router
	r := mux.NewRouter()

	// Binding endpoint 與對應的 controller method
	// .Method() 限制了可以訪問的 Verbs
	r.HandleFunc("/", index)
	r.HandleFunc("/login", showLogin).Methods("GET")
	r.HandleFunc("/login", doLogin).Methods("POST")
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/register", showRegister).Methods("GET")
	r.HandleFunc("/register", doRegister).Methods("POST")

	// 可以建立特定 prifix 下的 router map
	s := r.PathPrefix("/member").Subrouter()
	s.HandleFunc("", memberIndex)
	s.HandleFunc("/edit", memberShowEdit).Methods("GET")
	s.HandleFunc("/edit", memberDoEdit).Methods("POST")

	// Middleware Funtion，類似 Before Actions，讓 s 下的路由都必須先經過指定 Func
	s.Use(memberAuthHandler)

	// 初始化 CSRF 的 HTTP middleware protection
	CSRF := csrf.Protect(
		[]byte(`1234567890abcdefghijklmnopqrstuvwsyz!@#$%^&*()_+~<>?:{}|,./;'[]\`),
		csrf.RequestHeader("X-ATUH-Token"),
		// hidden input name
		csrf.FieldName("auth_token"),
		csrf.Secure(false),
	)

	log.Fatal(http.ListenAndServe(":8080", CSRF(r)))
}

type member struct {
	Email string `json:"email"`
}

func (m *member) String() string {
	memBytes, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(memBytes)
}

func generateHTML(w http.ResponseWriter, data interface{}, files ...string) {
	var tmp []string
	for _, f := range files {
		tmp = append(tmp, fmt.Sprintf("webUsingGorilla/templates/%s.html", f))
	}

	tmpl := template.Must(template.ParseFiles(tmp...))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func redirect(w http.ResponseWriter, target string) {
	w.Header().Set("Location", target)
	w.WriteHeader(http.StatusFound)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func showLogin(w http.ResponseWriter, r *http.Request) {
	// 產生 csrf token 傳給模板
	generateHTML(w, csrf.TemplateField(r), "layout", "login")
}

func doLogin(w http.ResponseWriter, r *http.Request) {

	// 用 Gorilla schema 處理 crsf 時，記得要加一個 token 欄位，可以不處理
	form := struct {
		Email    string `schema:"email"`
		Password string `schema:"password"`
		Token    string `schema:"auth_token"`
	}{}

	r.ParseForm()

	if err := schema.NewDecoder().Decode(&form, r.PostForm); err != nil {
		log.Println("schema decode:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mem := &member{
		Email: form.Email,
	}

	// 編碼 cookie value
	tmp, err := secureC.Encode(cookieName, mem)
	if err != nil {
		log.Println("encode secure cookie:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 建構 cookie
	cookie := &http.Cookie{
		Name: cookieName,
		// 可以是 String || Struct
		Value:  tmp,
		MaxAge: 0,
		Path:   "/",
	}

	http.SetCookie(w, cookie)
	redirect(w, "/member")
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	redirect(w, "/")
}

func showRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show register")
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "do resgier")
}

func memberIndex(w http.ResponseWriter, r *http.Request) {
	// 直接從 context 中讀取 request-scope 資料
	mem, ok := r.Context().Value(ctxKey(cookieName)).(*member)
	if !ok || mem == nil {
		log.Println(mem, ok)
		redirect(w, "/")
		return
	}

	fmt.Fprintln(w, "member:", mem)
}

func memberShowEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "member show edit")
}

func memberDoEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "member do edit")
}

func memberAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check cookie

		// 讀 cookie
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		value := &member{}
		// 解碼 cookie
		if err := secureC.Decode(cookieName, cookie.Value, value); err != nil {
			log.Println("decode secure cookie:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 轉換成 request-scope 資料 (存進ctx)
		newRequest := r.WithContext(context.WithValue(r.Context(), ctxKey(cookieName), value))

		// Next Handler
		next.ServeHTTP(w, newRequest)
	})
}

type ctxKey string

var (
	secureC *securecookie.SecureCookie
)

const (
	// hashKey: 32 or 64 bytes
	hashKey = "1234567890123456789012345678901234567890123456789012345678901234"
	// blockKey: 16 (AES-128), 24 (AES-192), 32 (AES-256) bytes
	blockKey = "0123456789abcdef"

	cookieName = "mytest"
)
