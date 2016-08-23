package main

import "net/http"

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes {
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "Accounts",
    "GET",
    "/accounts",
    GetAll,
  },
  Route{
    "New Account",
    "POST",
    "/accounts",
    CreateAccount,
  },
  Route{
    "Account Detail",
    "GET",
    "/accounts/{accountId}",
    GetAccount,
  },
}
