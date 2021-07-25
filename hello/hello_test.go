package hello_test

import (
	"testing"

	"example.com/go-demo-1/hello"
)

func TestHello(t *testing.T) {
	if hello.BaseHello() != "Go gopher" {
		t.Fatal("Worng hello:<")
	}
}
