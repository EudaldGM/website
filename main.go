package main

import (
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

/*var (
	host     = os.Getenv("DBHOST")
	port     = os.Getenv("DBPORT")
	user     = os.Getenv("DBUSER")
	password = os.Getenv("DBPASSWORD")
	dbname   = os.Getenv("DBNAME")
)*/

//var Q *dbsqlc.Queries

func main() {
	//port, err := strconv.Atoi(port)
	/*psdata := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psdata)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	Q = dbsqlc.New(db)*/

	http.HandleFunc("/", home)

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
