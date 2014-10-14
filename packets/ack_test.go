package packets

import (
	"testing"
	"reflect"
)

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

	//range on arrays and slices provides both the index and value for
	// each entry.
	for i, tc := range testCases {
		tid, err := ParseAckPacket(tc.packet)
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

func TestCreateAckPacket(t *testing.T) {
	testCases := []struct {
		tid      uint16
		expected []byte
	}{
		{
			tid:      1,
			expected: []byte{0, 4, 0, 1},
		},
		{
			tid:      14,
			expected: []byte{0, 4, 0, 14},
		},
	}

	for i, tc := range testCases {
		packet := CreateAckPacket(tc.tid)
		if !reflect.DeepEqual(packet, tc.expected) {
			t.Errorf("Expected and actual packet not equal (%d)", i)
			t.Error(packet)
		}
	}
}