package main

import (
	"gateway/datebase"
	"gateway/middlewares"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//ctx := context.Background()

	middlewares.SetupMiddlewares(mux)
	datebase.OpenConnection()

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
