package users_test

import (
	"go-practice-playground/web-application/echo/users"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestIndex(t *testing.T) {
	e := echo.New()

	// New a request
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	// New a Response Recorder (header, body, code)
	rec := httptest.NewRecorder()
	// New a context instance
	ctx := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, users.Index(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
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
func TestGet(t *testing.T) {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	// New a request
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	// New a Response Recorder (header, body, code)
	rec := httptest.NewRecorder()
	// New a context instance
	ctx := e.NewContext(req, rec)

	ctx.SetParamNames("id")
	// can't set integer????
	ctx.SetParamValues("1")

	// Assertions
	if assert.NoError(t, users.View(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreate(t *testing.T) {
	e := echo.New()

	// Build form values
	f := make(url.Values)
	f.Set("name", "Jon Snow")
	f.Set("email", "jon@labstack.com")

	// New a request
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(f.Encode()))
	// Set header "Content-type: application/x-www-form-urlencoded"
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	// New a Response Recorder (header, body, code)
	rec := httptest.NewRecorder()
	// New a context instance
	ctx := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, users.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestCreatefail(t *testing.T) {
	e := echo.New()

	// New a request
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(userJSON))
	// Set wrong header "Content-type: application/xml" but the correct one is "x-www-form-urlencoded"
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationXML)
	// New a Response Recorder (header, body, code)
	rec := httptest.NewRecorder()
	// New a context instance
	ctx := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, users.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestUpdate(t *testing.T) {
	e := echo.New()

	// Build form values
	f := make(url.Values)
	f.Set("name", "Jon Snow")
	f.Set("email", "jon@labstack.com")

	// New a request
	req := httptest.NewRequest(http.MethodPost, "/api/users/1", strings.NewReader(userJSON))
	// Set header "Content-type: application/json"
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// New a Response Recorder (header, body, code)
	rec := httptest.NewRecorder()
	// New a context instance
	ctx := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, users.Update(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
