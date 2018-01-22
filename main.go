package main

import (
	"fmt"
	// "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"errors"
)


type User struct {
	Name        string
	Email 			string
	Password  	string
	Validated   bool
	PhoneNumber   string
	Country   	string
	City   			string
	Address   	string
}

var (
	ErrRequest = errors.New("request error")
	ErrNameRequired = errors.New("the field 'name' is required")
	ErrEmailRequired = errors.New("the field 'email' is required")
	ErrPasswordRequired = errors.New("the field 'password' is required")
	ErrPhoneRequired = errors.New("the field 'phoneNumber' is required")
)

func ValidateInputs(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
    return ErrRequest
  }

	if len(r.FormValue("name")) == 0 {
		return ErrNameRequired
	}
	if len(r.FormValue("email")) == 0  {
		return ErrEmailRequired
	}
	if len(r.FormValue("password")) == 0  {
		return ErrPasswordRequired
	}
	if len(r.FormValue("phoneNumber")) == 0  {
		return ErrPhoneRequired
	}
	return nil
}

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {

	err := ValidateInputs(r)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	decoder := schema.NewDecoder()
// r.PostForm is a map of our POST form values
	user := new(User)
	derr := decoder.Decode(user, r.PostForm)

	if derr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(derr.Error()))
	}

	fmt.Println("**************************")
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	fmt.Printf("%+v\n", user)
	fmt.Println(user.PhoneNumber)

}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUserEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
