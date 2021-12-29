package settman

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// Allowed types for settings
var (
	Boolean = reflect.TypeOf((*bool)(nil)).Elem()

	Int  = reflect.TypeOf((*int)(nil)).Elem()
	Uint = reflect.TypeOf((*uint)(nil)).Elem()

	Int8  = reflect.TypeOf((*int8)(nil)).Elem()
	Int16 = reflect.TypeOf((*int16)(nil)).Elem()
	Int32 = reflect.TypeOf((*int32)(nil)).Elem()
	Int64 = reflect.TypeOf((*int64)(nil)).Elem()

	Uint8  = reflect.TypeOf((*uint8)(nil)).Elem()
	Uint16 = reflect.TypeOf((*uint16)(nil)).Elem()
	Uint32 = reflect.TypeOf((*uint32)(nil)).Elem()
	Uint64 = reflect.TypeOf((*uint64)(nil)).Elem()

	Float32 = reflect.TypeOf((*float32)(nil)).Elem()
	Float64 = reflect.TypeOf((*float64)(nil)).Elem()

	String = reflect.TypeOf((*string)(nil)).Elem()
)

type setting struct {
	name     string
	dataType reflect.Type
	value    interface{}
}

// If you do not specify a defaultValue, you make this a mandatory setting
func NewSetting(name string, dataType reflect.Type, defaultValue interface{}) (*setting, error) {
	if defaultValue != nil && reflect.TypeOf(defaultValue) != dataType {
		return nil, fmt.Errorf("Default value for setting can't have a different type than dataType")
	}

	return &setting{
		name:     name,
		dataType: dataType,
		value:    defaultValue,
	}, nil
}

// Get a value setting, or his default
func (s setting) Get() interface{} {
	return s.value
}

func (s *setting) Parse() error {
	//Get value from env
	val := os.Getenv(s.name)
	if len(val) == 0 {
		return nil
	}

	//No need to unmarshall for strings
	if s.dataType == String {
		s.value = val
		return nil
	}

	// Unmarshall into a new object of needed type
	zero := reflect.New(s.dataType)
	err := json.Unmarshal([]byte(val), zero.Interface())
	if err != nil {
		return err
	}

	//Convert from <interface *type> to <interface type>
	s.value = zero.Elem().Addr().Elem().Interface()
	return nil
}
