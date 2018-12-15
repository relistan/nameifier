package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/relistan/apistuff"
	"gopkg.in/relistan/rubberneck.v1"
)

type Config struct {
	Port int `envconfig:"PORT" default:"9001"`
}

func nameHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	namer := NewNameifier(".")

	if len(params["count"]) < 1 {
		apistuff.HttpError(w, "Bad count parameter", 400)
		return
	}

	if len(params["seed"]) < 1 {
		apistuff.HttpError(w, "Bad seed parameter", 400)
		return
	}

	count, err := strconv.Atoi(params["count"])
	if err != nil {
		apistuff.HttpError(w, "Bad count parameter", 400)
		return
	}

	var allNames []string
	for i := 0; i < count; i++ {
		name, err := namer.Nameify(fmt.Sprintf("%s-%d", params["seed"], i))
		if err != nil {
			apistuff.HttpError(w, "Failed to nameify!", 500)
			return
		}
		allNames = append(allNames, name)
	}

	w.Write([]byte(strings.Join(allNames, "\n")))
}

func blankHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(204)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}

func main() {

	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		println("Unable to parse config " + err.Error())
		os.Exit(1)
	}
	rubberneck.Print(config)

	uiFs := http.FileServer(http.Dir("ui"))
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	router.HandleFunc("/nameifier/{seed}/{count}", nameHandler).Methods("GET")
	router.HandleFunc("/nameifier/{seed}", blankHandler).Methods("GET")
	router.HandleFunc("/nameifier/", blankHandler).Methods("GET")
	router.HandleFunc("/nameifier/", blankHandler).Methods("GET")
	router.PathPrefix("/").Handler(uiFs)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Port), nil)
}
