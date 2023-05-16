package handlers

import (
	"encoding/json"
	"fmt"
	"horus/storage"
	"horus/types"
	"log"
	"net/http"
)

func HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	var s types.Subscribed
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		log.Println("Erro ao fazer decode do json", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = storage.SubscribeClient(s.Url)
	if err != nil {
		log.Println("Erro ao inserir cliente no redis", err.Error())
		http.Error(w, fmt.Sprint("Erro ao inserir cliente no redis: ", err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Println(s)
	for _, name := range s.Names {
		fmt.Println("name", name)
	}
	http.Error(w, "Subscribe efetuado com sucesso!", http.StatusOK)
}
