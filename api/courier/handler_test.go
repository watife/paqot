package courier

import (
	"bytes"
	h "deliva/pkg/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateNewCourier(t *testing.T) {
	repo := NewInmem()
	service := NewService(repo)
	c := NewFixtureCourier()
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(c)

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}


	ok, errs := h.ValidateInputs(*c)
	assert.NotNil(t, ok)
	assert.Nil(t, errs)

	r := mux.NewRouter()

	MakeCourierHandlers(r, service)

	path, err := r.GetRoute("createCourier").GetPathTemplate()

	assert.Nil(t, err)
	assert.Equal(t, "/v1/courier", path)

	cor, err := service.CreateCourier(c)

	assert.Nil(t, err)

	assert.NotNil(t, cor.ID)
	assert.Equal(t, 8089333186, cor.PhoneNumber)

	ht := CreateNewCourier(service)

	ts := httptest.NewServer(ht)

	defer ts.Close()

	resp, err := http.Post(ts.URL+"/v1/courier", "application/json", reqBodyBytes)

	if err != nil {
		assert.NotEqual(t, "Boluwatife", c.FirstName)
		_ = json.NewDecoder(resp.Body).Decode(&c)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	_ = json.NewDecoder(resp.Body).Decode(&c)
	assert.Equal(t, "Boluwatife", c.FirstName)
}