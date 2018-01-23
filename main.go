package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/judaro13/users_enpoint/publisher"
)

type User struct {
	Name        string
	Email       string
	Password    string
	Validated   bool
	PhoneNumber string
	Country     string
	City        string
	Address     string
}

var (
	ErrRequest          = errors.New("request error")
	ErrNameRequired     = errors.New("the field 'name' is required")
	ErrEmailRequired    = errors.New("the field 'email' is required")
	ErrPasswordRequired = errors.New("the field 'password' is required")
	ErrPhoneRequired    = errors.New("the field 'phoneNumber' is required")
)

func ValidateInputs(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return ErrRequest
	}

	if len(r.FormValue("name")) == 0 {
		return ErrNameRequired
	}
	if len(r.FormValue("email")) == 0 {
		return ErrEmailRequired
	}
	if len(r.FormValue("password")) == 0 {
		return ErrPasswordRequired
	}
	if len(r.FormValue("phoneNumber")) == 0 {
		return ErrPhoneRequired
	}
	return nil
}

func ReturnWithError(w http.ResponseWriter, err error, status int) bool {
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return true
	}
	return false
}

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	err := ValidateInputs(r)
	if ReturnWithError(w, err, http.StatusBadRequest) {
		return
	}

	decoder := schema.NewDecoder()
	user := new(User)
	derr := decoder.Decode(user, r.PostForm)

	if ReturnWithError(w, derr, http.StatusInternalServerError) {
		return
	}

	jsonUser, jerr := json.Marshal(user)
	if ReturnWithError(w, jerr, http.StatusInternalServerError) {
		return
	}

	serr := publisher.SendMessage(string(jsonUser))
	if ReturnWithError(w, serr, http.StatusInternalServerError) {
		return
	}

	w.Write([]byte(jsonUser))
}

func IndexEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("running app.."))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexEndpoint).Methods("GET")
	router.HandleFunc("/users", CreateUserEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
