/*

	nohup JWTSECRET= go run main.go &

	go get golang.org/x/crypto/bcrypt
	go get github.com/lib/pq
	go get github.com/dgrijalva/jwt-go

	CREATE DATABASE frenzydb
	CREATE TABLE users (uid BIGSERIAL, email TEXT, password BYTEA, ts TIMESTAMP, primary key (uid, email, ts));
	CREATE TABLE plans (pid BIGSERIAL, uid BIGINT, title TEXT, itema TEXT, ts TIMESTAMP, primary key (pid, uid, ts));

	go run server/main.go

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
	"os"
	"strconv"
	"time"
)

var jwtSecret []byte = []byte(os.Getenv("JWTSECRET"))
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
	if token != nil && err != nil {
		panic("-- token not valid")
	}
	return claims.UID
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- signupHandler")
	var (
		UID   int
		PID   string
		ItemA string
		Title string
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
	Title = "Blank Title"
	ItemA = "[]"

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	ts := time.Now()
	_ = db.QueryRow(
		"INSERT INTO users (email, password, ts) VALUES ($1, $2, $3) RETURNING uid", email, string(hashed), ts,
	).Scan(&UID)
	_ = db.QueryRow(
		"INSERT INTO plans (uid, title, itema, ts) VALUES ($1, $2, $3, $4) RETURNING pid", UID, Title, ItemA, ts,
	).Scan(&PID)

	claims := &Claims{
		UID:            UID,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}

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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- loginHandler")
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

	enableCors(&w)
	r.ParseForm()
	email = r.PostFormValue("email")
	password = r.PostFormValue("password")

	result := db.QueryRow("SELECT uid, password FROM users WHERE email=$1 ORDER BY ts DESC", email)
	err = result.Scan(&UID, &passwordS)
	if err != nil { // email not found in database
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordS), []byte(password))
	if err != nil {
		panic(err)
	}

	claims := &Claims{
		UID:            UID,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}

	result = db.QueryRow("SELECT pid, title, itema FROM plans WHERE uid=$1 ORDER BY ts DESC", claims.UID)
	err = result.Scan(&PID, &Title, &ItemA)
	if err != nil {
		ts := time.Now()
		_, err = db.Exec(
			"INSERT INTO plans (uid, title, itema, ts) VALUES ($1, 'Blank Title', '[]', $2)", UID, ts,
		)
		if err != nil {
			panic(err)
		}
		result = db.QueryRow("SELECT pid, title, itema FROM plans WHERE uid=$1 ORDER BY ts DESC", UID)
		err = result.Scan(&PID, &Title, &ItemA)
	}

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
	fmt.Println("-- refreshHandler")
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

	result := db.QueryRow("SELECT pid, title, itema FROM plans WHERE uid=$1 ORDER BY ts DESC", UID)
	err = result.Scan(&PID, &Title, &ItemA)

	if err != nil {
		ts := time.Now()
		_, err = db.Exec(
			"INSERT INTO plans (uid, title, itema, ts) VALUES ($1, 'Blank Title', '[]', $2)", UID, ts,
		)
		if err != nil {
			panic(err)
		}
		result = db.QueryRow("SELECT pid, title, itema FROM plans WHERE uid=$1 ORDER BY ts DESC", UID)
		err = result.Scan(&PID, &Title, &ItemA)
	}

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
	fmt.Println("-- saveHandler")
	// save a plan
	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	pid := r.PostFormValue("PID")
	itemA := r.PostFormValue("ItemA")
	uid := jwtValidate(TokenString)

	if err != nil {
		panic(err)
	}

	ts := time.Now()
	_, err = db.Exec(
		"UPDATE plans SET itema=$1, ts=$2 WHERE uid=$3 and pid=$4", itemA, ts, uid, pid,
	)
	if err != nil {
		panic(err)
	}
}

func titleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- titleHandler")
	// update title of plan
	type RespStruct struct {
		Title string
	}
	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	PID := r.PostFormValue("PID")
	Title := r.PostFormValue("Title")
	UID := jwtValidate(TokenString)

	ts := time.Now()
	_, err = db.Exec(
		"UPDATE plans SET title=$1, ts=$2 WHERE uid=$3 and pid=$4", Title, ts, UID, PID,
	)
	if err != nil {
		panic(err)
	}

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		Title: Title,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))

}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- listHandler")
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

	type RespStruct struct {
		PlanA []PlanStruct
	}

	rows, err := db.Query("SELECT pid, title FROM plans WHERE uid=$1 ORDER BY ts DESC", UID)
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

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		PlanA: planA,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- addHandler")
	var (
		PID   int
		ItemA string
		Title string
	)
	type RespStruct struct {
		PID   string
		ItemA string
		Title string
	}

	enableCors(&w)
	r.ParseForm()
	TokenString := r.PostFormValue("TokenString")
	UID := jwtValidate(TokenString)

	ts := time.Now()
	_ = db.QueryRow(
		"INSERT INTO plans (uid, title, itema, ts) VALUES ($1, 'Blank Title', '[]', $2)", UID, ts,
	).Scan(&PID)

	result := db.QueryRow("SELECT pid, title, itema FROM plans WHERE uid=$1 AND pid=$2 ORDER BY ts DESC", UID, PID)
	err = result.Scan(&PID, &Title, &ItemA)

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		PID:   strconv.Itoa(PID),
		ItemA: strconv.Itoa(PID),
		Title: Title,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- deleteHandler")
	enableCors(&w)
	r.ParseForm()
	PID := r.PostFormValue("PID")
	TokenString := r.PostFormValue("TokenString")
	UID := jwtValidate(TokenString)
	db.Query("DELETE FROM plans WHERE uid=$1 AND pid=$2", UID, PID)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- loadHandler")
	// load a plan
	var (
		Title string
		ItemA string
	)
	type RespStruct struct {
		PID   string
		ItemA string
		Title string
	}
	enableCors(&w)
	r.ParseForm()
	PID := r.PostFormValue("PID")
	TokenString := r.PostFormValue("TokenString")
	UID := jwtValidate(TokenString)

	result := db.QueryRow("SELECT pid, title, itema FROM plans WHERE uid=$1 AND pid=$2 ORDER BY ts DESC", UID, PID)
	err = result.Scan(&PID, &Title, &ItemA)

	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(RespStruct{
		PID:   PID,
		ItemA: ItemA,
		Title: Title,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/refresh", refreshHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/title", titleHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/load", loadHandler)
	http.HandleFunc("/delete", deleteHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
