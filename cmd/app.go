package cmd

import (
	"XM/router"
	"fmt"
	"log"
	"net/http"
)

//App struct
type App struct {
	appName string
}

//NewApp returns new app struct
func NewApp() *App {
	return &App{
		appName: "XM Assignment",
	}
}

//Serve to serve the server
func (a *App) Serve() {
	port, err := getPort()
	if err != nil {
		log.Println("Error : Can't find the server address")
		panic(err)
	}
	r := router.Setup()
	log.Println("App : Server is listening")
	fmt.Println(http.ListenAndServe("localhost:"+port, r))
}
