package main

import (
	"net/http"
)

type Route struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"/users", "GET", "UserList", UsersList},
	Route{"/users", "POST", "UserCreate", UsersCreate},
	Route{"/users/{name}", "GET", "UserDetail", UserDetail},
	Route{"/users/{name}", "PATCH", "UserUpdate", UserUpdate},
	Route{"/users/{name}", "DELETE", "UserDelete", UserDelete},

	Route{"/games", "GET", "GameList", GamesList},
	Route{"/games", "POST", "GameCreate", GamesCreate},
	Route{"/games/{name}", "GET", "GameDetail", GameDetail},
	Route{"/games/{name}", "PATCH", "GameUpdate", GameUpdate},
	Route{"/games/{name}", "DELETE", "GameDelete", GameDelete},
}