package kml

import (
	"encoding/xml"
	"os"
)

func DecodeFile(fpath string) (kmlDoc KML, err error) {
	f, err := os.Open(fpath)
	if err != nil {
		return
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	err = dec.Decode(&kmlDoc)
	return kmlDoc, err
}
