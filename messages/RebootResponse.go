package messages

import (
	"encoding/xml"
	"fmt"
	"github.com/jteeuwen/go-pkg-xmlx"
	"time"
)

//RebootResponse reboot response
type RebootResponse struct {
	ID   string
	Name string
}

type rebootResponseBodyStruct struct {
	Body rebootResponseStruct `xml:"cwmp:TransferCompleteResponse"`
}

type rebootResponseStruct struct {
}

//NewRebootResponse create RebootResponse object
func NewRebootResponse() (m *RebootResponse) {
	m = &RebootResponse{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetID get msg id
func (msg *RebootResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg type
func (msg *RebootResponse) GetName() string {
	return "RebootResponse"
}

//CreateXML encode into xml
func (msg *RebootResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id}
	reboot := rebootResponseStruct{}
	env.Body = rebootResponseBodyStruct{reboot}
	//output, err := xml.Marshal(env)
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *RebootResponse) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	return nil
}
