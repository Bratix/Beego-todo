package models

type Todo struct {
	Id   int    `form:"-" orm:"pk;auto" `
	Todo string `form:"Todo"`
}
