package human

import "testing"

func TestHumanMove(test *testing.T) {
	newHumanStruct := NewHuman("coma333ryu", 40, true)
	newHumanStruct.Move()
	newHumanStruct.Communicate()
}
