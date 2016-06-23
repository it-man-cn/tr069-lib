package messages

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateSetParameterValues(t *testing.T) {
	resp := NewSetParameterValues()
	params := make(map[string]ValueStruct)
	param := ValueStruct{XsdString, "abc"}
	params["InternetGatewayDevice.DeviceInfo.Manufacturer"] = param
	resp.Params = params
	out, err := resp.CreateXML()
	if err != nil {
		t.FailNow()
	} else {
		fmt.Println(string(out))
	}
	jsonstr, _ := json.Marshal(&resp)
	fmt.Println(string(jsonstr))

}

func TestParseSetParameterValues(t *testing.T) {

}
