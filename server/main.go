/*

	go get golang.org/x/crypto/bcrypt
	go get github.com/lib/pq
	go get github.com/dgrijalva/jwt-go

	CREATE DATABASE frenzydb
	CREATE TABLE users (uid BIGSERIAL, email TEXT, password BYTEA, primary key (uid, email));
	CREATE TABLE plans (pid BIGSERIAL, uid BIGINT, title TEXT, primary key (pid, uid));

	go run server/main.go

	curl --data "email='Tony'" --data "password='default'" localhost:3030/signup
	curl --data "email='Tony'" --data "password='wrong'" localhost:3030/login
	curl --data "email='Tony'" --data "password='default'" localhost:3030/login

	curl --data "email='Tony'" --data "password='default'" http://localhost:3030/login
	curl http://localhost:3030/login?email=Tony&password=default

*/

package main

import (
	"database/sql"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var jwtSecret []byte = []byte("getfromsecret")

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Claims struct {
	uid int
	jwt.StandardClaims
}

func validateToken(t string) interface{} {
	// validates that token exists and is valid
	if t == "" {
		return nil
	}
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return t, nil
	})
	if err != nil {
		panic(err)
	}
	if !tkn.Valid {
		return nil
	}
	return nil
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres " +
		"password=postgres dbname=frenzydb sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = r.ParseForm()
	if err != nil {
		panic(err)
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	_, err = db.Query(
		"INSERT INTO users (email, password) VALUES ($1, $2)", email, string(hashed),
	)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "new user added")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	psqlInfo := fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres " +
		"password=postgres dbname=frenzydb sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = r.ParseForm()
	if err != nil {
		panic(err)
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	log.Println(email)
	log.Println(password)

	result := db.QueryRow("SELECT password FROM users WHERE email=$1", email)
	if err != nil {
		panic(err)
	}

	var passwordStored string
	result.Scan(&passwordStored)

	err = bcrypt.CompareHashAndPassword([]byte(passwordStored), []byte(password))
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": 0, // TODO: get uid
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{ "token": "` + tokenString + `" }`))
	/*
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: tokenString,
		})
	*/

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// save a plan
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	// load a plan
	fmt.Fprint(w, "user logged in")
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/load", loadHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
