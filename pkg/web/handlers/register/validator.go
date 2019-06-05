package register


type RegisterSchema struct {
	Phone string `json:"phone" binding:"required"`
}

type VerifyDigitalSchema struct {
	Phone string `json:"phone" binding:"required"`
	Digital string `json:"digital" binding:"required"`
}