package server


import (
	"io/ioutil"
	"log"
	"testing"
	"common"
	"packets"
	"strings"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func acceptedMode(mode string) bool {
	switch strings.ToLower(mode) {
	case "netascii", "octet", "mail":
		return true
	}
	return false
}
func TestParseACKPacket(t *testing.T) {
	testCases := []struct {
		packet      []byte
		tid         uint16
		errExpected bool
	}{
		// Valid packet
		{
			packet:      []byte{0, 4, 0, 1},
			tid:         1,
			errExpected: false,
		},
		// Wrong opcode
		{
			packet:      []byte{0, 3, 0, 1},
			tid:         1,
			errExpected: true,
		},
	}

	for i, tc := range testCases {
		tid, err := packets.ParseAckPacket(tc.packet)
		if tc.errExpected && err == nil {
			t.Errorf("Expected an error, got nil (%d)", i)
			continue
		}
		if !tc.errExpected && err != nil {
			t.Errorf("Error: %v (%d)", err, i)
			continue
		}
		if tc.errExpected && err != nil {
			continue
		}
		if tid != tc.tid {
			t.Errorf("Expected tid: %d, got %d (%d)", tc.tid, tid, i)
		}
	}
}

func sampleRRQ() []byte {
	return []byte{0, 1, 'H', 'e', 'l', 'l', 'o', 'R', 0, 'n', 'e', 't', 'a', 's', 'c', 'i', 'i', 0}
}

func sampleWRQ() []byte {
	return []byte{0, 2, 'H', 'e', 'l', 'l', 'o', 'W', 0, 'n', 'e', 't', 'a', 's', 'c', 'i', 'i', 0}
}

func TestAcceptedMode(t *testing.T) {
	testCases := []struct {
		mode     string
		accepted bool
	}{
		// Three accepted modes
		{mode: "netascii", accepted: true},
		{mode: "octet", accepted: true},
		{mode: "mail", accepted: true},

		// Mixed case should be allowed
		{mode: "netAscii", accepted: true},
		{mode: "OcteT", accepted: true},
		{mode: "Mail", accepted: true},

		// Anything else should be rejected
		{mode: "", accepted: false},
		{mode: "mode", accepted: false},
		{mode: "blah", accepted: false},
	}

	for _, tc := range testCases {
		outcome := acceptedMode(tc.mode)
		if outcome != tc.accepted {
			t.Errorf("Expected mode, '%s' accepted = %v", tc.mode, tc.accepted)
		}
	}
}


func TestGetOpcode(t *testing.T) {
	testCases := []struct {
		data           []byte
		expectedOpcode packets.OpCode
		shouldError    bool
	}{
		// Standard RRQ
		{
			data:           []byte{0, 1},
			expectedOpcode: packets.OpRRQ,
			shouldError:    false,
		},
		// Empty data
		{
			data:           []byte{},
			expectedOpcode: packets.OpERROR,
			shouldError:    true,
		},
		// Unknown opcode
		{
			data:           []byte{0, 99},
			expectedOpcode: packets.OpERROR,
			shouldError:    true,
		},
		// Only 1 byte
		{
			data:           []byte{1},
			expectedOpcode: packets.OpERROR,
			shouldError:    true,
		},
		// More than 2 bytes
		{
			data:           []byte{0, 1, 2},
			expectedOpcode: packets.OpRRQ,
			shouldError:    false,
		},
	}

	for i, tc := range testCases {
		oc, err := common.GetOpCode(tc.data)
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
