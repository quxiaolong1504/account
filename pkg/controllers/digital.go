package digital

import (
	"fmt"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/rpcs/sms"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"math/rand"
	"time"
)

var CanSendDigitalTimeThreshold = time.Second * 30 + time.Minute * 4
/*
向用户发送注册验证码
	发送频率限制： 60s/次
	有效期: 5m
 */
func SendRegDigital(phone string) error {
	redisKey := fmt.Sprintf("phone:%s:reg:digital", phone)
	if err := InspectKey(redisKey); err != nil {
		return err
	}

	digital := GenDigital()

	if err := sms.SendSms(phone, GenSmsContent(digital)); err != nil {
		return err
	}
	storage.Cache.Set(redisKey, digital, time.Minute * 5)
	return nil
}

func VerifyDigital(phone, digital string) bool {
	redisKey := fmt.Sprintf("phone:%s:reg:digital", phone)
	dig, err := storage.Cache.Get(redisKey).Result()
	if  err != nil {
		return false
	}

	defer func(){
		// 仅用一次
		if digital == dig{
			storage.Cache.Del(redisKey)
		}
	}()
	return digital == dig
}

func GenDigital() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999 - 1000) + 1000
}

func GenSmsContent(digital int) string{
	return fmt.Sprintf(`您的验证码是:%d, 有效期 5分钟.`, digital)
}

func InspectKey(redisKey string) error {
	isExists, err := storage.Cache.Exists(redisKey).Result()
	if  err != nil {
		panic(err.Error())
	}

	// 如果存在则 检查频率
	if isExists == 1 {
		ttl, err := storage.Cache.TTL(redisKey).Result()
		if err != nil {
			panic(err.Error())
		}
		// 有效期大于 4 分 半 则认为可以频率太快了
		if ttl >= CanSendDigitalTimeThreshold{
			after := ttl - CanSendDigitalTimeThreshold
			return errs.BaseError{Reason: fmt.Sprintf("发送频率过快, 请在 %s 秒后再试!", after.String() )}
		}
	}
	return nil
}