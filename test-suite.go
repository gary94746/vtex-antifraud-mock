package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	ApprovedStatus  = "approved"
	ReceivedStatus  = "received"
	DeniedStatus    = "denied"
	UndefinedStatus = "undefined"
)

type AntiFraudResponse struct {
	AnalysisType string `json:"analysisType"`
	Code         string `json:"code"`
	ID           string `json:"id"`
	Message      string `json:"message"`
	Responses    struct {
	} `json:"responses"`
	Score  int    `json:"score"`
	Tid    string `json:"tid"`
	Status string `json:"status"`
}

type TestSuite struct {
}

var memCache = make(map[string]int8)

func (test *TestSuite) preAuth(transactionId string, hookUrl string) AntiFraudResponse {
	return test.getCase(transactionId, hookUrl)
}

func (test *TestSuite) transactions(transactionId string, hookUrl string) AntiFraudResponse {
	go test.pingUrl(hookUrl, transactionId)

	return AntiFraudResponse{
		AnalysisType: "sync",
		Code:         "ok",
		ID:           transactionId,
		Message:      "Testing",
		Responses:    struct{}{},
		Score:        100,
		Tid:          transactionId,
		Status:       ReceivedStatus,
	}
}

func (test *TestSuite) getTransaction(transactionId string) AntiFraudResponse {
	return test.getCase(transactionId, "")
}

func (test *TestSuite) getResponse(status string, transactionId string) AntiFraudResponse {
	isApproved := status == ApprovedStatus
	isDeclined := status == DeniedStatus

	if isApproved {
		return AntiFraudResponse{
			AnalysisType: "sync",
			Code:         "ok",
			ID:           transactionId,
			Message:      "Testing",
			Responses:    struct{}{},
			Score:        100,
			Tid:          transactionId,
			Status:       ApprovedStatus,
		}
	}

	if isDeclined {
		return AntiFraudResponse{
			AnalysisType: "sync",
			Code:         "ok",
			ID:           transactionId,
			Message:      "Testing",
			Responses:    struct{}{},
			Score:        100,
			Tid:          transactionId,
			Status:       DeniedStatus,
		}
	}

	return AntiFraudResponse{
		AnalysisType: "sync",
		Code:         "ok",
		ID:           transactionId,
		Message:      "Testing",
		Responses:    struct{}{},
		Score:        100,
		Tid:          transactionId,
		Status:       UndefinedStatus,
	}
}

func (test *TestSuite) getCase(transactionId string, hookUrl string) AntiFraudResponse {
	cases := make(map[string]AntiFraudResponse)
	lastCharacter := transactionId[len(transactionId)-1:]

	approvedResponse := test.getResponse(ApprovedStatus, transactionId)
	deniedResponse := test.getResponse(DeniedStatus, transactionId)

	cases["1"] = approvedResponse
	cases["2"] = deniedResponse
	if lastCharacter == "3" {
		cases["3"] = test.asyncResponse(transactionId)
	}
	if lastCharacter == "4" {
		cases["4"] = test.asyncResponse(transactionId)
	}
	if lastCharacter == "5" {
		cases["5"] = test.asyncResponse(transactionId)
	}
	if lastCharacter == "6" {
		cases["6"] = test.asyncResponse(transactionId)
	}

	shouldPing := lastCharacter == "5" || lastCharacter == "6"
	if shouldPing {
		go test.pingUrl(hookUrl, transactionId)
	}

	return cases[lastCharacter]
}

func (test *TestSuite) asyncResponse(transactionId string) AntiFraudResponse {
	isNotInCache := memCache[transactionId] == 0
	if isNotInCache {
		memCache[transactionId] = 1

		return test.getResponse(UndefinedStatus, transactionId)
	}

	lastCharacter := transactionId[len(transactionId)-1:]
	shouldBeApprove := lastCharacter == "3" || lastCharacter == "5"

	if shouldBeApprove {
		return test.getResponse(ApprovedStatus, transactionId)
	}

	return test.getResponse(DeniedStatus, transactionId)
}

func (test *TestSuite) pingUrl(url string, transactionId string) {
	shouldBeAvoided := url == ""
	if shouldBeAvoided {
		return
	}

	time.Sleep(time.Millisecond * 500)

	transactionStatus := test.getResponse(ApprovedStatus, transactionId)
	jsonBytes, err := json.Marshal(transactionStatus)
	if err != nil {
		panic(err)
	}

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Request done!", res.StatusCode)
}
