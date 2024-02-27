// Package jsonobject implements utilities for JSON Parsing and Usage
package jsonobject

import (
	"encoding/json"
)

// JsonObject Type Contains the Parsed JSON Value
type JsonObject struct {
	// JSON value
	// object, array, string, number, true/false, null
	value interface{}
}

// Parses the JSON data and Returns JsonObject.
// If the data contains Invalid JSON, Then Returns Error
func Parse(data []byte) (JsonObject, error) {

	// Store the Parsed Data into value
	var value any
	e := json.Unmarshal(data, &value)

	if e != nil {
		return JsonObject{}, e
	}

	return JsonObject{value}, nil
}

// Get the Corresponding Value
//
//	object     - map[string]interface{}
//	array      - []interface{}
//	string     - string
//	nummer     - float64
//	true/false - bool
//	null       - nil
func (jsonObj *JsonObject) Value() any {
	return jsonObj.value
}

// If The JSON is either object or array then Return
// the element corresponding to the key or index. Else Return The Element.
//
// If They Key for Lookup is Invalid (Not a String for object, Invalid Index),
// Then nil Value JsonObject is Returned.
func (jsonObj *JsonObject) Get(k any) *JsonObject {

	switch v := (jsonObj.value).(type) {
	case map[string]interface{}:
		if kv, kok := (k).(string); kok {
			return &JsonObject{v[kv]}
		}
	case []interface{}:
		if kv, kok := (k).(int); kok && 0 <= kv && kv < len(v) {
			return &JsonObject{v[kv]}
		}
	default:
		return jsonObj
	}
	return &JsonObject{}
}
