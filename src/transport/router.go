package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"oidc_project/dto"
	"oidc_project/models"
	"oidc_project/services"
)

type MainController struct {
	serv *services.MainService
}

func NewMainController(serv services.MainService) *MainController {
	return &MainController{serv: &serv}
}

func (m *MainController) GetData(w http.ResponseWriter, r *http.Request) {
	data, err := m.serv.ReadData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := dto.MainResponse{Data: data}

	json.NewEncoder(w).Encode(response)
}

func (m *MainController) WriteData(w http.ResponseWriter, r *http.Request) {
	var data models.BasicRequest
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = m.serv.WriteData(data.Data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (m *MainController) CleanData(w http.ResponseWriter, r *http.Request) {
	err := m.serv.CleanData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
