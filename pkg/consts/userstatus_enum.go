// Code generated by "enum -type=UserStatus,UserGender -linecomment=true"; DO NOT EDIT.

package consts

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

func userStatus() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Enable-1]
	_ = x[Disable-10]
}

// UserStatus Key Value Map
var userStatusKeyMap = map[UserStatus]string{
	Enable:  "enable",
	Disable: "disable",
}

func UserStatusByKey(key string) (UserStatus, error) {
	target := strings.ToLower(key)
	for k, v := range userStatusKeyMap {
		if v == target {
			return k, nil
		}
	}
	return 0, fmt.Errorf("invalid key: %s for UserStatus", key)
}

func (i UserStatus) String() string {
	return userStatusKeyMap[i]
}

// func for Borm
func (i *UserStatus) Value() (interface{}, error) {
	return int(*i), nil
}

func (i *UserStatus) SetValue(v interface{}) error {
	inst := UserStatus(cast.ToInt(v))
	_, ok := userStatusKeyMap[inst]
	if !ok {
		return fmt.Errorf("invalid provider %d", int(inst))
	}
	*i = inst
	return nil
}

// func for JSON
func (i UserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *UserStatus) UnmarshalJSON(data []byte) error {
	var key string
	err := json.Unmarshal(data, &key)
	if err != nil {
		return err
	}

	inst, err := UserStatusByKey(key)
	if err != nil {
		return err
	}

	*i = inst
	return nil
}

func userGender() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Male-0]
	_ = x[Female-1]
	_ = x[Unknown-2]
}

// UserGender Key Value Map
var userGenderKeyMap = map[UserGender]string{
	Male:    "male",
	Female:  "female",
	Unknown: "Unknown",
}

func UserGenderByKey(key string) (UserGender, error) {
	target := strings.ToLower(key)
	for k, v := range userGenderKeyMap {
		if v == target {
			return k, nil
		}
	}
	return 0, fmt.Errorf("invalid key: %s for UserGender", key)
}

func (i UserGender) String() string {
	return userGenderKeyMap[i]
}

// func for Borm
func (i *UserGender) Value() (interface{}, error) {
	return int(*i), nil
}

func (i *UserGender) SetValue(v interface{}) error {
	inst := UserGender(cast.ToInt(v))
	_, ok := userGenderKeyMap[inst]
	if !ok {
		return fmt.Errorf("invalid provider %d", int(inst))
	}
	*i = inst
	return nil
}

// func for JSON
func (i UserGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *UserGender) UnmarshalJSON(data []byte) error {
	var key string
	err := json.Unmarshal(data, &key)
	if err != nil {
		return err
	}

	inst, err := UserGenderByKey(key)
	if err != nil {
		return err
	}

	*i = inst
	return nil
}
