package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//Download type
const (
	FTFireware   string = "1 Firmware Upgrade Image"
	FTWebContent string = "2 Web Content"
	FTConfig     string = "3 Vendor Configuration File"
)

//Download tr069 download msg
type Download struct {
	ID             string
	Name           string
	NoMore         int
	CommandKey     string
	FileType       string
	URL            string
	Username       string
	Password       string
	FileSize       int
	TargetFileName string
	DelaySeconds   int
	SuccessURL     string
	FailureURL     string
}

type downloadBodyStruct struct {
	Body downloadStruct `xml:"cwmp:Download"`
}

type downloadStruct struct {
	CommandKey     string
	FileType       string
	URL            string
	Username       string
	Password       string
	FileSize       int
	TargetFileName string
	DelaySeconds   int
	SuccessURL     string
	FailureURL     string
}

func NewDownload() *Download {
	download := new(Download)
	download.ID = download.GetID()
	download.Name = download.GetName()
	return download
}

//GetID get download msg id(tr069 msg id)
func (msg *Download) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName name is msg object type, use for decode
func (msg *Download) GetName() string {
	return "Download"
}

//CreateXML encode xml
func (msg *Download) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	download := downloadStruct{
		CommandKey:     msg.CommandKey,
		FileType:       msg.FileType,
		URL:            msg.URL,
		Username:       msg.Username,
		Password:       msg.Password,
		FileSize:       msg.FileSize,
		TargetFileName: msg.TargetFileName,
		DelaySeconds:   msg.DelaySeconds,
		SuccessURL:     msg.SuccessURL,
		FailureURL:     msg.FailureURL}
	env.Body = downloadBodyStruct{download}
	//output, err := xml.Marshal(env)
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse parse from xml
func (msg *Download) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	msg.CommandKey = doc.SelectNode("*", "CommandKey").GetValue()
	msg.FileType = doc.SelectNode("*", "FileType").GetValue()
	msg.URL = doc.SelectNode("*", "URL").GetValue()
	return nil
}
