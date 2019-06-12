/*

	go get golang.org/x/crypto/bcrypt
	go get github.com/lib/pq

	CREATE DATABASE frenzydb
	CREATE TABLE users (uid BIGSERIAL, email TEXT, password BYTEA, primary key (uid, email));
	CREATE TABLE plans (pid BIGSERIAL, uid BIGINT, title TEXT, primary key (pid, uid));

	curl --data "email='Tony'" --data "password='default'" localhost:8080/signup
	curl --data "email='Tony'" --data "password='wrong'" localhost:8080/login
	curl --data "email='Tony'" --data "password='default'" localhost:8080/login

*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

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

	fmt.Fprint(w, "user logged in")
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
