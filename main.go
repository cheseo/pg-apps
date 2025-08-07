package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"html/template"
)

const (
	port = 5432
	user = "ubuntu"
	password = "ubuntu"
	dbname = "postgres"
)

var db *sql.DB
func main(){
	cs := fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable",  port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", cs)
	if err != nil {
		log.Panic("couldn't connect to postgres ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic("no reply ", err)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/", index)
	http.HandleFunc("/new", New)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func New(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		fmt.Fprintln(w, "empty name")
		return
	}
	id, err := Insert(name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "the new id is: ", id)
}

func index(w http.ResponseWriter, r *http.Request) {
	p :=`
<!DOCTYPE html>
<html>
<head><title>Postgres!</title></head>
<body>
<form action="/new" method="GET">
<input type="text" name="name" placeholder="name" >
<input type="submit">
</form>
<h3>The available users are</h3>
<table border="1">
`
	fmt.Fprintln(w, p)
	tmpl, err := template.New("users").Parse("<tr><td>{{.Id}}</td><td>{{.Name}}</tr>")
	if err != nil {
		fmt.Fprintln(w, err)
		log.Println(err)
	}
	u := GetUsers()
	for _, uu := range u {
		err = tmpl.Execute(w, uu)
		if err != nil {
			fmt.Fprintln(w, err)
			log.Println(err)
		}
	}
	e :=`
</table>
</body>
</html>
`
	fmt.Fprintln(w, e)
}


func Insert(name string) (int, error) {
	r := db.QueryRow("insert into users(name) values($1) returning id;", name)
	if r.Err() != nil {
		log.Println(r.Err())
		return 0, r.Err()
	}
	id := 0
	r.Scan(&id)
	return id, nil
}

type Users struct {
	Id int
	Name string
}

func GetUsers() ([]Users) {
	var u []Users
	r, err := db.Query("select id, name from users;")
	if err != nil {
		log.Println(err)
		return u
	}
	for r.Next() {
		id, n := 0, ""
		r.Scan(&id, &n)
		u = append(u, Users{id, n})
	}
	if r.Err() != nil {
		log.Println(err)
		return []Users{Users{0, fmt.Sprint("error: ", err)}}
	}
	return u
}
