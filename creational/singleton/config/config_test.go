package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConfig(t *testing.T) {
	conf := GetInstance()
	//If nil, error
	if conf == nil {
		t.Error("Expected pointer to singleton")
	}
	conf.Set("TestKey", "TestValue")
	conf2 := GetInstance() //should point to same
	if conf2.Get("TestKey") != "TestValue" {
		t.Errorf("Expected replica conf2 to have config with key 'TestKey'")
	}

	conf2.Set("TestKey2", "TestValue2")
	assert.Equal(t, "TestValue", conf2.Get("TestKey"))
	assert.Equal(t, "TestValue2", conf2.Get("TestKey2"))
}
