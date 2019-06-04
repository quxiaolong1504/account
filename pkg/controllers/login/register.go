package login

import (
	"context"
	"github.com/quxiaolong/account/pkg/rpcs/sms"
)

func SendVerifyCode(ctx context.Context, phone string) error {
	sms.SendSms(phone, "sms")
	return nil
}