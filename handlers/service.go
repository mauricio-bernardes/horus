package handlers

import (
	"encoding/json"
	"fmt"
	"horus-api/storage"
	"horus-api/types"
	"log"
	"net/http"
)

func HandleInserService(w http.ResponseWriter, r *http.Request) {
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

func HandleListServices(w http.ResponseWriter, r *http.Request) {
	var s types.Services

	services := storage.GetServicesNames()
	for _, service := range services {
		s.Services = append(s.Services, service)
	}

	json, err := json.Marshal(s)
	if err != nil {
		log.Println("Erro ao fazer encode do json", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, string(json), http.StatusOK)

}

func HandleRemoveService(w http.ResponseWriter, r *http.Request) {
	var s types.Services
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		log.Println("Erro ao fazer decode do json", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, name := range s.Services {
		err := storage.RemoveServiceName(name)
		if err != nil {
			log.Println("Erro ao remover serviço no redis", err.Error())
			http.Error(w, fmt.Sprint("Erro ao remover nome do serviço no redis: ", err.Error()), http.StatusInternalServerError)
			return
		}
	}
	http.Error(w, "Serviço removido com sucesso!", http.StatusOK)
}
