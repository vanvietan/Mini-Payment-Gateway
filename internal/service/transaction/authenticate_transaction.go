package transaction

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// AuthenticateTransaction authenticate transaction
func (i impl) AuthenticateTransaction(ctx context.Context, id int64, otp string) error {
	/* find and compare OTP between
	client and transaction created
	*/
	trans, err := i.txRepo.FindTransactionByID(ctx, id)
	if err != nil {
		log.Printf("error when find id transaction, error: %v", err)
		return err
	}
	if trans.OTP != otp {
		return errors.New("wrong otp")
	}
	/* update transaction status
	 */
	trans.Status = model.StatusAccepted
	_, errU := i.txRepo.UpdateTransaction(ctx, trans)
	if errU != nil {
		log.Printf("error when update transaction status, error: %v", errU)
		return err
	}
	return nil
}
