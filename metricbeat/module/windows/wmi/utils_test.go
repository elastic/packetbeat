package wmi

import (
	"fmt"
	"testing"
	"time"

	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/stretchr/testify/assert"
)

type MockWmiSession struct {
}

const MockTimeout = time.Second * 5

// Mock Implementation of QueryInstances function
// This simulate a long-running query
func (c *MockWmiSession) QueryInstances(queryExpression string) ([]*wmi.WmiInstance, error) {
	time.Sleep(MockTimeout)
	return []*wmi.WmiInstance{}, nil
}

func TestExecuteGuardedQueryInstances(t *testing.T) {
	mockSession := new(MockWmiSession)
	query := "SELECT * FROM Win32_OpeartingSystem"
	timeout := 200 * time.Millisecond

	startTime := time.Now()
	expectedError := fmt.Errorf("the execution of the query'%s' exceeded the threshold of %s", query, timeout)
	_, err := ExecuteGuardedQueryInstances(mockSession, query, timeout)
	// Make sure the return time is less than the MockTimeout
	assert.Less(t, time.Since(startTime), MockTimeout, "The return time should be less than the sleep time")
	// Make sure the error returned is the expected one
	assert.Equal(t, err, expectedError, "Expected the returned error to match the expected error")
}

func Test_ConversionFunctions(t *testing.T) {
	tests := []struct {
		name        string
		conversion  WmiStringConversionFunction
		input       string
		expected    interface{}
		expectErr   bool
		description string
	}{
		// Test cases for ConvertUint64
		{
			name:        "ConvertUint64 - valid input",
			conversion:  ConvertUint64,
			input:       "12345",
			expected:    uint64(12345),
			expectErr:   false,
			description: "Should convert string to uint64",
		},
		{
			name:        "ConvertUint64 - invalid input",
			conversion:  ConvertUint64,
			input:       "notANumber",
			expected:    nil,
			expectErr:   true,
			description: "Should return error for invalid uint64 string",
		},

		// Test cases for ConvertSint64
		{
			name:        "ConvertSint64 - valid input",
			conversion:  ConvertSint64,
			input:       "-12345",
			expected:    int64(-12345),
			expectErr:   false,
			description: "Should convert string to sint64",
		},
		{
			name:        "ConvertSint64 - invalid input",
			conversion:  ConvertSint64,
			input:       "notANumber",
			expected:    nil,
			expectErr:   true,
			description: "Should return error for invalid sint64 string",
		},

		// Test cases for ConvertDatetime
		{
			name:        "ConvertDatetime - valid input",
			conversion:  ConvertDatetime,
			input:       "20231224093045.123456-000",
			expected:    mustParseTime("20060102150405.999999-0700", "20231224093045.123456-0000"),
			expectErr:   false,
			description: "Should convert string to time.Time",
		},
		{
			name:        "ConvertDatetime - invalid input",
			conversion:  ConvertDatetime,
			input:       "invalidDatetime",
			expected:    nil,
			expectErr:   true,
			description: "Should return error for invalid datetime string",
		},
		// Test cases for ConvertString
		{
			name:        "ConvertString - valid input",
			conversion:  ConvertString,
			input:       "test string",
			expected:    "test string",
			expectErr:   false,
			description: "Should return the same string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.conversion(tt.input)

			if tt.expectErr {
				assert.Error(t, err, tt.description)
			} else {
				assert.NoError(t, err, tt.description)
				assert.Equal(t, tt.expected, result, tt.description)
			}
		})
	}
}

// Helper function to parse time
func mustParseTime(layout, value string) time.Time {
	parsed, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return parsed
}
