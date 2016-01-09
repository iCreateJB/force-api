package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"runtime"
	"github.com/bmizerany/pat"
	"github.com/satori/go.uuid"

	"database/sql"

	_ "github.com/lib/pq"
)

const appVersionStr = "1.1"

type Moral struct {
	Message  string  `json:"quote"`
	Category string  `json:"category"`
	Key      string  `json:"request_id"`
	Errors   []Error `json:"errors",omitempty`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func (m *Moral) create() (*Moral, bool) {
	m,v := m.valid()
	if v {
		_, err := db.Exec("insert into morals ( category,quote,created_on ) values ( $1, $2, now() )", m.Category, m.Message)
		if err != nil {
			log.Fatal(err)
		}
	}
	return m, v
}

func (m *Moral) valid() (*Moral, bool) {
	if m.Message == "" {
		m.Errors = append(m.Errors, Error{ Message: "Message can't be blank." })
	}
	return m, len(m.Errors) == 0
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "application/vnd.morals."+appVersionStr+"+json")
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	}
}

func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		fn(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}

func requestId() string {
	req_id := uuid.NewV4()
	reqId, _ := req_id.Value()
	return reqId.(string)
}

func totalMorals() int {
	var count int
	err := db.QueryRow("select count(*) from morals where category = 'philosophy'").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func track(moralId int) {
	_, err := db.Exec("insert into metrics ( moral_id, created_on ) values ( $1, now() )", moralId)
	if err != nil {
		log.Fatal(err)
	}
}

func getMoral() Moral {
	var moral Moral
	err := db.QueryRow("select quote,category from morals offset floor(random()*(select count(*) from morals where category is not null)) limit 1").Scan(&moral.Message, &moral.Category)
	if err != nil {
		log.Fatal(err)
	}
	// track(key)
	return moral
}

func MoralHandler(w http.ResponseWriter, r *http.Request) {
	quote := getMoral()
	moral := &Moral{Message: quote.Message, Category: strings.TrimSpace(quote.Category), Key: requestId()}
	resp, _ := json.Marshal(moral)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func NewMoralHandler(w http.ResponseWriter, r *http.Request) {
	moral := new(Moral)
	json.NewDecoder(r.Body).Decode(moral)
	moral, v  := moral.create()
	if v {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	resp, _ := json.Marshal(&Moral{Message: moral.Message, Category: moral.Category, Key: requestId(), Errors: moral.Errors})
	w.Write(resp)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(&Moral{})
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func portNumber() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}

func init() {
	var err error
	db, err = sql.Open("postgres", "user=readerwriter dbname=morals sslmode=disable")
	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	defer db.Close()
	m := pat.New()
	m.Get("/morals", commonHeaders(logHandler(MoralHandler)))
	m.Post("/morals", commonHeaders(logHandler(NewMoralHandler)))
	m.Get("/", commonHeaders(logHandler(IndexHandler)))
	http.Handle("/", m)
	http.ListenAndServe(":"+portNumber(), nil)
}
