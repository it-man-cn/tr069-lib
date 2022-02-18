package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//GetParameterNames get paramnames
type GetParameterNames struct {
	ID     string
	Name   string
	NoMore int
	Body   GetParameterNamesStruct
}

type GetParameterNamesBodyStruct struct {
	Body GetParameterNamesStruct `xml:"cwmp:GetParameterNames"`
}

type GetParameterNamesStruct struct {
	ParameterPath string `xml:"ParameterPath"`
	NextLevel     int    `xml:"NextLevel"`
}

//NewGetParameterNames create GetParameterNames object
func NewGetParameterNames() (m *GetParameterNames) {
	m = &GetParameterNames{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetName get type name
func (msg *GetParameterNames) GetName() string {
	return "GetParameterNames"
}

//GetID get tr069 msg id
func (msg *GetParameterNames) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *GetParameterNames) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	env.Body = GetParameterNamesBodyStruct{msg.Body}
	output, err := xml.MarshalIndent(env, "  ", "    ")
	//output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *GetParameterNames) Parse(doc *xmlx.Document) error {
	//TODO
	return nil
}
