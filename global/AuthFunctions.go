package global

import (
	"fmt"
	"net/http"
	"todoapp/models"
)

/* Authenticate user */
func Authenticate(userId int, td *models.TokenDetails, w http.ResponseWriter) {
	fmt.Println("Creating tokens!!!!!!!")
	cookie := CreateCookieWithJWT("AccessToken", td.AccessToken)
	http.SetCookie(w, cookie)
	cookie = CreateCookieWithJWT("RefreshToken", td.RefreshToken)
	http.SetCookie(w, cookie)
	CreateAuth(userId, td)
}

/* Log user out */
func Logout(key1, key2 string, w http.ResponseWriter) {
	DeleteCookieWithJWT("AccessToken", w)
	DeleteCookieWithJWT("RefreshToken", w)
	DeleteAuth(key1, key2)
}

/* Creating a http cookie with passed in token as value */
func CreateCookieWithJWT(name, token string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Secure:   true,
		HttpOnly: true, // change to true for production
	}
	return cookie
}

/* Same as create cookie, but it sets the maxage atr to -1, and deleting it in the process*/
func DeleteCookieWithJWT(name string, w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Secure:   true,
		HttpOnly: true, // change to true for production
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}
