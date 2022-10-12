package transaction

import (
	"context"
	"math/rand"
	"pg/internal/model"
	"strconv"
)

const (
	min = 10
	max = 999999
)

// GenerateOTP generate OTP and create a authentication
func (i impl) GenerateOTP(ctx context.Context, transaction model.Transaction) (string, error) {
	randNumber := randInt(min, max)
	OTP := strconv.Itoa(randNumber)
	transaction.OTP = OTP

	tx := i.gormDB.Create(&transaction)
	if tx.Error != nil {
		return "", tx.Error
	}
	return OTP, nil
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
