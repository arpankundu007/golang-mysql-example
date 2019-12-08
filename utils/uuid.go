package utils

import uuid "github.com/satori/go.uuid"

func GetUUID() (string, error){
	uuidVal, err := uuid.NewV4()
	if err!=nil{
		return "", err
	}
	return uuidVal.String(), nil
}
