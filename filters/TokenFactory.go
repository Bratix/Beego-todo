package filters

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

/* Creating a jwt token */
func CreateToken(userid int) (string, error) {
	var err error
	/* Getting the env variable specified in /conf/auth.evn */
	env := EnviromentVariable("AUTH_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(env))
	if err != nil {
		return "", err
	}
	return token, nil
}

/* Creating a http cookie with passed in token as value */
func CreateCookieWithJWT(token string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "JWTCookie",
		Value:    token,
		Secure:   true,
		HttpOnly: false, // change to true for production
	}
	return cookie
}

/* Same as create cookie, but it sets the maxage atr to -1, and deleting it in the process*/
func DeleteCookieWithJWT(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "JWTCookie",
		Value:    "",
		Secure:   true,
		HttpOnly: false, // change to true for production
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}

/* Extracting the token from JWTCookie */
func ExtractTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("JWTCookie")
	if err != nil {
		return "", err
	}
	token := cookie.Value
	return token, nil
}

/* Verify token */
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	/* First we need to extract token */
	tokenString, err := ExtractTokenFromCookie(r)
	/* Check to see if there was an error */
	if err != nil {
		fmt.Println("Error extracting token from cookie", err)
		return nil, err
	}

	/* The token (which is string) is getting parse with jwt.Parse and the signing method is checked */
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(EnviromentVariable("AUTH_SECRET")), nil
	})

	/* Check for error */
	if err != nil {
		return nil, err
	}

	/* Return the token */
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (int, error) {
	/* Verify token */
	token, err := VerifyToken(r)

	/* Check if there was an error and if the token is valid */
	if err != nil || !token.Valid {
		return 0, err
	}
	/* Extract user id from the token */
	claims, _ := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	return uid, nil
}

func EnviromentVariable(key string) string {

	/* Viper function to read /conf/auth.evn data */

	// name of config file (without extension)
	viper.SetConfigName("auth")
	// look for config in the working directory
	viper.AddConfigPath("./conf")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
