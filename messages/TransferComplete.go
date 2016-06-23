package messages

import (
	"encoding/xml"
	"fmt"
	"github.com/jteeuwen/go-pkg-xmlx"
	"strconv"
	"time"
)

//TransferComplete download complete
type TransferComplete struct {
	ID           string
	Name         string
	CommandKey   string
	StartTime    string
	CompleteTime string
	FaultCode    int
	FaultString  string
}

type transferCompleteBodyStruct struct {
	Body transferCompleteStruct `xml:"cwmp:TransferComplete"`
}

type transferCompleteStruct struct {
	CommandKey   string
	StartTime    string
	CompleteTime string
	Fault        interface{} `xml:"FaultStruct,ommitempty"`
}

//NewTransferComplete create TransferComplete object
func NewTransferComplete() (m *TransferComplete) {
	m = &TransferComplete{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetID get msg id
func (msg *TransferComplete) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg type
func (msg *TransferComplete) GetName() string {
	return "TransferComplete"
}

//CreateXML encode into mxl
func (msg *TransferComplete) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id}
	var transf transferCompleteStruct
	if len(msg.FaultString) > 0 {
		fault := FaultStruct{FaultCode: msg.FaultCode, FaultString: msg.FaultString}
		transf = transferCompleteStruct{
			CommandKey:   msg.CommandKey,
			StartTime:    msg.StartTime,
			CompleteTime: msg.CompleteTime,
			Fault:        fault,
		}
	} else {
		transf = transferCompleteStruct{
			CommandKey:   msg.CommandKey,
			StartTime:    msg.StartTime,
			CompleteTime: msg.CompleteTime,
		}
	}

	env.Body = transferCompleteBodyStruct{transf}
	//output, err := xml.Marshal(env)
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *TransferComplete) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	fmt.Println(doc.String())
	msg.CommandKey = doc.SelectNode("*", "CommandKey").GetValue()
	msg.CompleteTime = doc.SelectNode("*", "CompleteTime").GetValue()
	msg.StartTime = doc.SelectNode("*", "StartTime").GetValue()
	faultStringNode := doc.SelectNode("*", "FaultString")

	if faultStringNode != nil {
		msg.FaultString = doc.SelectNode("*", "FaultString").GetValue()
	}

	faultCodeNode := doc.SelectNode("*", "FaultCode")
	if faultCodeNode != nil {
		faultCode, err := strconv.Atoi(faultCodeNode.GetValue())
		if err != nil {
			return err
		}
		msg.FaultCode = faultCode
	}
	return nil
}
