package jsonobject

import "testing"

func TestJsonObjectParse(t *testing.T) {
	var data []byte
	var e error
	var jsonObj JsonObject

	// Parse Success
	data = []byte("\"Hello\"")
	_, e = Parse(data)
	if e != nil {
		t.Error("Error While Parsing: ", e)
	}

	// Parse Fail - Missing " Should Return Error
	data = []byte("\"Hello")
	jsonObj, e = Parse(data)
	if jsonObj.value != nil || e == nil {
		t.Error("No Parse Error. Should have been a Parse Error", jsonObj, e)
	}
}

func TestJsonObjectParseDataTypeObject(t *testing.T) {
	data := []byte(`{
		"Key": "Value"
	}`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	if obj, ok := jsonObj.Value().(map[string]interface{}); ok {
		for k, v := range obj {
			if k != "Key" || v != "Value" {
				t.Error("Key or Value Not Found")
			}
		}
	} else {
		t.Errorf("Object Type Not Matching")
	}
}

func TestJsonObjectParseDataTypeArray(t *testing.T) {
	data := []byte(`["a", "b", "c"]`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	expectedValues := map[int]interface{}{
		0: "a",
		1: "b",
		2: "c",
	}

	if obj, ok := jsonObj.Value().([]interface{}); ok {
		for k, v := range obj {
			if expectedValues[k] != v {
				t.Error("Value Not Matching ", k, v)
			}
		}
	} else {
		t.Errorf("Object Type Not Matching")
	}

}

func TestJsonObjectParseDataTypeString(t *testing.T) {
	data := []byte(`"a"`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	if obj, ok := jsonObj.Value().(string); ok {
		if obj != "a" {
			t.Error("Value Not Matching")
		}
	} else {
		t.Error("Object Type Not Matching")
	}
}

func TestJsonObjectParseDataTypeNumber(t *testing.T) {
	data := []byte(`1`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	if obj, ok := jsonObj.Value().(float64); ok {
		if obj != 1 {
			t.Error("Value Not Matching")
		}
	} else {
		t.Error("Object Type Not Matching")
	}
}

func TestJsonObjectParseDataTypeBool(t *testing.T) {
	data := []byte(`true`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	if obj, ok := jsonObj.Value().(bool); ok {
		if obj != true {
			t.Error("Value Not Matching")
		}
	} else {
		t.Error("Object Type Not Matching")
	}
}

func TestJsonObjectParseDataTypeNull(t *testing.T) {
	data := []byte(`null`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	if jsonObj.Value() != nil {
		t.Error("Value Not Matching")
	}
}

func TestJsonObjectGet(t *testing.T) {
	data := []byte(`{
		"string_array": ["asdf", "ghjk", "zxcv"],
		"nested_obj": {"k": 1}
	}`)
	jsonObj, e := Parse(data)
	if e != nil {
		t.Error("Parse Error: ", e)
	}

	// Get Value from Object Using Key
	if _, ok := jsonObj.Get("string_array").Value().([]interface{}); !ok {
		t.Error("Invalid Value for Key ", "string_array")
	}

	// Get Element from Array Using Index
	if arrayElement, ok := jsonObj.Get("string_array").Get(0).Value().(string); !ok || arrayElement != "asdf" {
		t.Error("Invalid Value for Index", "string_array[0]", arrayElement)
	}

	// Get Nested Element from Object Using Key
	if nestedValue, ok := jsonObj.Get("nested_obj").Get("k").Value().(float64); !ok || nestedValue != 1 {
		t.Error("Invalid Value for Index", "nested_obj[k]", nestedValue)
	}

	// Get Invalid Key from Object
	if jsonObj.Get("invalid_key").Value() != nil {
		t.Error("Found Not Null")
	}

	// Get Invalid Index from Array
	if jsonObj.Get("string_array").Get(-1).Value() != nil {
		t.Error("Found Not Null")
	}
}
