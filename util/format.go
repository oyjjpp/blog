package util

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

//@desc interface转换为整形
func Int(data interface{}) (int, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(data).Uint()), nil
	case string:
		return strconv.Atoi(data.(string))
	}
	return 0, errors.New("invalid value type")
}

//@desc interface 转换为浮点
func Float(data interface{}) (float64, error) {
	switch data.(type) {
	case json.Number:
		return data.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(data).Uint()), nil
	case string:
		return strconv.ParseFloat(data.(string), 64)
	}
	return 0, errors.New("invalid value type")
}

//转换为字符串
func String(data interface{}) (string, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(i, 10), nil
	case float32, float64:
		fdata := reflect.ValueOf(data).Float()
		return strconv.FormatFloat(fdata, 'f', -1, 64), nil
	case int, int8, int16, int32, int64:
		idata := reflect.ValueOf(data).Int()
		return strconv.FormatInt(idata, 10), nil
	case uint, uint8, uint16, uint32, uint64:
		udata := reflect.ValueOf(data).Uint()
		return strconv.FormatUint(udata, 10), nil
	case string:
		return data.(string), nil
	}
	return "", errors.New("invalid value type")
}
