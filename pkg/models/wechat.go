package models

type WeChatInfo struct {
	ID       uint   `gorm:"primary_key"`
	UnionID  string `gorm:"unionid"`
	OpenID   string `gorm:"openid"`
	Name     string `gorm:"name"`
	Gender   string `gorm:"gender"`
	Province string `gorm:"province"`
	City     string `gorm:"city"`
	Country  string `gorm:"country"`
	Avatar   string `gorm:"avatar"`
	DateTimeModel
}

func (w *WeChatInfo) TableName() string {
	return "wechat_info"
}

type UserWeChatShip struct {
	ID      uint   `gorm:"primary_key"`
	UID     string `gorm:"uid"`
	UnionID string `gorm:"unionid"`
	DateTimeModel
}

func (w *UserWeChatShip) TableName() string {
	return "user_wechat_ship"
}
