package messages

import (
	"github.com/jteeuwen/go-pkg-xmlx"
)

//ParseXML parse xml msg
func ParseXML(data []byte) (msg Message, err error) {
	doc := xmlx.New()
	err = doc.LoadBytes(data, nil)
	if err != nil {
		return
	}
	bodyNode := doc.SelectNode("*", "Body")
	if bodyNode != nil {
		var name string
		if len(bodyNode.Children) > 1 {
			name = bodyNode.Children[1].Name.Local
		} else {
			name = bodyNode.Children[0].Name.Local
		}
		switch name {
		case "Inform":
			msg = NewInform()
			err = msg.Parse(doc)
		case "GetParameterValuesResponse":
			msg = NewGetParameterValuesResponse()
			err = msg.Parse(doc)
		case "SetParameterValuesResponse":
			msg = NewSetParameterValuesResponse()
			err = msg.Parse(doc)
		case "DownloadResponse":
			msg = NewDownloadResponse()
			err = msg.Parse(doc)
		case "TransferComplete":
			msg = NewTransferComplete()
			err = msg.Parse(doc)
		case "GetRPCMethodsResponse":
			msg = NewGetRPCMethodsResponse()
			err = msg.Parse(doc)
		case "RebootResponse":
			msg = NewRebootResponse()
			err = msg.Parse(doc)
		case "Fault":
			msg = NewFault()
			err = msg.Parse(doc)
		}
	}
	return msg, err
}
