package service

import (
	"context"
	"testing"
	payment "github.com/MelodyDeep/TikTok-E-commerce/app/payment/kitex_gen/rpc/payment"
)

func TestHandlePaymentTimeout_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHandlePaymentTimeoutService(ctx)
	// init req and assert value

	req := &payment.PaymentTimeoutRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
