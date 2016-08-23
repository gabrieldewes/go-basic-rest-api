package main
import (
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Bem vindo!")
}

func GetAll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(accounts); err != nil {
    panic(err)
  }
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var accountId int
  var err error
  if accountId, err = strconv.Atoi(vars["accountId"]); err != nil {
    panic(err)
  }
  account := RepoFindById(accountId)
  if account.AccountId > 0 {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(account); err != nil {
			panic(err)
		}
		return
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
  if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
  var account Account
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &account); err != nil {
    w.Header().Set("Content-Type", "application/json charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  a := RepoSave(account)
  w.Header().Set("Content-Type", "application/json charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(a); err != nil {
    panic(err)
  }
}
