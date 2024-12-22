package main

import (
	"context"
	"database/sql"
	dbsqlc "eudald/interview/SQL"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	host     = os.Getenv("DBHOST")
	port     = os.Getenv("DBPORT")
	user     = os.Getenv("DBUSER")
	password = os.Getenv("DBPASSWORD")
	dbname   = os.Getenv("DBNAME")
)

var Q *dbsqlc.Queries

func main() {
	port, err := strconv.Atoi(port)
	psdata := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psdata)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	Q = dbsqlc.New(db)

	
	http.HandleFunc("/", home)
	http.HandleFunc("/get-users", getUsers)
	http.HandleFunc("/create-user", createUser)
	http.HandleFunc("/delete-user", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("HTML/index.html")
	if err != nil {
		log.Println(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := Q.GetUsers(context.Background())
	if err != nil {
		log.Println(err)
	}
	for _, v := range users {
		fmt.Fprintf(w, "<li> %v </li>", v)
	}
}

func createUser(_ http.ResponseWriter, r *http.Request) {
	lid, _ := Q.LastId(context.Background())
	lid++
	name := r.PostFormValue("name")

	err := Q.CreateUser(context.Background(), dbsqlc.CreateUserParams{ID: lid, Name: name, Status: sql.NullBool{Bool: false, Valid: false}})
	if err != nil {
		panic(err)
	}
}

func deleteUser(_ http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	err := Q.DeleteUser(context.Background(), name)
	if err != nil {
		log.Println(err)
	}
}
