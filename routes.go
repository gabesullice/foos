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
	Route{"/users/{mail}", "GET", "UserDetail", UserDetail},
	Route{"/users/{mail}", "PATCH", "UserUpdate", UserUpdate},
	Route{"/users/{mail}", "DELETE", "UserDelete", UserDelete},

	Route{"/games", "GET", "GameList", GamesList},
	Route{"/games", "POST", "GameCreate", GamesCreate},
	Route{"/games/{id}", "GET", "GameDetail", GameDetail},
	Route{"/games/{id}", "PATCH", "GameUpdate", GameUpdate},
	Route{"/games/{id}", "DELETE", "GameDelete", GameDelete},
}
