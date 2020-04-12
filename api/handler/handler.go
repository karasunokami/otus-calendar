package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/karasunokami/otus-calendar/internal/domain"
	"github.com/karasunokami/otus-calendar/internal/service"
	"github.com/karasunokami/otus-calendar/internal/version"
	"github.com/sirupsen/logrus"
)

// Handler for processing API endpoints
type Handler struct {
	Calendar *service.CalendarService
	Storage  domain.StorageInterface
	Logger   *logrus.Logger
}

// APIResponse Function for marshaling and sending data
type APIResponse func(resp http.ResponseWriter)

// Handle Wrap handler function for http server
func (h *Handler) Handle(fn func(req *http.Request) APIResponse) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Server", "GoCalendar API "+version.Version)
		fn(req)(resp)
	}
}

// JSON Send data in JSON format (success)
func (h *Handler) JSON(code int, data interface{}) APIResponse {
	return h.sendJSON(code, map[string]interface{}{"result": data})
}

// Error Send data in JSON format (error)
func (h *Handler) Error(code int, err error) APIResponse {
	return h.sendJSON(code, map[string]string{"error": err.Error()})
}

func (h *Handler) sendJSON(code int, data interface{}) APIResponse {
	var encodedData []byte
	var err error

	if data != nil {
		encodedData, err = json.Marshal(data)
		if err != nil {
			return h.Error(http.StatusInternalServerError, err)
		}
	}

	return func(resp http.ResponseWriter) {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.Header().Set("Content-Length", strconv.Itoa(len(encodedData)))
		resp.WriteHeader(code)
		_, _ = resp.Write(encodedData)
	}
}
