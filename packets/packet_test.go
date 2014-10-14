package packets

import (
	"testing"
)

func TestGetOpcode(t *testing.T) {
	testCases := []struct {
		data           []byte
		expectedOpcode OpCode
		shouldError    bool
	}{
		// Standard RRQ
		{
			data:           []byte{0, 1},
			expectedOpcode: OpRRQ,
			shouldError:    false,
		},
		// Empty data
		{
			data:           []byte{},
			expectedOpcode: OpERROR,
			shouldError:    true,
		},
		// Unknown opcode
		{
			data:           []byte{0, 67},
			expectedOpcode: OpERROR,
			shouldError:    true,
		},
		// Only 1 byte
		{
			data:           []byte{1},
			expectedOpcode: OpERROR,
			shouldError:    true,
		},
		// More than 2 bytes
		{
			data:           []byte{0, 1, 2},
			expectedOpcode: OpRRQ,
			shouldError:    false,
		},
	}

	for i, tc := range testCases {
		oc, err := GetOpCode(tc.data)
		if tc.shouldError && err == nil {
			t.Errorf("Expected error, didn't get one (%d)", i)
			continue
		}
		if !tc.shouldError && err != nil {
			t.Errorf("%v (%d)", err, i)
			continue
		}
		if oc != tc.expectedOpcode {
			t.Errorf("Expected: %v, got %v (%d)", tc.expectedOpcode, oc, i)
		}
	}
}