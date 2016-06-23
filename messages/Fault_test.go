package messages

import (
	"fmt"
	"testing"
)

func TestCreateFault(t *testing.T) {
	fault := NewFault()
	fault.ID = "ID:intrnl.unset.id.SetParameterValues1458897112421.1852997967"
	fault.FaultCode = "Client"
	fault.FaultString = "Client fault"
	fault.CwmpFaultCode = "9003"
	fault.CwmpFaultString = "Invalid arguments"
	fault.SetParameterValuesFault = SetParameterValuesFaultStruct{
		ParameterName: "InternetGatewayDevice.config.webauthglobal.successUrl",
		FaultCode:     "9008",
		FaultString:   "Attempt to set a non-writable parameter",
		ParameterKey:  "006aecb92b10459b91321e32c9b4502f",
	}
	out, err := fault.CreateXML()
	if err != nil {
		t.FailNow()
	} else {
		fmt.Println(string(out))
	}
}

func TestParseFault(t *testing.T) {
	data := `<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
	       xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/"
	       xmlns:xsi="http://www.w3.org/1999/XMLSchema-instance"
	       xmlns:xsd="http://www.w3.org/1999/XMLSchema"
	       xmlns:cwmp="urn:dslforum-org:cwmp-1-0">
	       <SOAP-ENV:Header>
	       <cwmp:ID SOAP-ENV:mustUnderstand="1">ID:intrnl.unset.id.SetParameterValues1458897112421.1852997967</cwmp:ID>
	       </SOAP-ENV:Header>
	       <SOAP-ENV:Body>
	       <SOAP-ENV:Fault>
	       <faultcode>Client</faultcode>
	       <faultstring>Client fault</faultstring>
	       <detail>
	       <cwmp:Fault>
	       <FaultCode>9003</FaultCode>
	       <FaultString>Invalid arguments</FaultString>
	       <SetParameterValuesFault>
	       <ParameterName>InternetGatewayDevice.config.webauthglobal.successUrl</ParameterName>
	       <FaultCode>9008</FaultCode>
	       <FaultString>Attempt to set a non-writable parameter</FaultString>
	       <ParameterKey>006aecb92b10459b91321e32c9b4502f</ParameterKey>
	       </SetParameterValuesFault>
	       </cwmp:Fault>
	       </detail>
	       </SOAP-ENV:Fault>
	       </SOAP-ENV:Body>
	       </SOAP-ENV:Envelope>`
	msg, _ := ParseXML([]byte(data))
	fault := msg.(*Fault)
	fmt.Println(fault.FaultCode)
	fmt.Println(fault.FaultString)
	fmt.Println(fault.CwmpFaultCode)
	fmt.Println(fault.CwmpFaultString)
	fmt.Println(fault.SetParameterValuesFault.FaultCode)
	fmt.Println(fault.SetParameterValuesFault.FaultString)
	fmt.Println(fault.SetParameterValuesFault.ParameterName)
	fmt.Println(fault.SetParameterValuesFault.ParameterKey)
}
