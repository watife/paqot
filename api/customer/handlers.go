package customer

import (
	"deliva/api/entities"
	h "deliva/pkg/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateNewCustomer(s Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := &entities.Customer{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(c)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		if ok, err := h.ValidateInputs(*c); !ok {
			h.ValidationError(w, http.StatusUnprocessableEntity, err)
			return
		}

		cor, err := s.CreateCustomer(c)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.RespondJSON(w, http.StatusCreated, cor)

	})
}

func FindCustomerByID(s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sId := mux.Vars(r)["id"]

		id, err := entities.StringToID(sId)
		if err != nil {
			h.RespondError(w, http.StatusInternalServerError, err.Error())
		}

		cor, err := s.FindCustomerByID(id)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.RespondJSON(w, http.StatusOK, cor)

	})
}


//MakeCustomerHandlers make url handlers
func MakeCustomerHandlers(r *mux.Router, service Service) {
	r.Handle("/v1/customer", CreateNewCustomer(service)).Methods("POST", "OPTIONS").Name("createCustomer")
	r.Handle("/v1/customer/{id}", FindCustomerByID(service)).Methods("GET", "OPTIONS").Name("getCustomer")
}
