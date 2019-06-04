//go:generate enum -type=UserStatus,UserGender -linecomment=true
package consts

type UserStatus int8

const (
	Enable		UserStatus = 1		// enable
	Disable	UserStatus = 10		// disable
)


type UserGender int8

const (
	Male	UserGender = 0 	// male
	Female	UserGender = 1	// female
	Unknown	UserGender = 2 	// Unknown
)