package common

import (
	"encoding/xml"
)

type EnvelopeRS[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	S       string   `xml:"xmlns:S,attr"`
	SoapEnv string   `xml:"xmlns:SOAP-ENV,attr"`
	Body    struct {
		XMLName xml.Name `xml:"S:Body"`
		Content T        `xml:",innerxml"`
	} `xml:"S:Body"`
}
