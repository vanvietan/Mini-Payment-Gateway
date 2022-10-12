package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// CompareOTP compare OTP clients with db
func (i impl) CompareOTP(ctx context.Context, input string) error {
	trans, err := i.txRepo.CompareOTP(ctx, input)
	if err != nil {
		log.Printf("error when compare OTP, error: %v", err)
		return err
	}
	trans.Status = model.StatusAccepted

	return nil
}
