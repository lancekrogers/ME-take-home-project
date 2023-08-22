package utils

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeStructForDB(t *testing.T) {
	value := struct {
		Name string
		Age  int
	}{"John", 30}

	result, err := EncodeStructForDB(value)
	assert.NoError(t, err)
	assert.True(t, result.Valid)

	var decoded map[string]interface{}
	json.Unmarshal(result.RawMessage, &decoded)
	assert.Equal(t, "John", decoded["Name"])
	assert.Equal(t, float64(30), decoded["Age"])
}

func TestDecodeJson(t *testing.T) {
	raw := json.RawMessage(`{"name": "John", "age": 30}`)
	result, err := DecodeJson(raw)
	assert.NoError(t, err)
	assert.Equal(t, "John", result["name"])
	assert.Equal(t, float64(30), result["age"])
}

func TestMapToStruct(t *testing.T) {
	dataMap := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	result, err := MapToStruct(dataMap)
	assert.NoError(t, err)

	// Verify the result using reflection
	value := reflect.ValueOf(result)
	assert.Equal(t, "John", value.FieldByName("Name").Interface())
	assert.Equal(t, 30, value.FieldByName("Age").Interface())
}
