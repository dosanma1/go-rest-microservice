package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dosanma1/go-rest-microservice/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	router, err := setup()

	w := performRequest(router, http.MethodGet, "/health")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["status"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "success", value)
}

func TestGetProduct_Ok(t *testing.T) {
	router, err := setup()

	product := model.Product{
		Name:     "sun cream 15",
		Quantity: 240,
	}

	w := performBodyRequest(router, http.MethodPost, "/api/v1/product/", product)

	var createResponse model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &createResponse)

	w = performRequest(router, http.MethodGet, fmt.Sprintf("/api/v1/product/%s", createResponse.Response.ProductID))

	assert.Equal(t, http.StatusOK, w.Code)

	var getResponse model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &getResponse)

	assert.Nil(t, err)
	assert.Equal(t, "success", getResponse.Status)
	assert.Equal(t, createResponse.Response.ProductID, getResponse.Response.ProductID)
}

func TestCreateProduct_Ok(t *testing.T) {
	router, err := setup()

	product := model.Product{
		Name:     "sun cream 15",
		Quantity: 24,
	}

	w := performBodyRequest(router, http.MethodPost, "/api/v1/product/", product)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, "success", response.Status)
	assert.Equal(t, product.Name, response.Response.Name)
}

func TestUpdateProduct_Ok(t *testing.T) {
	router, err := setup()

	product := model.Product{
		Name:     "sun cream 50",
		Quantity: 23,
	}

	w := performBodyRequest(router, http.MethodPost, "/api/v1/product/", product)

	var createResponse model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &createResponse)

	createResponse.Response.Quantity--

	w = performBodyRequest(router, http.MethodPut, "/api/v1/product/", createResponse.Response)

	assert.Equal(t, http.StatusOK, w.Code)

	var updateResponse model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &updateResponse)

	assert.Nil(t, err)
	assert.Equal(t, "success", updateResponse.Status)
	assert.Equal(t, createResponse.Response.Quantity, updateResponse.Response.Quantity)
}

func TestDeleteProduct_Ok(t *testing.T) {
	router, err := setup()

	product := model.Product{
		Name:     "sun cream 30",
		Quantity: 240,
	}

	w := performBodyRequest(router, http.MethodPost, "/api/v1/product/", product)

	var createResponse model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &createResponse)

	w = performRequest(router, http.MethodDelete, fmt.Sprintf("/api/v1/product/%s", createResponse.Response.ProductID))

	assert.Equal(t, http.StatusOK, w.Code)

	var deleteResponse model.ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &deleteResponse)

	assert.Nil(t, err)
	assert.Equal(t, "success", deleteResponse.Status)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	r.ServeHTTP(w, req)
	return w
}

func performBodyRequest(r http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	postBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(postBody))
	r.ServeHTTP(w, req)
	return w
}
