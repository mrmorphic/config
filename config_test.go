package config

import (
	"testing"
)

func TestBasicProps(t *testing.T) {
	conf, e := ReadFromFile("./test.json")
	if e != nil {
		t.Errorf("Error reading config files: '%s'", e)
	}

	// read intprop
	i := conf.Get("test1.intprop")
	ii, ok := i.(float64)
	if !ok {
		t.Errorf("Expected test1.intprop to be a float")
	} else if ii != 123.0 {
		t.Errorf("Expected test1.intprop to have value 123, has %s", ii)
	}

	// read strprop
	s := conf.Get("test1.strprop")
	ss, ok := s.(string)
	if !ok {
		t.Errorf("Expected test1.strprop to be a string")
	} else if ss != "str" {
		t.Errorf("Expected test1.strprop to have value 'str', has '%s'", ii)
	}

	// read boolprop
	b := conf.Get("test1.boolprop")
	bb, ok := b.(bool)
	if !ok {
		t.Errorf("Expected test1.boolprop to be a bool")
	} else if !bb {
		t.Errorf("Expected test1.boolprop to be true, obviously wasn't")
	}

	// read non-existent property
	n := conf.Get("non-existent")
	if n != nil {
		t.Errorf("Got a value for a non-existent property: %s", n)
	}
}

func TestAsInt(t *testing.T) {
	conf, _ := ReadFromFile("./test.json")

	// test intprop using GetInt
	i, e := conf.AsInt("test1.intprop")
	if e != nil {
		t.Errorf("Error reading test1.intprop using AsInt")
	} else if i != 123 {
		t.Errorf("Expected test1.intprop to have value 123, has %s", i)
	}

	// test stringprop using GetInt, expecting an error
	_, e = conf.AsInt("test1.strprop")
	if e == nil {
		t.Errorf("Expected an error reading test1.strprop using AsInt, didn't get an error")
	}
}

func TestAsString(t *testing.T) {
	conf, _ := ReadFromFile("./test.json")

	if conf.AsString("test1.intprop") != "123" {
		t.Errorf("Expected test1.intprop to have value \"123\" via AsString")
	}

	if conf.AsString("test1.strprop") != "str" {
		t.Errorf("Expected test1.intprop to have value \"str\" via AsString")
	}

	if conf.AsString("non-existent") != "" {
		t.Errorf("Expected test1.intprop to have value \"\" via AsString")
	}
}

func TestNestedProps(t *testing.T) {
	conf, _ := ReadFromFile("./test.json")

	if conf.AsString("test1.child1.child1_2.x") != "c1_2" {
		t.Errorf("Expected test1.child1.child1_2.x to have value \"c1_2\" via AsString")
	}
}

func TestArrayProps(t *testing.T) {
	conf, _ := ReadFromFile("./test.json")

	// read arrayprop
	a := conf.Get("test1.arrayprop")
	aa, ok := a.([]interface{})
	if !ok {
		t.Errorf("Expected test1.arrayprop to be an array")
	} else if len(aa) != 3 {
		t.Errorf("Expected test1.arrayprop to have length 3, has '%d'", len(aa))
	}
}
