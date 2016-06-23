package messages

import (
	"encoding/xml"
	"fmt"
	"github.com/jteeuwen/go-pkg-xmlx"
	"strconv"
	"time"
)

//DownloadResponse download response
type DownloadResponse struct {
	ID           string
	Name         string
	Status       int
	StartTime    string
	CompleteTime string
}

type downloadResponseBodyStruct struct {
	DownResp downloadResponseStruct `xml:"cwmp:DownloadResponse"`
}

type downloadResponseStruct struct {
	Status       int
	StartTime    string
	CompleteTime string
}

//NewDownloadResponse create DownloadResponse object
func NewDownloadResponse() (m *DownloadResponse) {
	m = &DownloadResponse{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetID tr069 msg id
func (msg *DownloadResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName msg type name
func (msg *DownloadResponse) GetName() string {
	return "DownloadResponse"
}

//CreateXML encode into xml
func (msg *DownloadResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id}
	download := downloadResponseStruct{
		StartTime:    msg.StartTime,
		CompleteTime: msg.CompleteTime,
		Status:       msg.Status,
	}
	env.Body = downloadResponseBodyStruct{download}
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode into struct
func (msg *DownloadResponse) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	statusNode := doc.SelectNode("*", "Status")
	if statusNode != nil {
		var err error
		msg.Status, err = strconv.Atoi(statusNode.GetValue())
		if err != nil {
			return err
		}
	}
	msg.StartTime = doc.SelectNode("*", "StartTime").GetValue()
	msg.CompleteTime = doc.SelectNode("*", "CompleteTime").GetValue()
	return nil
}
