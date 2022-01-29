package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//GetRPCMethods get rpc methods
type GetRPCMethods struct {
	ID     string
	Name   string
	NoMore int
}

type getRPCMethodsBodyStruct struct {
	Body getRPCMethodsStruct `xml:"cwmp:GetRPCMethods"`
}

type getRPCMethodsStruct struct {
}

//NewGetRPCMethods create GetRPCMethods object
func NewGetRPCMethods() (m *GetRPCMethods) {
	m = &GetRPCMethods{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetID get msg id
func (msg *GetRPCMethods) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get type name
func (msg *GetRPCMethods) GetName() string {
	return "GetRPCMethods"
}

//CreateXML encode into xml
func (msg *GetRPCMethods) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	rpcMethods := getRPCMethodsStruct{}
	env.Body = getRPCMethodsBodyStruct{rpcMethods}
	//output, err := xml.Marshal(env)
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *GetRPCMethods) Parse(doc *xmlx.Document) error {
	//TODO
	return nil
}
