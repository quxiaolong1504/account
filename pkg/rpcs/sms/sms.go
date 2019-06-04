package sms

import "github.com/quxiaolong/account/pkg/utils/logger"

func SendSms(contact, content string) {
	logger.Logger.Infof("[sms] `%s` -> `%s`", content, contact)
}