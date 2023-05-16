package main

import (
	//"fmt"
	"horus/configs"
	//"horus/handlers"
	"horus/util"
	//"net/http"
	"time"
	//"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	// r := chi.NewRouter()
	// r.Post("/subscribe", handlers.HandleSubscribe)
	// r.Post("/add", handlers.HandleInserService)

	// http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPIConfig().Port), r)
	// fmt.Println("Server running on port", configs.GetAPIConfig().Port)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			util.GetServicesStatus()
		}
	}
}
