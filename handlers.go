package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	g "github.com/gabesullice/foos/game"
	u "github.com/gabesullice/foos/user"
)

func UsersList(w http.ResponseWriter, r *http.Request) {
	users, err := u.GetUsers(storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["userNotFound"])
		panic(err)
	}

	response := marshal(w, users)
	respond(w, string(response), http.StatusOK)
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	body := readBody(w, r)
	var users []*u.User
	unmarshal(w, body, &users)
	usersCreate(w, &users)
	response := marshal(w, users)
	respond(w, string(response), http.StatusCreated)
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := u.GetUser(vars["mail"], storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["userNotFound"])
		panic(err)
	}

	response := marshal(w, user)
	respond(w, string(response), http.StatusOK)
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := u.GetUser(vars["mail"], storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["userNotFound"])
		panic(err)
	}

	body := readBody(w, r)
	unmarshal(w, body, &user)

	user.Save(storage.GetSession())

	response := marshal(w, user)
	respond(w, string(response), http.StatusAccepted)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := u.GetUser(vars["mail"], storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["userNotFound"])
		panic(err)
	}

	err = user.Delete(storage.GetSession())
	if err != nil {
		ServeError(w, ServeErrors["badDbOp"])
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func GamesList(w http.ResponseWriter, r *http.Request) {
	games, err := g.GetGames(storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["gameNotFound"])
		panic(err)
	}

	response := marshal(w, games)
	respond(w, string(response), http.StatusOK)
}

func GamesCreate(w http.ResponseWriter, r *http.Request) {
	body := readBody(w, r)
	var games []*g.Game
	unmarshal(w, body, &games)
	gamesCreate(w, &games)
	response := marshal(w, games)
	respond(w, string(response), http.StatusOK)
}

func GameDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, err := g.GetGame(vars["id"], storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["gameNotFound"])
		panic(err)
	}

	response := marshal(w, game)
	respond(w, string(response), http.StatusOK)
}

func GameUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, err := g.GetGame(vars["id"], storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["gameNotFound"])
		panic(err)
	}

	body := readBody(w, r)
	unmarshal(w, body, &game)

	game.Save(storage.GetSession())

	response := marshal(w, game)
	respond(w, string(response), http.StatusAccepted)
}

func GameDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, err := g.GetGame(vars["id"], storage.GetSession())

	if err != nil {
		ServeError(w, ServeErrors["gameNotFound"])
		panic(err)
	}

	err = game.Delete(storage.GetSession())
	if err != nil {
		ServeError(w, ServeErrors["badDbOp"])
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func readBody(w http.ResponseWriter, r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		ServeError(w, ServeErrors["badPostBody"])
		panic(err)
	}
	return body
}

func respond(w http.ResponseWriter, response string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	fmt.Fprintf(w, response)
}

func marshal(w http.ResponseWriter, v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		ServeError(w, ServeErrors["badResponse"])
		panic(err)
	}
	return data
}

func unmarshal(w http.ResponseWriter, data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		ServeError(w, ServeErrors["badJSON"])
		panic(err)
	}
}

func usersCreate(w http.ResponseWriter, users *[]*u.User) {
	for _, user := range *users {
		s := storage.GetSession()
		if err := user.Check(s); err != nil {
			ServeError(w, ServeErrors["badUser"])
			panic(err)
		}
		err := user.Save(s)
		if err != nil {
			ServeError(w, ServeErrors["badDbOp"])
			panic(err)
		}
	}
	return
}

func usersDelete(w http.ResponseWriter, users []u.User) []u.User {
	for _, user := range users {
		s := storage.GetSession()
		err := user.Delete(s)
		if err != nil {
			ServeError(w, ServeErrors["badDbOp"])
			panic(err)
		}
	}
	return users
}

func gamesCreate(w http.ResponseWriter, games *[]*g.Game) {
	for _, game := range *games {
		s := storage.GetSession()
		err := game.Save(s)
		if err != nil {
			ServeError(w, ServeErrors["badDbOp"])
			panic(err)
		}
	}
	return
}

func gamesDelete(w http.ResponseWriter, games []g.Game) []g.Game {
	for _, game := range games {
		s := storage.GetSession()
		err := game.Delete(s)
		if err != nil {
			ServeError(w, ServeErrors["badDbOp"])
			panic(err)
		}
	}
	return games
}
