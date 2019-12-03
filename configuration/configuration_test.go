package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestLoadConfig(test *testing.T) {
	LoadConfig(map[string]interface{}{})
}

func TestGetFlag(test *testing.T) {
	key := "bob"
	flag := "--" + key
	expected := "mary"
	actual := getFlag(map[string]interface{}{flag: expected}, key)
	assert.Equal(test, expected, actual, "Same?")
}
