package messages

import (
	"encoding/xml"
	"fmt"
	"time"
	"strconv"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//ScheduleInform get paramnames
type ScheduleInform struct {
	ID     string
	Name   string
	NoMore int
	Data   ScheduleInformStruct
}

type ScheduleInformBodyStruct struct {
	Body ScheduleInformStruct `xml:"cwmp:ScheduleInform"`
}

type ScheduleInformStruct struct {
	DelaySeconds uint   `xml:"DelaySeconds"`
	CommandKey   string `xml:"CommandKey"`
}

//NewScheduleInform create GetParameterNames object
func NewScheduleInformStruct() (m *ScheduleInform) {
	m = &ScheduleInform{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetName get type name
func (msg *ScheduleInform) GetName() string {
	return "ScheduleInformStruct"
}

//GetID get tr069 msg id
func (msg *ScheduleInform) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *ScheduleInform) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	getParam := ScheduleInformStruct{DelaySeconds: msg.Data.DelaySeconds, CommandKey: msg.Data.CommandKey}
	env.Body = ScheduleInformBodyStruct{getParam}
	output, err := xml.MarshalIndent(env, "  ", "    ")
	//output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *ScheduleInform) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	delaySecs := doc.SelectNode("*", "DelaySeconds").GetValue()
	d,_ := strconv.Atoi(delaySecs)
	msg.Data.DelaySeconds = uint(d)
	msg.Data.CommandKey = doc.SelectNode("*", "CommandKey").GetValue()
	return nil
}
