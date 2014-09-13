// Config package provides an implementation of ConfigProvider. Currently it only reads from a JSON file, but
// this may be extended to support other forms.
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// Concrete storage for the configuration values read. The map keys are dot-delimited.
type Config map[string]interface{}

// ReadFromFile reads the configuration from JSON in the provided path. The config file
// will generally contain a single object. The properties of that object form the
// top-level name space for Get().
func ReadFromFile(path string) (Config, error) {
	// read file
	data, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	// decode json
	var v interface{}
	e = json.Unmarshal(data, &v)
	if e != nil {
		return nil, e
	}

	// Get this as a map
	nested := v.(map[string]interface{})

	// make map
	var result Config
	result = make(map[string]interface{})

	result.nestedMerge(nested, "")

	return result, nil
}

func (c Config) nestedMerge(object map[string]interface{}, prefix string) {
	p := prefix + "."
	if p == "." {
		p = ""
	}

	for k, v := range object {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			// if 'v' is a map of interface{}, recursively add.
			c.nestedMerge(v.(map[string]interface{}), p+k)
		} else {
			// otherwise just add the property, using the prefix
			c[p+k] = v
		}
	}
}

// Get looks up an object in the map via a key. The key can have "." separators for names;
// this will go into the structure as appropriate. It will return nil if a key maps to an undefined
// property, or where a partial key is not an object.
func (c Config) Get(key string) interface{} {
	return c[key]
}

// AsString returns a key from the configuration (using Get), but returning it as a string.
// If the key is not defined, it returns "".
func (c Config) AsString(key string) string {
	v := c.Get(key)
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

// AsInt returns a key from the configuration as an integer. Integer values in the
// json file are retrieved as float64. This will return an error if it's not a float64.
// Otherwise will convert it to an int (with truncation).
func (c Config) AsInt(key string) (int, error) {
	v := c.Get(key)
	vv, ok := v.(float64)
	if !ok {
		return 0, fmt.Errorf("Expected config property to be a numeric value, but wasn't: '%s'", v)
	}
	return int(vv), nil
}
