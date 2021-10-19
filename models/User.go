package models

type User struct {
	Id       int    `form:"-" orm:"pk;auto" `
	Username string `form:"username"  orm:"unique" `
	Password string `form:"password" `
}
