/*

	go get golang.org/x/crypto/bcrypt
	go get github.com/lib/pq
	go get github.com/dgrijalva/jwt-go

	CREATE DATABASE frenzydb
	CREATE TABLE users (uid BIGSERIAL, email TEXT, password BYTEA, primary key (uid, email));
	CREATE TABLE plans (pid BIGSERIAL, uid BIGINT, title TEXT, itemA TEXT, primary key (pid, uid));

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
	"encoding/json"
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

func jwtValidate(TokenString string) int {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(TokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if !token.Valid || err != nil {
		panic("-- token not valid")
	}
	return claims.UID
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-- signup")
	var (
		uid       int
		passwordS string
		PID       string
		ItemA     string
		Title     string
	)
	type RespStruct struct {
		UID         string
		TokenString string
		PID         string
		ItemA       string
		Title       string
	}

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
	err = result.Scan(&uid, &passwordS)

	_, err = db.Query(
		"INSERT INTO plans (uid, title, itemA) VALUES ($1, $2, $3)", uid, "", "[]",
	)
	if err != nil {
		panic(err)
	}

	result = db.QueryRow("SELECT pid, title, itemA FROM plans WHERE uid=$1", uid)
	err = result.Scan(&PID, &Title, &ItemA)

	claims := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		UID:         strconv.Itoa(uid),
		TokenString: TokenString,
		PID:         PID,
		ItemA:       ItemA,
		Title:       Title,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var (
		UID       int
		email     string
		password  string // inputted password
		passwordS string // stored password
		PID       string
		ItemA     string
		Title     string
	)
	type RespStruct struct {
		UID         string
		TokenString string
		PID         string
		ItemA       string
		Title       string
	}

	log.Println("-- login")
	enableCors(&w)
	r.ParseForm()
	email = r.PostFormValue("email")
	password = r.PostFormValue("password")

	result := db.QueryRow("SELECT uid, password FROM users WHERE email=$1", email)
	err = result.Scan(&UID, &passwordS)
	if err != nil { // email not found in database
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordS), []byte(password))
	if err != nil {
		panic(err)
	}

	claims := &Claims{
		UID: UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}

	result = db.QueryRow("SELECT pid, title, itemA FROM plans WHERE uid=$1", claims.UID)
	err = result.Scan(&PID, &Title, &ItemA)

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		UID:         strconv.Itoa(UID),
		TokenString: TokenString,
		PID:         PID,
		ItemA:       ItemA,
		Title:       Title,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-- refresh")
	var (
		UID   int
		PID   int
		Title string
		ItemA string
	)
	type RespStruct struct {
		UID   string
		PID   string
		ItemA string
		Title string
	}

	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	UID = jwtValidate(TokenString)

	result := db.QueryRow("SELECT pid, title, itemA FROM plans WHERE uid=$1", UID)
	err = result.Scan(&PID, &Title, &ItemA)

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		UID:   strconv.Itoa(UID),
		PID:   strconv.Itoa(PID),
		ItemA: ItemA,
		Title: Title,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// save a plan
	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	pid := r.PostFormValue("PID")
	itemA := r.PostFormValue("ItemA")
	log.Println(TokenString)
	uid := jwtValidate(TokenString)

	if err != nil {
		log.Println(err)
	}

	_, err = db.Query(
		"UPDATE plans SET itemA=$1 WHERE uid=$2 and pid=$3", itemA, uid, pid,
	)
	if err != nil {
		panic(err)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	// returns a list of plans
	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	UID := jwtValidate(TokenString)

	type PlanStruct struct {
		PID   string
		Title string
	}
	var planA []PlanStruct

	rows, err := db.Query("SELECT pid, title FROM plans WHERE uid=$1", UID)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var PID int
		var Title string
		err := rows.Scan(&PID, &Title)
		if err != nil {
			panic(err)
		}
		planA = append(planA, PlanStruct{
			PID:   strconv.Itoa(PID),
			Title: Title,
		})
	}
	log.Println(json.Marshal(planA))
}

func addPlanHandler(w http.ResponseWriter, r *http.Request) {
	var PID int
	type RespStruct struct {
		PID string
	}

	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	UID := jwtValidate(TokenString)

	_ = db.QueryRow(
		"INSERT INTO plans (uid, title, itemA) VALUES ($1, $2, $3) RETURNING pid", UID, "New Plan", "[]",
	).Scan(&PID)

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		PID: strconv.Itoa(PID),
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	// load a plan
	enableCors(&w)
	r.ParseForm()
	PID := r.PostFormValue("pid")
	TokenString := r.PostFormValue("TokenString")
	UID := jwtValidate(TokenString)
	log.Println(PID)
	log.Println(UID)
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/refresh", refreshHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/load", loadHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
