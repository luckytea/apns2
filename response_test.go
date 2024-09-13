package apns2_test

import (
	"encoding/json"
	"net/http"
	"testing"

	apns "github.com/luckytea/apns2"
	"github.com/stretchr/testify/assert"
)

// Unit Tests

func TestResponseSent(t *testing.T) {
	assert.Equal(t, http.StatusOK, apns.StatusSent)
	assert.True(t, (&apns.Response{StatusCode: 200}).Sent())
	assert.False(t, (&apns.Response{StatusCode: 400}).Sent())
}

func TestIntTimestampParse(t *testing.T) {
	response := &apns.Response{}
	payload := "{\"reason\":\"Unregistered\", \"timestamp\":1458114061260}"
	json.Unmarshal([]byte(payload), &response)
	assert.Equal(t, int64(1458114061260)/1000, response.Timestamp.Unix())
}

func TestInvalidTimestampParse(t *testing.T) {
	response := &apns.Response{}
	payload := "{\"reason\":\"Unregistered\", \"timestamp\": \"2016-01-16 17:44:04 +1300\"}"
	err := json.Unmarshal([]byte(payload), &response)
	assert.Error(t, err)
}
