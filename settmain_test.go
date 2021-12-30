package settman

import (
	"os"
	"testing"
)

var optionalUint, mandatoryUint *setting
var optionalString, mandatoryString *setting
var optionalBoolean, mandatoryBoolean *setting

func TestMain(m *testing.M) {
	optionalUint = NewSetting(
		"optionalUint",
		Uint8,
		uint8(1),
	)

	optionalString = NewSetting(
		"optionalString",
		String,
		"HolaTest",
	)

	optionalBoolean = NewSetting(
		"optionalBoolean",
		Boolean,
		true,
	)

	mandatoryUint = NewSetting(
		"mandatoryUint",
		Uint8,
		nil,
	)

	mandatoryString = NewSetting(
		"mandatoryString",
		String,
		nil,
	)

	mandatoryBoolean = NewSetting(
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

	//This are going to use default value
	if nil != opUint.Parse() {
		t.Error("Parse on opUint expected to succeed but failed")
		t.Fail()
	}

	if nil != opString.Parse() {
		t.Error("Parse on opString expected to succeed but failed")
		t.Fail()
	}

	if nil != opBoolean.Parse() {
		t.Error("Parse on opBoolean expected to succeed but failed")
		t.Fail()
	}

	if nil == maUint.Parse() {
		t.Error("Parse on maUint expected to fail but succeeded")
		t.Fail()
	}

	if nil == maString.Parse() {
		t.Error("Parse on maString expected to fail but succeeded")
		t.Fail()
	}

	if nil == maBoolean.Parse() {
		t.Error("Parse on maBoolean expected to fail but succeeded")
		t.Fail()
	}

	//Test parse with invalid types
	_ = os.Setenv("mandatoryUint", "100.0") //Expect uint but got float
	_ = os.Setenv("mandatoryBoolean", "1")  //Expect boolean but got a number

	if nil == maUint.Parse() {
		t.Error("Parse on maUint expected to fail but succeeded")
		t.Fail()
	}

	if nil == maBoolean.Parse() {
		t.Error("Parse on maBoolean expected to fail but succeeded")
		t.Fail()
	}

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
