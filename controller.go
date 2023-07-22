package main

import (
	"github.com/gin-gonic/gin"
)

var test = TestSuite{}

type PreAuthRequest struct {
	Id   string `json:"id"`
	Hook string `json:"hook"`
}

func preAuthHandler(c *gin.Context) {
	var body PreAuthRequest
	isTestSuite := c.GetHeader("x-provider-api-is-testsuite") == "true"

	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, gin.H{
			"message": "transactionId is missing",
		})

		return
	}

	if isTestSuite {
		preAuthResponse := test.preAuth(body.Id, body.Hook)
		c.JSON(200, preAuthResponse)

		return
	}

	c.JSON(500, gin.H{
		"message": "Not implemented",
	})
}

func transactions(c *gin.Context) {
	var body PreAuthRequest
	isTestSuite := c.GetHeader("x-provider-api-is-testsuite") == "true"

	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, gin.H{
			"message": "Unexpected body",
		})

		return
	}

	if isTestSuite {
		preAuthResponse := test.transactions(body.Id, body.Hook)
		c.JSON(200, preAuthResponse)

		return
	}

	c.JSON(500, gin.H{
		"message": "Not implemented",
	})
}

func getTransaction(c *gin.Context) {
	isTestSuite := c.GetHeader("x-provider-api-is-testsuite") == "true"
	transactionId := c.Param("id")

	if isTestSuite {
		preAuthResponse := test.getTransaction(transactionId)
		c.JSON(200, preAuthResponse)

		return
	}

	c.JSON(500, gin.H{
		"message": "Not implemented",
	})
}
