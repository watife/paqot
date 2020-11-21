package courier

import (
	"encoding/json"
	h "github.com/fakorede-bolu/deliva/pkg/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateNewCourier(s Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := &Courier{}
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

		cor, err := s.CreateCourier(c)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.RespondJSON(w, http.StatusCreated, cor)

	})
}


//MakeCourierHandlers make url handlers
func MakeCourierHandlers(r *mux.Router, service Service) {
	r.Handle("/v1/courier", CreateNewCourier(service)).Methods("POST", "OPTIONS").Name("createCourier")
}