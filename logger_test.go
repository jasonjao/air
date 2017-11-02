package air

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	assert.NotNil(t, loggerSingleton)
	assert.Nil(t, loggerSingleton.template)
	assert.NotNil(t, loggerSingleton.mutex)

	buf := &bytes.Buffer{}
	LoggerOutput = buf

	func() {
		loggerSingleton.log("INFO", "foo", "bar")
	}()
	assert.Zero(t, buf.Len())

	LoggerEnabled = true

	loggerSingleton.log("INFO", "foo", "bar")
	assert.Zero(t, buf.Len())

	func() {
		loggerSingleton.log("INFO", "foo", "bar")
	}()

	m := map[string]interface{}{}
	assert.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, "foobar", m["message"])

	LoggerEnabled = false
	LoggerOutput = os.Stdout
	loggerSingleton.template = nil
}