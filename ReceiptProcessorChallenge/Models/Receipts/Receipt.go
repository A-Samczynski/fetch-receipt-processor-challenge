package Receipts

import (
	Items "ReceiptProcessorChallenge/Models/Receipts/Items"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Receipt struct {
    Retailer     string  `json:"retailer"`
    PurchaseDate  string  `json:"purchaseDate"`
    PurchaseTime string  `json:"purchaseTime"`
    Items []Items.Item `json:"items"`
	Total string `json:"total"`
}

func (receipt Receipt) CalculatePoints() int{
    totalPoints := 0
    
    totalPoints += receipt.retailerNamePoints(receipt.Retailer)
    totalPoints += receipt.roundDollarPoints(receipt.Total)
    totalPoints += receipt.quarterMultiplePoints(receipt.Total)
    totalPoints += receipt.itemPairsPoints(receipt.Items)
    totalPoints += receipt.trimmedLengthPoints(receipt.Items, receipt.Total)
    totalPoints += receipt.oddDayPoints(receipt.PurchaseDate)
    totalPoints += receipt.purchaseTimePoints(receipt.PurchaseTime)

    return totalPoints
}
func (receipt Receipt) retailerNamePoints(retailer string) int{
    //any character that is not a thru z, A thru Z, or 0-9 --> [^a-z0-9][^A-Z]
    regularExpression := regexp.MustCompile(`[^a-zA-Z0-9]+`)

    retailer = regularExpression.ReplaceAllString(retailer, "")
    points := 0
    points += len(retailer)
    return points
}
func (receipt Receipt) roundDollarPoints(total string) int{
    points := 0
    if (strings.Contains(total, ".00")){
        points += 50
    }
    return points
}
func (receipt Receipt) quarterMultiplePoints(total string) int{
    points := 0
    if (strings.Contains(total, ".00") || strings.Contains(total, ".25") ||
        strings.Contains(total, ".50") || strings.Contains(total, ".75")){
        points += 25
    }
    return points
}
func (receipt Receipt) itemPairsPoints(items []Items.Item) int{
    points := 0
    points += 5*(len(items)/2) //automatically performs floor division for integers
    return points
}
func (receipt Receipt) trimmedLengthPoints(items []Items.Item, totalString string) int{
    points := 0
    
    for i:= 0; i < len(items); i++{
        itemDescription := strings.Trim(items[i].ShortDescription, " ")
        
        if ((len(itemDescription)) % 3 == 0){
            price, err := strconv.ParseFloat(items[i].Price, 64)
            if err != nil {
                panic(err)
            }
            result := price * float64(0.2)
            points += int(math.Ceil(result))
        }
    }
    return points
}
func (receipt Receipt) oddDayPoints(purchaseDate string) int{
    points := 0
    purchaseDateArray := strings.Split(purchaseDate, "-")
    purchaseDay, err := strconv.Atoi(purchaseDateArray[2])
    if err != nil {
        panic(err)
    }
    if (purchaseDay%2 == 1){
        points += 6
    }  
    return points
}
func (receipt Receipt) purchaseTimePoints(purchaseTime string) int{
    points := 0
    purchaseTimeArray := strings.Split(purchaseTime, ":")
    purchaseHour, err := strconv.Atoi(purchaseTimeArray[0])
    if err != nil {
        panic(err)
    }
    if (purchaseHour >= 14 && purchaseHour <= 16){
        points += 10
    } 
    return points
}



type ReceiptID struct{
	ID string `json:"id"`
}

func (receiptID *ReceiptID) SetID(shaID string){
    receiptID.ID = shaID
}
func (receiptID ReceiptID) GetID() string{
    return receiptID.ID
}
func (receiptID *ReceiptID) GenerateID(receipt Receipt) string{    
    var temp_buffer bytes.Buffer
    enc := gob.NewEncoder(&temp_buffer) // Will write to temp_buffer.

    err := enc.Encode(receipt)
    if err != nil {
        log.Fatal("Error during Encode():", err)
    }

    hasher := sha256.New()
    hasher.Write(temp_buffer.Bytes())
    sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
    receiptID.SetID(sha)
    return sha
}

type ReceiptPoints struct{
	Points string `json:"points"`
}
func (receiptPoints *ReceiptPoints) SetPoints(points int){
    pointString := strconv.Itoa(points)
    receiptPoints.Points = pointString
}
func (receiptPoints ReceiptPoints) GetPoints() string{
    return receiptPoints.Points
}







