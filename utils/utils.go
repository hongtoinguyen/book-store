package utils

import (
	"os"
	"strconv"
)

//GetEnv get key environment variable if exist otherwise return defalutValue
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

//ConvertStringToInt ...
func ConvertStringToInt(str string) (int, error) {
	res, err := strconv.Atoi(str)
	return res, err
}

func InterfaceToInt(val interface{}) int {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

func InterfaceToInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int:
		return int64(v)
	case int64:
		return v
	case float64:
		return int64(v)
	default:
		return 0
	}
}

func InterfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}
	switch v := val.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func InterfaceToBytes(val interface{}) []byte {
	var x []byte
	if val == nil {
		return x
	}
	switch v := val.(type) {
	case []byte:
		return v
	default:
		return x
	}
}

func MaxInt64(v1, v2 int64) int64 {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
