package http

import (
	"encoding/json"
	"net/http"
)

type Agentrequest struct {
	Promt  string `json:"promt"`
	APIKey string `json:"api_key"`
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandlePrompt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) //405
		w.Write([]byte("Разрешение только POST!"))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 5*1024*1024)

	var req Agentrequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка: сломанный json-пакет:", err.Error()))
		return
	}

	if req.Promt == "" || req.APIKey == "" {
		w.WriteHeader((http.StatusUnprocessableEntity))
		w.Write([]byte("Ошибка: Поля promt и api_key пустые"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{status":"queued"}`))
}

func main() {

}
