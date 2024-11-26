package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//AddObject adds object to cpe
type DeleteObject struct {
	ID           string
	Name         string
	NoMore       int
	ObjectName   string
	ParameterKey string
}

type delObjectBodyStruct struct {
	Body delObjectStruct `xml:"cwmp:DeleteObject"`
}

type delObjectStruct struct {
	ObjectName   string
	ParameterKey string
}

//GetID get msg id
func (msg *DeleteObject) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg name
func (msg *DeleteObject) GetName() string {
	return "DeleteObject"
}

//CreateXML encode into xml
func (msg *DeleteObject) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	delObject := delObjectStruct{ObjectName: msg.ObjectName, ParameterKey: msg.ParameterKey}
	env.Body = delObjectBodyStruct{delObject}
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *DeleteObject) Parse(doc *xmlx.Document) error {
	//TODO
	return nil
}
