package main

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

)

//our key for jwt to sign up
var jwtkey = []byte("secret")

//create a simple data
var user = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

type (
	Credentials struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}

	Claims struct {
		Username string `json: "username"`
		jwt.StandardClaims
	}
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SigningString(jwt)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name: "token",
			Value: tokenString,
			ExpiresAt: expirationTime
		}
	)

}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err:= r.Cookie("token")
	if err != nil{
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(
		tokenString,
		claims,
	func(t *jwt.Token) (interface{}, error){
		return jwtkey, nil
	}
	)

	if err != nil{
		err == jwt.ErrSingnatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(ftm.Sprintf("Welcom,%s",claims.Username))) //to user
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	//todo
}
