package encode_decode

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestMessage(t *testing.T) {
	//key: field << 3 | wire

	//value: 二进制 -> 7位一组 -> 逆序 -> 每组最高补1（第1组例外）
	varintsA := VarintsA{
		A: 723,
	}
	byteArr, _ := proto.Marshal(&varintsA)
	result1 := fmt.Sprintf("%X", byteArr)
	if result1 != "08D305" {
		t.Errorf("varints expected: 08D305, real: %s", result1)
	}

	// N>0 -> 2*N; N<0 -> 2*(-N)-1; varints
	zigZagA := ZigZagA{
		B: -362,
	}
	byteArr, _ = proto.Marshal(&zigZagA)
	result2 := fmt.Sprintf("%X", byteArr)
	if result2 != "08D305" {
		t.Errorf("zigzag expected: 08D305, real: %s", result2)
	}

	var varintsB VarintsA
	_ = proto.Unmarshal(byteArr, &varintsB)
	result3 := varintsB.A
	if result3 != 723 {
		t.Errorf("varint expected: 723, real: %d", result3)
	}

	// size + string
	stringA := StringA{
		A: "helloworld",
	}
	byteArr, _ = proto.Marshal(&stringA)
	result4 := fmt.Sprintf("%X", byteArr)
	if result4 != "3A0A68656C6C6F776F726C64" {
		t.Errorf("string expected: 3A0A68656C6C6F776F726C64, real: %s", result4)
	}

	// size + values(ordered)
	repeatedA := ReapeatedA{
		A: []int32{1,2,3,4,5},
	}
	byteArr, _ = proto.Marshal(&repeatedA)
	result5 := fmt.Sprintf("%X", byteArr)
	if result5  != "1A050102030405" {
		t.Errorf("repeated expected: 1A050102030405, real: %s", result5)
	}

	multiFieldA := MultiFieldA{
		A: 723,
		B: "helloworld",
	}
	byteArr, _ = proto.Marshal(&multiFieldA)
	result6 := fmt.Sprintf("%X", byteArr)
	if result6  != "08D3053A0A68656C6C6F776F726C64" {
		t.Errorf("multiFieldA expected: 08D3053A0A68656C6C6F776F726C64, real: %s", result6)
	}

	embeddedA := EmbeddedA{
		A: &StringA{
			A: "helloworld",
		},
	}
	byteArr, _ = proto.Marshal(&embeddedA)
	result7 := fmt.Sprintf("%X", byteArr)
	if result7 != "3A0C3A0A68656C6C6F776F726C64" {
		t.Errorf("embeddedA expected: 3A0C3A0A68656C6C6F776F726C64, real: %s", result7)
	}
}