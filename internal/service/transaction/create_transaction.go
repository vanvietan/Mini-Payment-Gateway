package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"pg/internal/model"
	"pg/internal/util"
	"strconv"
)

const (
	min = 10
	max = 999999
)

var getNextIDFunc = util.GetNextId
var randomFunc = randInt

// CreateTransaction create a transaction
func (i impl) CreateTransaction(ctx context.Context, cardID int64, orderID int64) (model.Transaction, error) {
	var inputTX model.Transaction

	//generate ID for transaction input
	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return model.Transaction{}, err
	}
	inputTX.ID = ID
	inputTX.CardID = cardID
	inputTX.OrderID = orderID
	inputTX.Status = model.StatusPending

	//generate otp
	randNumber := randomFunc(min, max)
	inputOTP := strconv.Itoa(randNumber)
	inputTX.OTP = inputOTP

	trans, errR := i.txRepo.CreateTransaction(ctx, inputTX)
	if errR != nil {
		log.Printf("error when generate OTP %v ", err)
		return model.Transaction{}, err
	}
	return trans, nil
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
