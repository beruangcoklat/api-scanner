package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/beruangcoklat/api-scanner/domain"
	"github.com/gorilla/mux"
)

type apiDataHandler struct {
	apiDataUsecase domain.APIDataUsecase
}

func NewAPIDataHandler(router *mux.Router, apiDataUsecase domain.APIDataUsecase) {
	handler := &apiDataHandler{
		apiDataUsecase: apiDataUsecase,
	}
	router.HandleFunc("/api-data", handler.create).Methods(http.MethodPost)
	router.HandleFunc("/api-data", handler.get).Methods(http.MethodGet)
	router.HandleFunc("/api-data/{id}", handler.getByID).Methods(http.MethodGet)
	router.HandleFunc("/api-data/{id}/scan", handler.scan).Methods(http.MethodGet)
	router.HandleFunc("/api-data/{id}/is-scan-running", handler.isScanRunning).Methods(http.MethodGet)
}

func (h *apiDataHandler) create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}

	var param domain.APIData
	err = json.Unmarshal(body, &param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}

	param.ScanResult = []domain.APIDataScanResult{}
	err = h.apiDataUsecase.Create(r.Context(), param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}
}

func (h *apiDataHandler) get(w http.ResponseWriter, r *http.Request) {
	result, err := h.apiDataUsecase.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *apiDataHandler) getByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.apiDataUsecase.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *apiDataHandler) scan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.apiDataUsecase.PublishScanMessage(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}
}

func (h *apiDataHandler) isScanRunning(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.apiDataUsecase.IsScanRunning(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type ResponseError struct {
	Message string `json:"message"`
}
