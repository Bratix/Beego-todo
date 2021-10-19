package models

/* Todo struct with form and orm mapping */
type Todo struct {
	Id   int    `form:"-" orm:"pk;auto" `
	Todo string `form:"Todo"`
	User *User  `orm:"rel(fk)"`
}
