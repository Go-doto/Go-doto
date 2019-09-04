package dota_api

import "testing"

var postitions = map[PlayerSlot]uint{
	0:   0,
	1:   1,
	2:   2,
	3:   3,
	4:   4,
	128: 0,
	129: 1,
	130: 2,
	131: 3,
	132: 4,
}

func TestPlayerSlot_IsDire(t *testing.T) {
	for slot, _ := range postitions {
		if slot > 128 && !slot.IsDire() {
			t.Error("Invalid position")
		}
	}
}

func TestPlayerSlot_IsRadiant(t *testing.T) {
	for slot, _ := range postitions {
		if slot < 128 && !slot.IsRadiant() {
			t.Error("Invalid position")
		}
	}
}

func TestPlayerSlot_GetPosition(t *testing.T) {
	for slot, pos := range postitions {
		if slot.GetPosition() != pos {
			t.Error("Invalid position")
		}
	}
}
