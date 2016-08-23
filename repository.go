package main
import (
  "fmt"
  "time"
)

var currentId int
var accounts Accounts

func init() {
  RepoSave(Account{Name:"Gabriel", BirthDate: time.Date(1996, time.August, 8, 10, 0, 45, 22, time.UTC), IsActive:true})
  RepoSave(Account{Name:"Dewes",   BirthDate: time.Date(1995, time.November, 23, 9, 0, 56, 10, time.UTC), IsActive:false})
  RepoSave(Account{Name:"Abel",    BirthDate: time.Date(2000, time.January, 5, 13, 33, 54, 0, time.UTC), IsActive:true})
}

func RepoFindById(id int) Account {
  for _, a := range accounts {
    if a.AccountId == id {
      return a
    }
  }
  return Account{}
}

func RepoSave(a Account) Account {
  currentId += 1
  a.AccountId = currentId
  accounts = append(accounts, a)
  return a
}

func RepoDelete(id int) error {
  for i, a := range accounts {
    if (a.AccountId == id) {
      accounts = append(accounts[:i], accounts[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Não foi possível excluir. Motivo: ID %d não encontrado.")
}
