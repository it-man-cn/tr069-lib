package messages

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//SetParameterValues set param
type SetParameterValues struct {
	ID           string
	Name         string
	NoMore       int
	Params       map[string]ValueStruct
	ParameterKey string
}

type setParameterValuesBodyStruct struct {
	Body setParameterValuesStruct `xml:"cwmp:SetParameterValues"`
}

type setParameterValuesStruct struct {
	ParamList    ParameterListStruct `xml:"ParameterList"`
	ParameterKey string
}

//NewSetParameterValues create SetParameterValues object
func NewSetParameterValues() (m *SetParameterValues) {
	m = &SetParameterValues{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetName get msg type
func (msg *SetParameterValues) GetName() string {
	return "SetParameterValues"
}

//GetID get msg id
func (msg *SetParameterValues) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *SetParameterValues) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}

	paramLen := strconv.Itoa(len(msg.Params))
	paramList := ParameterListStruct{Type: "cwmp:ParameterValueStruct[" + paramLen + "]"}
	for k, v := range msg.Params {
		param := ParameterValueStruct{
			Name:  NodeStruct{Value: k},
			Value: NodeStruct{Type: v.Type, Value: v.Value}}
		paramList.Params = append(paramList.Params, param)
	}
	setParam := setParameterValuesStruct{
		ParamList:    paramList,
		ParameterKey: msg.ParameterKey,
	}
	env.Body = setParameterValuesBodyStruct{setParam}
	//output, err := xml.MarshalIndent(env, "  ", "    ")
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *SetParameterValues) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	paramsNode := doc.SelectNode("*", "ParameterList")
	if len(strings.TrimSpace(paramsNode.String())) > 0 {
		params := make(map[string]ValueStruct)
		var name, value string
		for _, param := range paramsNode.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name = param.SelectNode("", "Name").GetValue()
				value = param.SelectNode("", "Value").GetValue()
				params[name] = ValueStruct{Value: value}
			}
		}
		msg.Params = params
	}
	return nil
}
