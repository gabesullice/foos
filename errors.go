package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerError struct {
	ResponseCode    int
	Status, Message string
}

type ServerErrors map[string]ServerError

var ServeErrors = ServerErrors{
	"userNotFound": ServerError{
		http.StatusNotFound,
		"Not Found", "User not found",
	},
	"gameNotFound": ServerError{
		http.StatusNotFound,
		"Not Found", "Game not found",
	},
	"badPostBody": ServerError{
		http.StatusInternalServerError,
		"Internal Server Error", "Could not read post body",
	},
	"badJSON": ServerError{
		http.StatusBadRequest,
		"Bad Request", "Unable to parse JSON",
	},
	"badGame": ServerError{
		422, // StatusUnprocessableEntity
		"Unprocessable Entity", "The game(s) could not be created",
	},
	"badKey": ServerError{
		422, // StatusUnprocessableEntity
		"Unprocessable Entity", "The key(s) could not be created",
	},
	"badDbOp": ServerError{
		http.StatusInternalServerError,
		"Internal Server Error", "Error operating on database",
	},
	"badResponse": ServerError{
		http.StatusInternalServerError,
		"Internal Server Error", "Could not create response body",
	},
}

func (err ServerError) Error() string {
	return fmt.Sprintf("%d %s: %s", err.ResponseCode, err.Status, err.Message)
}

func ServeError(w http.ResponseWriter, err ServerError) {
	response, marshalErr := json.Marshal(err)
	if marshalErr != nil {
		panic(marshalErr)
	}
	w.WriteHeader(err.ResponseCode)
	fmt.Fprintf(w, string(response))
}

func (err ServerError) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(
		"{\"Status\":\"%d %s\",\"Message\":\"%s\"}",
		err.ResponseCode, err.Status, err.Message,
	)), nil
}
