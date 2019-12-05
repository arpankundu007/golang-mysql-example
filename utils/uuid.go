package utils

import uuid "github.com/satori/go.uuid"

func GetUUID() string{
	uuidVal, err := uuid.NewV4()
	if err!=nil{
		panic(err.Error())
		return ""
	}
	return uuidVal.String()
}
