package settman

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sync"
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
	name         string
	dataType     reflect.Type
	defaultValue interface{}
	value        interface{}
	m            sync.RWMutex
}

// If you do not specify a defaultValue, you make this a mandatory setting
func NewSetting(name string, dataType reflect.Type, defaultValue interface{}) *setting {
	if defaultValue != nil && reflect.TypeOf(defaultValue) != dataType {
		panic(fmt.Errorf("default value for setting %s has a diferent type than the one configured", name))
		return nil
	}

	return &setting{
		name:         name,
		dataType:     dataType,
		defaultValue: defaultValue,
		value:        defaultValue,
	}
}

// Get a value setting, or his default
// It is safe to guess that this is always going to return a valid value, safe to convert to expected type, if the setting has been parsed
func (s *setting) Get() interface{} {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.value
}

// Set value for a setting. Internal use only
func (s *setting) set(v interface{}) {
	s.m.Lock()
	defer s.m.Unlock()

	s.value = v
}

func (s *setting) Parse() {
	defer s.checkConsistency()

	//Get value from env
	val := os.Getenv(s.name)

	//No env value set, use default
	if len(val) == 0 {
		s.set(s.defaultValue)
		return
	}

	//No need to unmarshall for string type
	if s.dataType == String {
		s.set(val)
		return
	}

	// Unmarshall into a new object of needed type
	zero := reflect.New(s.dataType)
	if nil == json.Unmarshal([]byte(val), zero.Interface()) {
		//Convert from <interface *type> to <interface type>
		s.set(zero.Elem().Addr().Elem().Interface())
	} else {
		// invalid type provided on environment variable, unable to unmarshall. Use default
		s.set(s.defaultValue)
	}

	return
}

// Check for valid value set on setting - This is only going to panic if mandatory setting is missing or invalid
func (s *setting) checkConsistency() {
	if s.Get() == nil || reflect.TypeOf(s.Get()) != s.dataType {
		panic(fmt.Errorf("mandatory setting is missing or invalid type: %s", s.name))
	}
}
