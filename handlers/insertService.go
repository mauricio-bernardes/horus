package handlers

import (
	"encoding/json"
	"fmt"
	"horus/storage"
	"horus/types"
	"log"
	"net/http"
)

func HandleInserService(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	var s types.Services
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		log.Println("Erro ao fazer decode do json", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, name := range s.Services {
		err := storage.SetServiceName(name)
		if err != nil {
			log.Println("Erro ao inserir serviço no redis", err.Error())
			http.Error(w, fmt.Sprint("Erro ao inserir nome do serviço no redis: ", err.Error()), http.StatusInternalServerError)
			return
		}
	}
	http.Error(w, "Serviço adicionado com sucesso!", http.StatusOK)
}
