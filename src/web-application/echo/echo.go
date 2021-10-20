package echo

import (
	"context"
	"go-practice-playground/web-application/echo/router"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type echoInstance struct{}

var Service = new(echoInstance)

func (ei *echoInstance) Run() {
	// 初始化 Echo
	e := echo.New()

	// 測試 定義 Router mapping
	defineRouter(e)

	// 測試 使用 echo.Context.Bind() 從 request取值到struct
	valueBinding(e)

	/* 實作 */

	// 初始化 驗證器
	e.Validator = &CustomValidator{validator: validator.New()}

	// 初始化 router
	router.AddRules(e)

	// 於背景執行 Go server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// 建立 channel 來等待 interrupt signal 來 gracefully shutdown
	// 由於 sender 不會 block，所以必須給予足夠的 buffer 以防 deadlock，
	// 而 Notify 會有一個 sinal value，所以必須給予 1 buffer size。
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// 主程式會被 block 在此處，直到 quit recevier 取到值
	<-quit

	// {要關閉的東西}

	// 建立一個 context timeout ，設定 5 秒的時間讓 server 進行關閉，超過就強制關閉所有連線
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Request 取值
// - ctx.QueryParam("key"): 取得 Query String
// - ctx.FormValue("key"): 取得 Body Params
// - ctx.Param("key"): 取得 Path Params; path param 寫法: /users/:key

// Response 回傳類型
// - Context.String: 一般字串 / Content-Type: text/plain; charset=UTF-8
// - Context.JSON: 回傳 JSON / Content-Type: application/json; charset=UTF-8
// - Context.JSONPretty: 同上
// - Context.HTML: 回傳 html 網頁 / Content-Type: text/html; charset=UTF-8
// - Context.Render: 自定義模板 Engine，回傳 html 網頁 / Content-Type: text/html; charset=UTF-8
func defineRouter(e *echo.Echo) {

	// Requet Handler: 直接指定 GET, POST, PUT, DELETE, ANY
	e.GET("/", func(c echo.Context) error {

		// response text
		return c.String(http.StatusOK, "Hello, World!")
	})

	// test: http://localhost:8080/queryString?a="123"
	e.GET("/queryString", func(c echo.Context) error {

		// 使用 QueryParam("key") 去接取 Query String
		return c.String(http.StatusOK, c.QueryParam("a"))
	})

	// test: curl -v -d "name=Joe" -d "email=joe@labstack.com" http://127.0.0.1:8080/formPost
	e.POST("/formPost", func(c echo.Context) error {

		// 使用 FormValue("key")
		return c.String(http.StatusOK, c.FormValue("name")+c.FormValue("email"))
	})

	// test: curl -v http://127.0.0.1:8080/users/7558
	e.GET("/users/:id", func(c echo.Context) error {
		// 使用 Param("param")
		return c.JSONPretty(
			http.StatusOK,
			echo.Map{
				"id":   c.Param("id"),
				"name": "rachel",
			}, "\t")
	})
}

func valueBinding(e *echo.Echo) {
	// Strcut tag
	// - json: 使用 json 格式上傳的資料
	// - param: URL PATH 的參數值
	// - query: Query string 參數值
	// - form: form submit 參數值
	type User struct {
		ID    int    `json:"id" param:"id"`
		Name  string `json:"name" query:"name" form:"name"`
		Email string `json:"email" query:"email" form:"email"`
	}

	// curl --include -X PUT -d "name=abc" -d "email=abc@gmail" http://127.0.0.1:8080/user/1\?name\="aaa"
	e.PUT("/users/:id", func(c echo.Context) error {

		user := new(User)

		// 同時利用 form 與 query string 傳入 name。
		// Echo Bind 的順序會是先 Body (json, xml, form)，然後是 Query String (如果 method 是 GET/DELETE)，最後是 URL PATH。
		// 因此，範例結果 User.Name 為 abc 而非 aaa
		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSONPretty(http.StatusOK, echo.Map{
			"cdoe": "0000",
			"data": user,
		}, "\t")
	})
}
