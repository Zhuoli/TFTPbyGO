package common

import (
	"testing"
	"reflect"
	"packets"
)

func TestParseRequestPacket(t *testing.T) {
	testCases := []struct {
		packet         []byte
		expectedPacket *RequestPacket
		shouldError    bool
	}{
		// Nil packet
		{
			packet:         nil,
			expectedPacket: nil,
			shouldError:    true,
		},
		// Empty packet
		{
			packet:         []byte{},
			expectedPacket: nil,
			shouldError:    true,
		},
		// RRQ
		{
			packet: []byte{0, 1, 'H', 'e', 'l', 'l', 'o', 0, 'M', 'o', 'd', 'e', 0},
			expectedPacket: &RequestPacket{
				OpCode:   packets.OpRRQ,
				Filename: "Hello",
				Mode:     "Mode",
			},
			shouldError: false,
		},
		// WRQ
		{
			packet: []byte{0, 2, 66, 0, 66, 0},
			expectedPacket: &RequestPacket{
				OpCode:   packets.OpWRQ,
				Filename: "B",
				Mode:     "B",
			},
			shouldError: false,
		},
		// Invalid name
		{
			packet:         []byte{0, 1, 'H', 'e', 'l', 'l', 'o'},
			expectedPacket: nil,
			shouldError:    true,
		},
		// Invalid mode
		{
			packet:         []byte{0, 1, 'H', 'e', 'l', 'l', 'o', 0, 'A'},
			expectedPacket: nil,
			shouldError:    true,
		},
		// Invalid opcode
		{
			packet:         []byte{1, 1, 'H', 'e', 'l', 'l', 'o'},
			expectedPacket: nil,
			shouldError:    true,
		},
	}

	for i, tc := range testCases {
		packet, err := ParseRequestPacket(tc.packet)
		if tc.shouldError && err == nil {
			t.Errorf("Expected error, didn't get one (%d)", i)
		}
		if !tc.shouldError && err != nil {
			t.Errorf("%v (%d)", err, i)
		}
		if !reflect.DeepEqual(tc.expectedPacket, packet) {
			t.Errorf("Test case %d failed", i)
			t.Errorf("Expected")
			t.Errorf("%v", tc.expectedPacket)
			t.Errorf("Got")
			t.Errorf("%v", packet)
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
		oc, err := packets.GetOpCode(tc.data)
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
