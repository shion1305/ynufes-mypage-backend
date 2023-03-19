package typecast

import (
	"errors"
	"fmt"
)

func ConvertToStringMapInterface(v interface{}) (map[string]interface{}, error) {
	m, ok := v.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("ConvertToStringMapInterface: %v is not a map", v)
	}
	res := make(map[string]interface{})
	for k, v := range m {
		key, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("ConvertToStringMapInterface: map key %v is not a string", k)
		}
		res[key] = v
	}
	return res, nil
}

func ConvertToStringMapString(v interface{}) (map[string]string, error) {
	m, ok := v.(map[string]string)
	if ok {
		return m, nil
	}
	m1, ok := v.(map[string]interface{})
	if !ok {
		return nil, errors.New("cannot convert to map[string]string")
	}
	m = make(map[string]string, len(m1))
	for k, v := range m1 {
		s, ok := v.(string)
		if !ok {
			return nil, errors.New("cannot convert to map[string]string")
		}
		m[k] = s
	}
	return m, nil
}

func ConvertToStringMapFloat64(v interface{}) (map[string]float64, error) {
	m, ok := v.(map[string]float64)
	if ok {
		return m, nil
	}
	m1, ok := v.(map[string]interface{})
	if !ok {
		return nil, errors.New("cannot convert to map[string]interface{}")
	}
	m = make(map[string]float64, len(m1))
	for k, v := range m1 {
		f, ok := v.(float64)
		if !ok {
			return nil, errors.New("cannot convert to map[string]float64")
		}
		m[k] = f
	}
	return m, nil
}
