package customer

import (
	"encoding/json"
	h "github.com/fakorede-bolu/deliva/pkg/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateNewCustomer(s Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := &Customer{}
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


//MakeCustomerHandlers make url handlers
func MakeCustomerHandlers(r *mux.Router, service Service) {
	r.Handle("/v1/customer", CreateNewCustomer(service)).Methods("POST", "OPTIONS").Name("createCustomer")
}
