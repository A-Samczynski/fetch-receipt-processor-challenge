package main

import (
	ReceiptController "ReceiptProcessorChallenge/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    
	router.POST("/receipts/process", ReceiptController.PostReceiptProcess) 

	router.GET("/receipts/:id/points", ReceiptController.GetReceiptPoints)

    router.Run("localhost:8080")
}