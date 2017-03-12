package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
func (t *AppTest) TestThatLoginPageWorks() {
	t.Get("/Login")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
func (t *AppTest) TestThatRegisterPageWorks() {
	t.Get("/Register")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
func (t *AppTest) After() {
	println("Tear down")
}
