package jlib

import (
	"encoding/json"
	"errors"
)

type Object = map[string]interface{}
type Array = []interface{}

type JSONObject struct {
	data Object
	err  error
}

type JSONArray struct {
	data Array
	err  error
}

var (
	KeyNotFoundError      = errors.New("Key not found")
	TypeError             = errors.New("invalid value type")
	IndexOutOfLengthError = errors.New("index out of length")
)

func NewObject(i ...interface{}) *JSONObject {
	var data Object
	if len(i) == 1 {
		switch i[0].(type) {
		case string:
			{
				json.Unmarshal([]byte(i[0].(string)), &data)
			}
		case []byte:
			{
				json.Unmarshal(i[0].([]byte), &data)
			}
		case Object:
			{
				data = i[0].(Object)
			}
		}
	}
	return &JSONObject{data: data, err: nil}
}

func NewArray(i ...interface{}) *JSONArray {
	var data []interface{}
	if len(i) == 1 {
		switch i[0].(type) {
		case string:
			{
				json.Unmarshal([]byte(i[0].(string)), &data)
			}
		case []byte:
			{
				json.Unmarshal(i[0].([]byte), &data)
			}
		case Array:
			{
				data = i[0].([]interface{})
			}
		}
	}
	return &JSONArray{data: data, err: nil}
}

func (j *JSONObject) String(key string) string {
	if j.data == nil {
		j.err = KeyNotFoundError
		return ""
	}
	switch obj := j.data[key]; obj.(type) {
	case string:
		j.err = nil
		return j.data[key].(string)
	case nil:
		j.err = KeyNotFoundError
		return ""
	}
	j.err = TypeError
	return ""
}

func (j *JSONObject) ToMap() Object {
	return j.data
}

func (j *JSONObject) Int(key string) int {
	if j.data == nil {
		j.err = KeyNotFoundError
		return 0
	}
	switch obj := j.data[key]; obj.(type) {
	case int:
		j.err = nil
		return obj.(int)
	case float64:
		j.err = nil
		return int(obj.(float64))
	case nil:
		j.err = KeyNotFoundError
		return 0
	}
	j.err = TypeError
	return 0
}

func (j *JSONObject) Float(key string) float64 {
	if j.data == nil {
		j.err = KeyNotFoundError
		return 0
	}
	switch obj := j.data[key]; obj.(type) {
	case int:
		j.err = nil
		return float64(obj.(int))
	case float64:
		j.err = nil
		return j.data[key].(float64)
	}
	j.err = TypeError
	return 0
}

func (j *JSONObject) Bool(key string) bool {
	if j.data == nil {
		j.err = KeyNotFoundError
		return false
	}
	switch obj := j.data[key]; obj.(type) {
	case bool:
		j.err = nil
		return j.data[key].(bool)
	}
	j.err = TypeError
	return false
}

func (j *JSONObject) Object(key string) *JSONObject {
	if j.data == nil {
		j.err = KeyNotFoundError
		return &JSONObject{}
	}
	switch obj := j.data[key]; obj.(type) {
	case Object:
		j.err = nil
		cpy_obj := make(map[string]interface{})
		for k, v := range obj.(Object) {
			cpy_obj[k] = v
		}
		return &JSONObject{
			data: cpy_obj,
		}
	case nil:
		j.err = KeyNotFoundError
		return &JSONObject{}
	}
	j.err = TypeError
	return &JSONObject{}
}

func (j *JSONObject) Array(key string) *JSONArray {
	if j.data == nil {
		j.err = KeyNotFoundError
		return &JSONArray{}
	}
	switch obj := j.data[key]; obj.(type) {
	case Array:
		j.err = nil
		arr := obj.(Array)
		tmp := make(Array, len(arr))
		copy(tmp, arr)
		return &JSONArray{
			data: tmp,
		}
	case nil:
		j.err = KeyNotFoundError
		return &JSONArray{}
	}
	j.err = TypeError
	return &JSONArray{}
}

func (j *JSONObject) Size() int {
	return len(j.data)
}

func (j *JSONObject) Error() error {
	return j.err
}

func (j *JSONObject) Put(key string, v interface{}) *JSONObject {
	if j.data == nil {
		j.data = make(map[string]interface{})
	}
	switch obj := v; obj.(type) {
	case *JSONObject:
		j.data[key] = v.(*JSONObject).ToMap()
		return j
	case *JSONArray:
		j.data[key] = v.(*JSONArray).ToArray()
		return j
	}
	j.data[key] = v
	return j
}

func (j *JSONObject) Delete(key ...string) *JSONObject {
	if j.data != nil {
		for _, k := range key {
			delete(j.data, k)
		}
	}
	return j
}

func (j *JSONObject) ToString() (string, error) {
	b, err := json.Marshal(j.data)
	return string(b), err
}

func (j *JSONObject) ToPrettyString() (string, error) {
	b, err := json.MarshalIndent(j.data, "", "\t")
	return string(b), err
}

func (j *JSONArray) String(index int) string {
	if index < 0 || index >= len(j.data) {
		j.err = IndexOutOfLengthError
		return ""
	}
	switch obj := j.data[index]; obj.(type) {
	case string:
		j.err = nil
		return obj.(string)
	}
	j.err = TypeError
	return ""
}

func (j *JSONArray) ToArray() Array {
	return j.data
}

func (j *JSONArray) Int(index int) int {
	if index < 0 || index >= len(j.data) {
		j.err = IndexOutOfLengthError
		return 0
	}
	switch obj := j.data[index]; obj.(type) {
	case int:
		j.err = nil
		return obj.(int)
	case float64:
		j.err = nil
		return int(obj.(float64))
	}
	j.err = TypeError
	return 0
}

func (j *JSONArray) Float(index int) float64 {
	if index < 0 || index >= len(j.data) {
		j.err = IndexOutOfLengthError
		return 0
	}
	switch obj := j.data[index]; obj.(type) {
	case int:
		j.err = nil
		return float64(obj.(int))
	case float64:
		j.err = nil
		return obj.(float64)
	}
	j.err = TypeError
	return 0
}

func (j *JSONArray) Bool(index int) bool {
	if index < 0 || index >= len(j.data) {
		j.err = IndexOutOfLengthError
		return false
	}
	switch obj := j.data[index]; obj.(type) {
	case bool:
		j.err = nil
		return obj.(bool)
	}
	j.err = TypeError
	return false
}

func (j *JSONArray) Object(index int) *JSONObject {
	if index < 0 || index >= len(j.data) {
		j.err = IndexOutOfLengthError
		return &JSONObject{}
	}
	switch obj := j.data[index]; obj.(type) {
	case Object:
		j.err = nil
		return &JSONObject{
			data: j.data[index].(Object),
		}
	}
	j.err = TypeError
	return &JSONObject{}
}

func (j *JSONArray) Array(index int) *JSONArray {
	if index < 0 || index >= len(j.data) {
		j.err = IndexOutOfLengthError
		return &JSONArray{}
	}
	switch obj := j.data[index]; obj.(type) {
	case Array:
		j.err = nil
		arr := obj.(Array)
		tmp := make(Array, len(arr))
		copy(tmp, arr)
		return &JSONArray{
			data: tmp,
		}
	}
	j.err = TypeError
	return &JSONArray{}
}

func (j *JSONArray) Error() error {
	return j.err
}

func (j *JSONArray) Add(v interface{}) *JSONArray {
	j.data = append(j.data, v)
	return j
}

func (j *JSONArray) DeleteAt(i int) *JSONArray {
	if j.data != nil {
		j.data = append(j.data[:i], j.data[i+1:]...)
	}
	return j
}

func (j *JSONArray) Size() int {
	return len(j.data)
}

func (j *JSONArray) ToString() (string, error) {
	b, err := json.Marshal(j.data)
	return string(b), err
}

func (j *JSONArray) ToPrettyString() (string, error) {
	b, err := json.MarshalIndent(j.data, "", "\t")
	return string(b), err
}
