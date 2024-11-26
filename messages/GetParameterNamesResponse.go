package messages

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

type ParameterNameStruct struct {
	Name     string
	Writable bool
}

//GetParameterNamesResponse getParameterNames response
type GetParameterNamesResponse struct {
	ID    string
	Name  string
	Names []ParameterNameStruct
}

//NewGetParameterNamesResponse create GetParameterNamesResponse object
func NewGetParameterNamesResponse() (m *GetParameterNamesResponse) {
	m = &GetParameterNamesResponse{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

type getParameterNamesResponseBodyStruct struct {
	Body getParameterNamesResponseStruct `xml:"cwmp:GetParameterNamesResponse"`
}

type getParameterNamesResponseStruct struct {
	Params ParameterListStruct `xml:"ParameterList"`
}

//GetName get type name
func (msg *GetParameterNamesResponse) GetName() string {
	return "GetParameterNamesResponse"
}

//GetID get msg id
func (msg *GetParameterNamesResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *GetParameterNamesResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id}

	paramLen := strconv.Itoa(len(msg.Names))
	params := ParameterListStruct{Type: "cwmp:ParameterInfoStruct[" + paramLen + "]"}
	for _, v := range msg.Names {
		s := "false"
		if v.Writable {
			s = "true"
		}
		param := ParameterInfoStruct{
			Name:     NodeStruct{Type: XsdString, Value: v.Name},
			Writable: NodeStruct{Type: XsdString, Value: s}}
		params.Names = append(params.Names, param)
	}
	info := getParameterNamesResponseStruct{Params: params}
	env.Body = getParameterNamesResponseBodyStruct{info}
	output, err := xml.MarshalIndent(env, "  ", "    ")
	//output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *GetParameterNamesResponse) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	paramsNode := doc.SelectNode("*", "ParameterList")
	if len(strings.TrimSpace(paramsNode.String())) > 0 {
		params := make([]ParameterNameStruct, 0)
		var name string
		var writable bool
		for _, param := range paramsNode.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name = param.SelectNode("", "Name").GetValue()
				writable = false
				var writeableVal = param.SelectNode("", "Writable").GetValue()
				if writeableVal == "true" || writeableVal == "1" {
					writable = true
				}
				params = append(params, ParameterNameStruct{Name: name, Writable: writable})
			}

		}
		msg.Names = params
	}
	return nil
}
