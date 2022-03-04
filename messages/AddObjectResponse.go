package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//AddObjectResponse add object response
type AddObjectResponse struct {
	ID   string
	Name string
}

type addObjectResponseBodyStruct struct {
	Body addObjectResponseStruct `xml:"cwmp:AddObjectResponse"`
}

type addObjectResponseStruct struct {
}

//NewAddObjectResponse create AddObjectResponse object
func NewAddObjectResponse() (m *AddObjectResponse) {
	m = &AddObjectResponse{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetID get msg id
func (msg *AddObjectResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg type
func (msg *AddObjectResponse) GetName() string {
	return "AddObjectResponse"
}

//CreateXML encode into xml
func (msg *AddObjectResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id}
	addobj := addObjectResponseStruct{}
	env.Body = addObjectResponseBodyStruct{addobj}
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *AddObjectResponse) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	return nil
}
