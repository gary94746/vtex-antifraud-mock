package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestSyncApproved(t *testing.T) {
	const transactionId = "transaction1"
	transactionsRecorder := httptest.NewRecorder()
	rawBody := fmt.Sprintf(`{"id": "%s", "hook": ""}`, transactionId)
	transactionsBody := []byte(rawBody)
	transactionsBodyReader := bytes.NewReader(transactionsBody)
	transactionsRequest, _ := http.NewRequest("POST", "/api/v1/anti-fraud/transactions", transactionsBodyReader)
	transactionsRequest.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(transactionsRecorder, transactionsRequest)

	var transactionsResponse map[string]interface{}
	transactionsError := json.NewDecoder(transactionsRecorder.Body).Decode(&transactionsResponse)
	if transactionsError != nil {
		panic(transactionsError)
	}

	assert.Equal(t, 200, transactionsRecorder.Code)
	assert.Equal(t, "received", transactionsResponse["status"])

	attempt1Recorder := httptest.NewRecorder()
	attempt1Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt1Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt1Recorder, attempt1Request)

	var attempt1Response map[string]interface{}
	attempt1ResponseError := json.NewDecoder(attempt1Recorder.Body).Decode(&attempt1Response)
	if attempt1ResponseError != nil {
		log.Fatal("Error attempt1: ", attempt1ResponseError)
	}

	assert.Equal(t, 200, attempt1Recorder.Code)
	assert.Equal(t, "approved", attempt1Response["status"])
}

func TestSyncDenied(t *testing.T) {
	const transactionId = "transaction2"
	transactionsRecorder := httptest.NewRecorder()
	rawBody := fmt.Sprintf(`{"id": "%s", "hook": ""}`, transactionId)
	transactionsBody := []byte(rawBody)
	transactionsBodyReader := bytes.NewReader(transactionsBody)
	transactionsRequest, _ := http.NewRequest("POST", "/api/v1/anti-fraud/transactions", transactionsBodyReader)
	transactionsRequest.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(transactionsRecorder, transactionsRequest)

	var transactionsResponse map[string]interface{}
	transactionsError := json.NewDecoder(transactionsRecorder.Body).Decode(&transactionsResponse)
	if transactionsError != nil {
		panic(transactionsError)
	}

	assert.Equal(t, 200, transactionsRecorder.Code)
	assert.Equal(t, "received", transactionsResponse["status"])

	attempt1Recorder := httptest.NewRecorder()
	attempt1Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt1Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt1Recorder, attempt1Request)

	var attempt1Response map[string]interface{}
	attempt1ResponseError := json.NewDecoder(attempt1Recorder.Body).Decode(&attempt1Response)
	if attempt1ResponseError != nil {
		log.Fatal("Error attempt1: ", attempt1ResponseError)
	}

	assert.Equal(t, 200, attempt1Recorder.Code)
	assert.Equal(t, "denied", attempt1Response["status"])
}

func TestAsyncApproved(t *testing.T) {
	const transactionId = "transaction3"
	transactionsRecorder := httptest.NewRecorder()
	rawBody := fmt.Sprintf(`{"id": "%s", "hook": ""}`, transactionId)
	transactionsBody := []byte(rawBody)
	transactionsBodyReader := bytes.NewReader(transactionsBody)
	transactionsRequest, _ := http.NewRequest("POST", "/api/v1/anti-fraud/transactions", transactionsBodyReader)
	transactionsRequest.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(transactionsRecorder, transactionsRequest)

	var transactionsResponse map[string]interface{}
	transactionsError := json.NewDecoder(transactionsRecorder.Body).Decode(&transactionsResponse)
	if transactionsError != nil {
		panic(transactionsError)
	}

	assert.Equal(t, 200, transactionsRecorder.Code)
	assert.Equal(t, "received", transactionsResponse["status"])

	attempt1Recorder := httptest.NewRecorder()
	attempt1Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt1Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt1Recorder, attempt1Request)

	var attempt1Response map[string]interface{}
	attempt1ResponseError := json.NewDecoder(attempt1Recorder.Body).Decode(&attempt1Response)
	if attempt1ResponseError != nil {
		log.Fatal("Error attempt1: ", attempt1ResponseError)
	}

	assert.Equal(t, 200, attempt1Recorder.Code)
	assert.Equal(t, "undefined", attempt1Response["status"])

	attempt2Recorder := httptest.NewRecorder()
	attempt2Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt2Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt2Recorder, attempt2Request)

	var attempt2Response map[string]interface{}
	attempt2ResponseError := json.NewDecoder(attempt2Recorder.Body).Decode(&attempt2Response)
	if attempt2ResponseError != nil {
		panic(attempt2ResponseError)
	}

	assert.Equal(t, 200, attempt2Recorder.Code)
	assert.Equal(t, "approved", attempt2Response["status"])
}

func TestAsyncDenied(t *testing.T) {
	const transactionId = "transaction4"
	transactionsRecorder := httptest.NewRecorder()
	rawBody := fmt.Sprintf(`{"id": "%s", "hook": ""}`, transactionId)
	transactionsBody := []byte(rawBody)
	transactionsBodyReader := bytes.NewReader(transactionsBody)
	transactionsRequest, _ := http.NewRequest("POST", "/api/v1/anti-fraud/transactions", transactionsBodyReader)
	transactionsRequest.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(transactionsRecorder, transactionsRequest)

	var transactionsResponse map[string]interface{}
	transactionsError := json.NewDecoder(transactionsRecorder.Body).Decode(&transactionsResponse)
	if transactionsError != nil {
		panic(transactionsError)
	}

	assert.Equal(t, 200, transactionsRecorder.Code)
	assert.Equal(t, "received", transactionsResponse["status"])

	attempt1Recorder := httptest.NewRecorder()
	attempt1Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt1Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt1Recorder, attempt1Request)

	var attempt1Response map[string]interface{}
	attempt1ResponseError := json.NewDecoder(attempt1Recorder.Body).Decode(&attempt1Response)
	if attempt1ResponseError != nil {
		log.Fatal("Error attempt1: ", attempt1ResponseError)
	}

	assert.Equal(t, 200, attempt1Recorder.Code)
	assert.Equal(t, "undefined", attempt1Response["status"])

	attempt2Recorder := httptest.NewRecorder()
	attempt2Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt2Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt2Recorder, attempt2Request)

	var attempt2Response map[string]interface{}
	attempt2ResponseError := json.NewDecoder(attempt2Recorder.Body).Decode(&attempt2Response)
	if attempt2ResponseError != nil {
		panic(attempt2ResponseError)
	}

	assert.Equal(t, 200, attempt2Recorder.Code)
	assert.Equal(t, "denied", attempt2Response["status"])
}

func TestHookApproved(t *testing.T) {
	const transactionId = "transaction5"
	transactionsRecorder := httptest.NewRecorder()
	rawBody := fmt.Sprintf(`{"id": "%s", "hook": ""}`, transactionId)
	transactionsBody := []byte(rawBody)
	transactionsBodyReader := bytes.NewReader(transactionsBody)
	transactionsRequest, _ := http.NewRequest("POST", "/api/v1/anti-fraud/transactions", transactionsBodyReader)
	transactionsRequest.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(transactionsRecorder, transactionsRequest)

	var transactionsResponse map[string]interface{}
	transactionsError := json.NewDecoder(transactionsRecorder.Body).Decode(&transactionsResponse)
	if transactionsError != nil {
		panic(transactionsError)
	}

	assert.Equal(t, 200, transactionsRecorder.Code)
	assert.Equal(t, "received", transactionsResponse["status"])

	attempt1Recorder := httptest.NewRecorder()
	attempt1Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt1Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt1Recorder, attempt1Request)

	var attempt1Response map[string]interface{}
	attempt1ResponseError := json.NewDecoder(attempt1Recorder.Body).Decode(&attempt1Response)
	if attempt1ResponseError != nil {
		log.Fatal("Error attempt1: ", attempt1ResponseError)
	}

	assert.Equal(t, 200, attempt1Recorder.Code)
	assert.Equal(t, "undefined", attempt1Response["status"])

	attempt2Recorder := httptest.NewRecorder()
	attempt2Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt2Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt2Recorder, attempt2Request)

	var attempt2Response map[string]interface{}
	attempt2ResponseError := json.NewDecoder(attempt2Recorder.Body).Decode(&attempt2Response)
	if attempt2ResponseError != nil {
		panic(attempt2ResponseError)
	}

	assert.Equal(t, 200, attempt2Recorder.Code)
	assert.Equal(t, "approved", attempt2Response["status"])
}

func TestHookDenied(t *testing.T) {
	const transactionId = "transaction6"
	transactionsRecorder := httptest.NewRecorder()
	rawBody := fmt.Sprintf(`{"id": "%s", "hook": ""}`, transactionId)
	transactionsBody := []byte(rawBody)
	transactionsBodyReader := bytes.NewReader(transactionsBody)
	transactionsRequest, _ := http.NewRequest("POST", "/api/v1/anti-fraud/transactions", transactionsBodyReader)
	transactionsRequest.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(transactionsRecorder, transactionsRequest)

	var transactionsResponse map[string]interface{}
	transactionsError := json.NewDecoder(transactionsRecorder.Body).Decode(&transactionsResponse)
	if transactionsError != nil {
		panic(transactionsError)
	}

	assert.Equal(t, 200, transactionsRecorder.Code)
	assert.Equal(t, "received", transactionsResponse["status"])

	attempt1Recorder := httptest.NewRecorder()
	attempt1Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt1Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt1Recorder, attempt1Request)

	var attempt1Response map[string]interface{}
	attempt1ResponseError := json.NewDecoder(attempt1Recorder.Body).Decode(&attempt1Response)
	if attempt1ResponseError != nil {
		log.Fatal("Error attempt1: ", attempt1ResponseError)
	}

	assert.Equal(t, 200, attempt1Recorder.Code)
	assert.Equal(t, "undefined", attempt1Response["status"])

	attempt2Recorder := httptest.NewRecorder()
	attempt2Request, _ := http.NewRequest("GET", "/api/v1/anti-fraud/transactions/"+transactionId, nil)
	attempt2Request.Header.Add("x-provider-api-is-testsuite", "true")
	router.ServeHTTP(attempt2Recorder, attempt2Request)

	var attempt2Response map[string]interface{}
	attempt2ResponseError := json.NewDecoder(attempt2Recorder.Body).Decode(&attempt2Response)
	if attempt2ResponseError != nil {
		panic(attempt2ResponseError)
	}

	assert.Equal(t, 200, attempt2Recorder.Code)
	assert.Equal(t, "denied", attempt2Response["status"])
}
