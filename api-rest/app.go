package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "os"
	"time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) getRecipe(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
        return
    }

    rec := recipe{ID: id}
    if err := rec.getRecipe(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Recipe not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    respondWithJSON(w, http.StatusOK, rec)
}

func (a *App) getRecipes(w http.ResponseWriter, r *http.Request) {   
    /*
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))
    
    if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
    }
    */
    recipes, err := getRecipes(a.DB)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, recipes)
}

func (a *App) createRecipe(w http.ResponseWriter, r *http.Request) {
    
    var rec recipe
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&rec); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return 
    }
    defer r.Body.Close()

    if err := rec.createRecipe(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, rec)
}

func (a *App) updateRecipe(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
        return
    }

    var rec recipe
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&rec); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    rec.ID = id

    if err := rec.updateRecipe(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, rec)
}

func (a *App) deleteRecipe(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
        return
    }

    rec := recipe{ID: id}
    if err := rec.deleteRecipe(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {    
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {    
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/recipes", a.getRecipes).Methods("GET")
    a.Router.HandleFunc("/recipe", a.createRecipe).Methods("POST")
    a.Router.HandleFunc("/recipe/{id:[0-9]+}", a.getRecipe).Methods("GET")
    a.Router.HandleFunc("/recipe/{id:[0-9]+}", a.updateRecipe).Methods("PUT")
    a.Router.HandleFunc("/recipe/{id:[0-9]+}", a.deleteRecipe).Methods("DELETE")
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
    a.DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
    handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(a.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
    a.initializeRoutes()
	log.Fatal(newServer.ListenAndServe())
    
 }

func (a *App) Run(addr string) { 
    //log.Fatal(http.ListenAndServe(addr, a.Router))
}