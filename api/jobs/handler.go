package jobs

import (
	"encoding/json"
	h "github.com/fakorede-bolu/deliva/pkg/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateNewJob(s Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		j := &Jobs{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(j)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		if ok, err := h.ValidateInputs(*j); !ok {
			h.ValidationError(w, http.StatusUnprocessableEntity, err)
			return
		}

		jb, err := s.CreateJob(j)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.RespondJSON(w, http.StatusCreated, jb)

	})
}

func AssignCourierToJob(s Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		j := &JobCourier{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(j)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		if ok, err := h.ValidateInputs(*j); !ok {
			h.ValidationError(w, http.StatusUnprocessableEntity, err)
			return
		}

		ord, err := s.AssignCourierToJob(j.ID, j.CourierID)

		if err != nil {
			h.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.RespondJSON(w, http.StatusCreated, ord)

	})
}


//MakeJobsHandlers make url handlers
func MakeJobsHandlers(r *mux.Router, service Service) {
	r.Handle("/v1/jobs", CreateNewJob(service)).Methods("POST", "OPTIONS").Name("createOrders")
	r.Handle("/v1/jobs/assign-courier", AssignCourierToJob(service)).Methods("POST", "OPTIONS").Name("assignCourierToOrder")
}
