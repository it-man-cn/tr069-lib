package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//AddObject adds object to cpe
type AddObject struct {
	ID           string
	Name         string
	NoMore       int
	ObjectName   string
	ParameterKey string
}

type addObjectBodyStruct struct {
	Body addObjectStruct `xml:"cwmp:AddObject"`
}

type addObjectStruct struct {
	ObjectName   string
	ParameterKey string
}

//GetID get msg id
func (msg *AddObject) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg name
func (msg *AddObject) GetName() string {
	return "AddObject"
}

//CreateXML encode into xml
func (msg *AddObject) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	addObject := addObjectStruct{ObjectName: msg.ObjectName, ParameterKey: msg.ParameterKey}
	env.Body = addObjectBodyStruct{addObject}
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *AddObject) Parse(doc *xmlx.Document) error {
	//TODO
	return nil
}
