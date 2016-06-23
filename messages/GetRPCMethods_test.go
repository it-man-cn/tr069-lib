package messages

import (
	"fmt"
	"testing"
)

func TestCreateGetRPCMethods(t *testing.T) {
	resp := NewGetRPCMethods()
	out, err := resp.CreateXML()
	if err != nil {
		t.FailNow()
	} else {
		fmt.Println(string(out))
	}
}
