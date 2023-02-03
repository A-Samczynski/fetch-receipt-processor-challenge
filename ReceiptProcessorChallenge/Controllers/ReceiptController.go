package Controllers

import (
	Receipts "ReceiptProcessorChallenge/Models/Receipts"

	"github.com/gin-gonic/gin"
)

//this map will act as a database
var	receiptDatabase = make(map[string]Receipts.Receipt)
var pointsDatabase = make(map[string] Receipts.ReceiptPoints)


func PostReceiptProcess(c *gin.Context) {
        var receipt Receipts.Receipt
		var receiptID Receipts.ReceiptID
        var receiptPoints Receipts.ReceiptPoints

        c.BindJSON(&receipt)
		receiptID.GenerateID(receipt)
        receiptPoints.SetPoints(receipt.CalculatePoints())
		receiptDatabase[receiptID.GetID()] = receipt
        pointsDatabase[receiptID.GetID()] = receiptPoints
		c.IndentedJSON(200, receiptID)
    }

func GetReceiptPoints(c *gin.Context){
    id := c.Param("id")
    c.IndentedJSON(200, pointsDatabase[id])


    
}





     


