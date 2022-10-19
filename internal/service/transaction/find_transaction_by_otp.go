package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// FindTransactionByOTP compare OTP clients with db
func (i impl) FindTransactionByOTP(ctx context.Context, input string) (model.Transaction, error) {
	/* find and compare OTP between
	client and transaction created
	*/
	trans, err := i.txRepo.FindTransactionByOTP(ctx, input)
	if err != nil {
		log.Printf("error when compare OTP, error: %v", err)
		return model.Transaction{}, err
	}
	/* update transaction status
	 */
	trans.Status = model.StatusAccepted
	tx, errU := i.txRepo.UpdateTransaction(ctx, trans)
	if errU != nil {
		log.Printf("error when update transaction status, error: %v", errU)
		return model.Transaction{}, err
	}

	return tx, nil
}
