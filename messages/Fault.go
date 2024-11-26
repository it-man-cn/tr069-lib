package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//Fault error response
type Fault struct {
	ID                      string
	Name                    string
	NoMore                  int
	FaultCode               string
	FaultString             string
	CwmpFaultCode           string
	CwmpFaultString         string
	SetParameterValuesFault SetParameterValuesFaultStruct
}

type faultBodyStruct struct {
	Fault faultStruct `xml:"SOAP-ENV:Fault"`
}
type faultStruct struct {
	FaultCode   string            `xml:"faultcode"`
	FaultString string            `xml:"faultstring"`
	FaultDetail faultDetailStruct `xml:"detail"`
}

type faultDetailStruct struct {
	CwmpFault cwmpFaultStruct `xml:"cwmp:Fault"`
}

type cwmpFaultStruct struct {
	FaultCode               string
	FaultString             string
	SetParameterValuesFault SetParameterValuesFaultStruct
}

//SetParameterValuesFaultStruct setParameterValues Fault
type SetParameterValuesFaultStruct struct {
	ParameterName string
	FaultCode     string
	FaultString   string
	ParameterKey  string
}

//NewFault create Fault object
func NewFault() (m *Fault) {
	m = &Fault{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetName get msg type
func (msg *Fault) GetName() string {
	return "Fault"
}

//GetID get msg id
func (msg *Fault) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *Fault) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id}
	setParamFault := SetParameterValuesFaultStruct{
		FaultCode:     msg.SetParameterValuesFault.FaultCode,
		FaultString:   msg.SetParameterValuesFault.FaultString,
		ParameterName: msg.SetParameterValuesFault.ParameterName,
		ParameterKey:  msg.SetParameterValuesFault.ParameterKey,
	}
	cwmpFault := cwmpFaultStruct{
		FaultCode:               msg.CwmpFaultCode,
		FaultString:             msg.CwmpFaultString,
		SetParameterValuesFault: setParamFault,
	}
	detail := faultDetailStruct{CwmpFault: cwmpFault}
	fault := faultStruct{
		FaultCode:   msg.FaultCode,
		FaultString: msg.FaultString,
		FaultDetail: detail,
	}
	env.Body = faultBodyStruct{fault}
	//output, err := xml.MarshalIndent(env, "  ", "    ")
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *Fault) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	faultNode := doc.SelectNode("*", "Fault")
	msg.FaultCode = faultNode.SelectNode("", "faultcode").GetValue()
	msg.FaultString = faultNode.SelectNode("", "faultstring").GetValue()
	detailNode := faultNode.SelectNode("*", "detail")
	detailFaultNode := detailNode.SelectNode("*", "Fault")
	msg.CwmpFaultCode = detailFaultNode.SelectNode("*", "FaultCode").GetValue()
	msg.CwmpFaultString = detailFaultNode.SelectNode("*", "FaultString").GetValue()
	setParaFaultNode := detailFaultNode.SelectNode("*", "SetParameterValuesFault")
	if setParaFaultNode != nil {
		msg.SetParameterValuesFault.FaultCode = setParaFaultNode.SelectNode("*", "FaultCode").GetValue()
		msg.SetParameterValuesFault.FaultString = setParaFaultNode.SelectNode("*", "FaultString").GetValue()
		msg.SetParameterValuesFault.ParameterName = setParaFaultNode.SelectNode("*", "ParameterName").GetValue()
		// Sometime the ParameterKey is not present on some devices
		pk := setParaFaultNode.SelectNode("*", "ParameterKey")
		if pk != nil {
		  msg.SetParameterValuesFault.ParameterKey = pk.GetValue()
		}
	}
	return nil
}
