package tests

import "github.com/revel/revel"

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
}
