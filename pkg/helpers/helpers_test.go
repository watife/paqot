package helpers

//import (
//	"errors"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestRespondJSON(t *testing.T) {
//	httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		RespondJSON(w, http.StatusOK, 1)
//	}))
//}
//func TestRespondError(t *testing.T) {
//	e := errors.New("bad request")
//	httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		RespondError(w, http.StatusBadRequest, e.Error())
//	}))
//}
