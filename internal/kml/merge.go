package kml

func Merge(kdoc1, kdoc2 *KMLOut) *KMLOut {
	if kdoc1 == nil {
		return kdoc2
	}

	// merge Placemarks
	kdoc1.Document.Placemarks = append(kdoc1.Document.Placemarks, kdoc2.Document.Placemarks...)
	return kdoc1
}
