package messages

import (
	"fmt"
	"testing"
)

func TestCreateDownloadResponse(t *testing.T) {
	resp := NewDownloadResponse()
	resp.Status = 1
	resp.StartTime = "2015-02-12T13:40:07"
	resp.CompleteTime = "2015-02-12T13:40:07"
	out, err := resp.CreateXML()
	if err != nil {
		t.FailNow()
	} else {
		fmt.Println(string(out))
	}
}

func TestParseDownloadResponse(t *testing.T) {
	data := `<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:cwmp="urn:dslforum-org:cwmp-1-0">
	      <SOAP-ENV:Header>
	          <cwmp:ID SOAP-ENV:mustUnderstand="1">ID:intrnl.unset.id.DownloadResponse1439551393.1439551393594569717</cwmp:ID>
	      </SOAP-ENV:Header>
	      <SOAP-ENV:Body>
	          <cwmp:DownloadResponse>
	              <StartTime>2015-02-12T13:40:07</StartTime>
	              <CompleteTime>2015-02-12T13:40:07</CompleteTime>
	              <Status>1</Status>
	          </cwmp:DownloadResponse>
	      </SOAP-ENV:Body>
	  </SOAP-ENV:Envelope>`
	msg, _ := ParseXML([]byte(data))
	resp := msg.(*DownloadResponse)
	fmt.Println(resp.Status)
	fmt.Println(resp.StartTime)
	fmt.Println(resp.CompleteTime)
}
