package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/sqlc-dev/pqtype"
)

func LoadConfig[T interface{}](path string) (*T, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c *T
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot decode into struct")
	}
	return c, nil
}

func PrintStructFieldTypes(value interface{}) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		fmt.Println("Value is not a struct")
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i).Type
		fieldName := t.Field(i).Name
		fmt.Printf("Field: %s, Type: %s, Value: %v\n", fieldName, fieldType, field.Interface())
	}
}

func EncodeStructForDB(value interface{}) (pqtype.NullRawMessage, error) {
	rawJSON, err := json.Marshal(value)
	if err != nil {
		return pqtype.NullRawMessage{}, err
	}
	return pqtype.NullRawMessage{
		RawMessage: rawJSON,
		Valid:      true,
	}, nil
}

func DecodeJson(raw json.RawMessage) (map[string]interface{}, error) {
	var value map[string]interface{}
	if err := json.Unmarshal(raw, &value); err != nil {
		return nil, err
	}
	return value, nil
}

func MapToStruct(dataMap map[string]interface{}) (interface{}, error) {
	// Create a new type dynamically
	var fields []reflect.StructField
	for key, value := range dataMap {
		fieldType := reflect.TypeOf(value)
		field := reflect.StructField{
			Name: strings.Title(key),
			Type: fieldType,
		}
		fields = append(fields, field)
	}

	dynamicType := reflect.StructOf(fields)
	dynamicValue := reflect.New(dynamicType).Elem()

	for key, value := range dataMap {
		field := dynamicValue.FieldByName(strings.Title(key))
		field.Set(reflect.ValueOf(value))
	}

	return dynamicValue.Interface(), nil
}

