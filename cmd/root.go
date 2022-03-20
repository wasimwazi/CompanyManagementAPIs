package cmd

import "log"

//Begin is the beginning of the app
func Begin() {
	err := checkEnv()
	if err != nil {
		log.Println("Error in environment variable", err.Error())
		panic(err)
	}
	log.Println("App : Database connected successfully")
	app := NewApp()
	app.Serve()
}
