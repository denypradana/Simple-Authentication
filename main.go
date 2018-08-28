package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

// User mewakili semua informasi mengenai user
type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
}

func connectDb() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/go_db")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func routes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)
}

func main() {
	connectDb()
	routes()

	defer db.Close()

	fmt.Println("Server berjalan pada port : 8888")
	http.ListenAndServe(":8888", nil)
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}

// QueryUser untuk membaca list user di dalam database
func QueryUser(username string) User {
	var users = User{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
		)
	return users
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	username := r.FormValue("email")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	password := r.FormValue("password")

	users := QueryUser(username)

	if (User{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &firstName, &lastName)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w, r, "/register", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		http.Redirect(w, r, "/home", 302)
	}
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	//deskripsi dan compare password
	var passwordTes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if passwordTes == nil {
		// login berhasil
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("name", users.FirstName)
		http.Redirect(w, r, "/home", 302)
	} else {
		// login gagal
		http.Redirect(w, r, "/login", 302)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	// Jika user belum login maka akan diarahkan ke form login
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
	// Jika user sudah login, map data pengguna untuk ditampilkan pada halaman Home
	var data = map[string]string{
		"username": session.GetString("name"),
		"message":  "Selamat datang di Go Login !",
	}
	var t, err = template.ParseFiles("views/home.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, data)
	return

}

func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}
