package messages

import (
	"fmt"
	"testing"
)

func TestCreateGetParameterValues(t *testing.T) {
	resp := NewGetParameterValues()
	var names []string
	names = append(names, "InternetGatewayDevice.DeviceInfo.Manufacturer", "InternetGatewayDevice.DeviceInfo.ProvisioningCode")
	resp.ParameterNames = names
	out, err := resp.CreateXML()
	if err != nil {
		t.FailNow()
	} else {
		fmt.Println(string(out))
	}
}

func TestParseGetParameterValues(t *testing.T) {

}
