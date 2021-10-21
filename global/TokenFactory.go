package global

import (
	"fmt"
	"net/http"
	"time"
	"todoapp/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

/* Creating a jwt token */
func CreateToken(userid int) (*models.TokenDetails, error) {

	td := &models.TokenDetails{}
	td.AccessUuid = uuid.NewV4().String()
	td.AtExpires = time.Now().Add(time.Minute * 10).Unix()

	td.RefreshUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 1).Unix()

	var err error
	/* Getting the env variable specified in /conf/auth.evn */
	env := EnviromentVariable("AUTH_SECRET")

	/* Creating access token */
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(env))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["access_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(env))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func RefreshToken(w http.ResponseWriter, r *http.Request) error {

	refreshTokenString, err := ExtractTokenMetadata("RefreshToken", r)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = DeleteRefreshToken(refreshTokenString.Uuid)

	if err != nil {
		fmt.Println("Error deleting refresh token from redis!")
		return err
	}

	td, err := CreateToken(refreshTokenString.UserId)
	if err != nil {
		fmt.Println("Error creating token!")
		return err
	}

	Authenticate(refreshTokenString.UserId, td, w)
	return nil

}

/* Extracting the token from JWTCookie */
func ExtractTokenFromCookie(tokenName string, r *http.Request) (string, error) {
	cookie, err := r.Cookie(tokenName)
	if err != nil {
		return "", err
	}
	token := cookie.Value
	return token, nil
}

/* Verify token */
func VerifyToken(tokenName string, r *http.Request) (*jwt.Token, error) {
	/* First we need to extract token */
	tokenString, err := ExtractTokenFromCookie(tokenName, r)
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

func ExtractTokenMetadata(tokenName string, r *http.Request) (*models.ExtractedTokenData, error) {
	/* Verify token */
	token, err := VerifyToken(tokenName, r)

	/* Check if there was an error and if the token is valid */
	if err != nil || !token.Valid {
		return nil, err
	}
	/* Extract user id from the token */
	claims, _ := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	accessUuid := claims["access_uuid"].(string)

	var accessDetails = &models.ExtractedTokenData{
		Uuid:   accessUuid,
		UserId: uid,
	}

	return accessDetails, nil
}
