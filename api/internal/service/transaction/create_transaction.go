package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math/rand"
	model2 "pg/api/internal/model"
	"pg/api/internal/util"
	"strconv"
	"time"
)

const (
	min = 10
	max = 999999
)

var getNextIDFunc = util.GetNextId
var randomFunc = randInt

// CreateTransaction create a transaction
func (i impl) CreateTransaction(ctx context.Context, cardID int64, orderID int64) (model2.Transaction, error) {
	var inputTX model2.Transaction

	//generate ID for transaction input
	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return model2.Transaction{}, err
	}
	inputTX.ID = ID
	inputTX.CardID = cardID
	inputTX.OrderID = orderID
	inputTX.Status = model2.StatusPending

	//generate otp
	randNumber := randomFunc(min, max)
	inputOTP := strconv.Itoa(randNumber)
	inputTX.OTP = inputOTP

	trans, errR := i.txRepo.CreateTransaction(ctx, inputTX)
	if errR != nil {
		log.Printf("error when generate OTP %v ", err)
		return model2.Transaction{}, err
	}
	return trans, nil
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
