package service

import (
	"context"
	user "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"testing"
)

func TestUpdateUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateUserService(ctx)
	// init req and assert value

	req := &user.UpdateUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
