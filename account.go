package main

import "time"

type Account struct {
  AccountId int       `json:"id"`
  Name      string    `json:"name"`
  BirthDate time.Time `json:"birthDate"`
  IsActive  bool      `json:"isActive"`
}

type Accounts []Account
