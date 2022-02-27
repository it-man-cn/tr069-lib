package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//Upload type
const (
	//FTConfig     string = "3 Vendor Configuration File"
	FTLog string = "4 Vendor Log File 1"
)

//Upload tr069 upload msg
type Upload struct {
	ID           string
	Name         string
	NoMore       int
	CommandKey   string
	FileType     string
	URL          string
	Username     string
	Password     string
	DelaySeconds int
}

type uploadBodyStruct struct {
	Body uploadStruct `xml:"cwmp:Upload"`
}

type uploadStruct struct {
	CommandKey   string
	FileType     string
	URL          string
	Username     string
	Password     string
	DelaySeconds int
}

//NewUpload create a inform messages
func NewUpload() *Upload {
	upload := new(Upload)
	upload.ID = upload.GetID()
	upload.Name = upload.GetName()
	return upload
}

//GetID get upload msg id(tr069 msg id)
func (msg *Upload) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName name is msg object type, use for decode
func (msg *Upload) GetName() string {
	return "Upload"
}

//CreateXML encode xml
func (msg *Upload) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	upload := uploadStruct{
		CommandKey:   msg.CommandKey,
		FileType:     msg.FileType,
		URL:          msg.URL,
		Username:     msg.Username,
		Password:     msg.Password,
		DelaySeconds: msg.DelaySeconds,
	}
	env.Body = uploadBodyStruct{upload}
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse parse from xml
func (msg *Upload) Parse(doc *xmlx.Document) error {
	//TODO
	return nil
}
