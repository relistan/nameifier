package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/relistan/apistuff"
	"gopkg.in/relistan/rubberneck.v1"
)

var Config struct {
	Port       int    `short:"p" help:"Port to start service on" default:"9001" env:"PORT"`
	SeedString string `short:"s" help:"Seed string to use for naming" default:"" env:"SEED_STRING"`
	Count      int    `short:"c" help:"Count of unique names to return" default:1 env:"COUNT"`
}

//go:embed ui/index.html
var uiFS embed.FS

func nameHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	namer := NewNameifier()

	if len(params["count"]) < 1 {
		apistuff.HttpError(w, "Bad count parameter", 400)
		return
	}

	if len(params["count"]) < 1 {
		apistuff.HttpError(w, "Bad count parameter", 400)
		return
	}

	if len(params["seed"]) < 1 {
		apistuff.HttpError(w, "Bad seed parameter", 400)
		return
	}

	count, err := strconv.Atoi(params["count"])
	if err != nil || count > 100000 {
		apistuff.HttpError(w, "Bad count parameter. Seriously, you don't need that many.", 400)
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
	kong.Parse(&Config)

	// By default start server
	if Config.SeedString == "" {
		rubberneck.Print(Config)
		serverRoot, err := fs.Sub(uiFS, "ui")
		if err != nil {
			println("Unable to read ui filesystem ", err.Error())
			os.Exit(1)
		}

		uiFs := http.FileServer(http.FS(serverRoot))
		router := mux.NewRouter()
		router.Use(loggingMiddleware)
		router.HandleFunc("/nameifier/{seed}/{count}", nameHandler).Methods("GET")
		router.HandleFunc("/nameifier/{seed}", blankHandler).Methods("GET")
		router.HandleFunc("/nameifier/", blankHandler).Methods("GET")
		router.HandleFunc("/nameifier/", blankHandler).Methods("GET")
		router.PathPrefix("/").Handler(uiFs)
		http.Handle("/", router)
		http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Config.Port), nil)
	} else {
		var allNames []string
		namer := NewNameifier()
		for i := 0; i < Config.Count; i++ {
			name, err := namer.Nameify(fmt.Sprintf("%s-%d", Config.SeedString, i))
			if err != nil {
				fmt.Printf("Failed to nameify! %s", err)
				return
			}
			allNames = append(allNames, name)
		}
		fmt.Println(strings.Join(allNames, "\n"))
	}
}
