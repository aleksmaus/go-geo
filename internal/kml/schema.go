package kml

import "encoding/xml"

type KMLBase struct {
	XMLName  xml.Name `xml:"kml"`
	XMLNS    string   `xml:"xmlns,attr"`
	Document Document `xml:"Document"`
}

// Different structure for marshalling/unmarshalling due to assymetry of namespace attributes decoding/encoding
type KML struct {
	KMLBase
	XMLNSGX string `xml:"xmlns gx,attr"`
}

type KMLOut struct {
	KMLBase
	XMLNSGX string `xml:"xmlns:gx,attr"`
}

type Document struct {
	XMLName    xml.Name    `xml:"Document"`
	Name       string      `xml:"name"`
	StyleMaps  []StyleMap  `xml:"StyleMap"`
	Styles     []Style     `xml:"Style"`
	Placemarks []Placemark `xml:"Placemark"`
}

type StyleMap struct {
	XMLName xml.Name `xml:"StyleMap"`
	ID      string   `xml:"id,attr"`
	Pair    string   `xml:",innerxml"`
}

type Style struct {
	XMLName struct{} `xml:"Style"`
	ID      string   `xml:"id,attr"`
	Style   string   `xml:",innerxml"`
}

type Placemark struct {
	XMLName struct{} `xml:"Placemark"`
	Style   string   `xml:",innerxml"`
}
