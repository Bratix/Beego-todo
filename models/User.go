package models

/* User struct with form and orm mapping */
type User struct {
	Id       int     `form:"-" orm:"pk;auto" `
	Username string  `form:"username"  orm:"unique" `
	Password string  `form:"password,password" `
	Todos    []*Todo `orm:"reverse(many)"`
}
