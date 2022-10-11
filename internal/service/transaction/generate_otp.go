package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
	"pg/internal/util"
)

var getNextIDFunc = util.GetNextId

func (i impl) GenerateOTP(ctx context.Context, cardID int64, orderID int64) (string, error) {
	var inputTX model.Transaction
	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return "", err
	}
	inputTX.ID = ID
	inputTX.CardID = cardID
	inputTX.OrderID = orderID
	inputTX.Status = model.StatusPending

	OTP, errR := i.txRepo.GenerateOTP(ctx, inputTX)
	if errR != nil {
		log.Printf("error when generate OTP %v ", err)
		return "", err
	}
	return OTP, nil
}
