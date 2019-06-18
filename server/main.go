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
	"strconv"
	"time"
)

var jwtSecret []byte = []byte("getfromsecret")
var psqlInfo = fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres " +
	"password=postgres dbname=frenzydb sslmode=disable")
var db, err = sql.Open("postgres", psqlInfo)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Claims struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-- signup")
	enableCors(&w)
	r.ParseForm()
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	_, err = db.Query(
		"INSERT INTO users (email, password) VALUES ($1, $2)", email, string(hashed),
	)
	if err != nil {
		panic(err)
	}

	result := db.QueryRow("SELECT uid, password FROM users WHERE email=$1", email)
	var (
		uid            int
		passwordStored string
	)
	err = result.Scan(&uid, &passwordStored)

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	log.Println(tokenString)

	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"tokenString": "` + tokenString + `", "uid": ` + strconv.Itoa(uid) + `}`))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-- login")
	enableCors(&w)
	r.ParseForm()
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	result := db.QueryRow("SELECT uid, password FROM users WHERE email=$1", email)
	var (
		uid            int
		passwordStored string
	)
	err := result.Scan(&uid, &passwordStored)

	if err != nil { // email not found in database
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordStored), []byte(password))
	if err != nil {
		panic(err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"tokenString": "` + tokenString + `", "uid": ` + strconv.Itoa(uid) + `}`))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-- refresh")
	enableCors(&w)
	r.ParseForm()
	tokenString := r.PostFormValue("tokenString")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if !token.Valid {
		panic("-- token not valid")
	}
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"uid": ` + strconv.Itoa(claims.UID) + `}`))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// save a plan
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	// load a plan
	enableCors(&w)
	r.ParseForm()
	plan := r.PostFormValue("plan")
	tokenString := r.PostFormValue("tokenString")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tkn.Valid && err == nil {
		log.Println(claims.UID)
		log.Println(plan)
	}
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/refresh", refreshHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/load", loadHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
