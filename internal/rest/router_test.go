package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestValidatorHandlerWithValidCard(t *testing.T) {
	logger := logrus.New()
	router := getRouter(logger)

	request, _ := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(`{"card_number":"5395577169979803", "expiration_month":12, "expiration_year":9999}`)))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "card valid", response.Body.String())
}

func TestValidatorHandlerWithInvalidCard(t *testing.T) {
	logger := logrus.New()
	router := getRouter(logger)

	request, _ := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(`{"card_number":"123456781234567", "expiration_month":12, "expiration_year":2023}`)))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestValidatorHandlerWithExpiredCard(t *testing.T) {
	logger := logrus.New()
	router := getRouter(logger)

	request, _ := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(`{"card_number":"1234567812345678", "expiration_month":1, "expiration_year":2020}`)))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}
