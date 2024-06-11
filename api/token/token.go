package token

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "Secret key for jwt token"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateJWTToken(user_id int, user_role string, username string) *Tokens {

	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["role"] = user_role
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(3 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating access token : ", err)
	}

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_id"] = user_id
	rftclaims["role"] = user_role
	rftclaims["username"] = username
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ValidateToken(tokenStr string) (bool, error) {
	claims, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	if claims["role"] == "admin" {
		return true, nil
	}

	return false, errors.New("forbidden")
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
