package tests

import (
	"bytes"
	"encoding/json"
	"github.com/fakorede-bolu/deliva/api/courier"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateNewCourier(t *testing.T) {
	repo := newInmem()
	service := courier.NewService(repo)
	controller := courier.CreateNewCourier(service)
	//create a new HTTP POST REQUEST
	jsonReq := []byte(`{
			"phoneNumber": 8089333186,
			"DOB":"1996-03-17",
			"address": "eleyele ibadan",
			"countryCode": "+234",
			"firstName": "boluwatife",
			"lastName": "fakorede",
			
		}`)
	req, _ := http.NewRequest("POST", "/v1/courier", bytes.NewBuffer(jsonReq))
	//Assign HTTP Handler function (controller AddCourier function)
	handler := http.Handler(controller)

	//Record HTTP response
	res := httptest.NewRecorder()

	//Dispatch HTTP request
	handler.ServeHTTP(res, req)

	//Add assertions on the HTTP status code and the response
	status := res.Code

	if status != http.StatusCreated {
		t.Errorf("Handlers returned a wrong status code: got %v want %v", status, http.StatusCreated)
	}

	assert.Equal(t, status, http.StatusCreated)

	var cor courier.Courier

	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&cor)

	if err != nil {
		assert.Nil(t, err)
	}

	assert.NotNil(t, cor.ID)
	assert.Equal(t, 8089333186, cor.PhoneNumber)
}