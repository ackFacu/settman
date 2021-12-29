package settman

import (
	"os"
	"testing"
)

var optionalUint, mandatoryUint *setting
var optionalString, mandatoryString *setting
var optionalBoolean, mandatoryBoolean *setting

func TestMain(m *testing.M) {
	optionalUint, _ = NewSetting(
		"optionalUint",
		Uint8,
		uint8(1),
	)

	optionalString, _ = NewSetting(
		"optionalString",
		String,
		"HolaTest",
	)

	optionalBoolean, _ = NewSetting(
		"optionalBoolean",
		Boolean,
		true,
	)

	mandatoryUint, _ = NewSetting(
		"mandatoryUint",
		Uint8,
		nil,
	)

	mandatoryString, _ = NewSetting(
		"mandatoryString",
		String,
		nil,
	)

	mandatoryBoolean, _ = NewSetting(
		"mandatoryBoolean",
		Boolean,
		nil,
	)

	os.Exit(m.Run())
}

// Get values before parsing
func TestGetWithoutEnv(t *testing.T) {
	opUint := optionalUint
	opString := optionalString
	opBoolean := optionalBoolean
	maUint := mandatoryUint
	maString := mandatoryString
	maBoolean := mandatoryBoolean

	opUint.Parse()
	opString.Parse()
	opBoolean.Parse()
	maUint.Parse()
	maString.Parse()
	maBoolean.Parse()

	if opUint.Get().(uint8) != uint8(1) {
		t.Error("Fail to get default value")
		t.Fail()
	}

	if opString.Get().(string) != "HolaTest" {
		t.Error("Fail to get default value")
		t.Fail()
	}

	if opBoolean.Get().(bool) != true {
		t.Error("Fail to get default value")
		t.Fail()
	}

	if maUint.Get() != nil {
		t.Error("Fail to create new mandatory setting")
		t.Fail()
	}

	if maString.Get() != nil {
		t.Error("Fail to create new mandatory setting")
		t.Fail()
	}

	if maBoolean.Get() != nil {
		t.Error("Fail to create new mandatory setting")
		t.Fail()
	}
}

// Get values after parsing
func TestGetWithEnv(t *testing.T) {
	opUint := optionalUint
	opString := optionalString
	opBoolean := optionalBoolean
	maUint := mandatoryUint
	maString := mandatoryString
	maBoolean := mandatoryBoolean

	os.Setenv("optionalUint", "100")
	os.Setenv("optionalString", "Nuevo string value")
	os.Setenv("optionalBoolean", "false")
	os.Setenv("mandatoryUint", "100")
	os.Setenv("mandatoryString", "Nuevo string value")
	os.Setenv("mandatoryBoolean", "false")

	opUint.Parse()
	opString.Parse()
	opBoolean.Parse()
	maUint.Parse()
	maString.Parse()
	maBoolean.Parse()

	if opUint.Get().(uint8) != uint8(100) {
		t.Error("Fail to get value from .env for optional setting uint")
		t.Fail()
	}

	if opString.Get().(string) != "Nuevo string value" {
		t.Error("Fail to get value from .env for optional setting string")
		t.Fail()
	}

	if opBoolean.Get().(bool) != false {
		t.Error("Fail to get value from .env for optional setting boolean")
		t.Fail()
	}

	if maUint.Get().(uint8) != uint8(100) {
		t.Error("Fail to get value from .env for mandatory setting uint")
		t.Fail()
	}

	if maString.Get().(string) != "Nuevo string value" {
		t.Error("Fail to get value from .env for mandatory setting string")
		t.Fail()
	}

	if maBoolean.Get().(bool) != false {
		t.Error("Fail to get value from .env for mandatory setting boolean")
		t.Fail()
	}
}
