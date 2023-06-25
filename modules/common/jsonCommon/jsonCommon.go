package jsonCommon

import (
	"encoding/json"
)

func ObjectToJsonBin[T interface{}](obj T) (*[]byte, error) {
	jsonBin, err := json.Marshal(obj)

	if err != nil {
		return nil, err
	}

	return &jsonBin, err
}

func JsonBinTOObject[T interface{}](bin []byte) (*T, error) {
	var newObj T

	err := json.Unmarshal(bin, newObj)

	if err != nil {
		return nil, err
	}

	return &newObj, nil
}

func ObjectToModel[T interface{}](obj interface{}, model *T) error {
	bin, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bin, model)

	if err != nil {
		return err
	}
	return nil
}
