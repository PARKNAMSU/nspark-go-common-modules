package common

import "encoding/json"

func CheckError(err error) bool {
	return err != nil
}

func ObjectToJsonBin[T interface{}](obj T) (*[]byte, error) {
	jsonBin, err := json.Marshal(obj)

	if CheckError(err) {
		return nil, err
	}

	return &jsonBin, err
}

func JsonBinTOObject[T interface{}](bin []byte) (*T, error) {
	var newObj T

	err := json.Unmarshal(bin, newObj)

	if CheckError(err) {
		return nil, err
	}

	return &newObj, nil
}
