package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/scvillon/workshop-uva/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Handlers são Todos os handlers da aplicação
type Handlers struct {
	Db *mgo.Collection
}

// CreatePersonHandler cadastra uma nova pessoa
// observe o pointer receiver!
func (h *Handlers) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	pessoa := &models.Person{}
	if err := json.NewDecoder(r.Body).Decode(&pessoa); err != nil {
		log.Print(r.Body)
		log.Printf("Erro ao parsear o corpo da requisição: %s", err.Error())
		http.Error(w, "Error to parse body", http.StatusBadRequest)
		return
	}
	if err := h.Db.Insert(pessoa); err != nil {
		log.Printf("Erro inserir pessoa no Banco: %s", err.Error())
		http.Error(w, "Erro inserir pessoa no Banco", http.StatusInternalServerError)
		return
	}
	resp, _ := json.Marshal(pessoa)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

// FindPersonHandler procura uma pessoa
func (h *Handlers) FindPersonHandler(w http.ResponseWriter, r *http.Request) {
	result := models.Person{}
	term := r.URL.Query().Get(":name")
	err := h.Db.Find(bson.M{"name": term}).One(&result)
	if err != nil {
		log.Print("Erro ao buscar " + term + ": " + err.Error())
		http.Error(w, "Pessoa não encontrada!", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Phone: "+result.Phone)
}

// ListPeopleHandler lista todas as Pessoas
func (h *Handlers) ListPeopleHandler(w http.ResponseWriter, r *http.Request) {
	result := []models.Person{}
	err := h.Db.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(w, "Hi there!")
	w.Header().Set("Content-Type", "application/json charset=utf-8")
	if resp, err := json.Marshal(result); err != nil {
		fmt.Fprintf(w, "Erro ao serializar as pessoas!")
	} else {
		fmt.Fprintf(w, string(resp))
	}
}
