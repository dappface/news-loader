package model

import "encoding/json"

type Coordinates struct {
	Coordinates     []float64 `firestore:"coordinates"     json:"coordinates"` // [longitude, latitude]
	CoordinatesType string    `firestore:"coordinatesType" json:"coordinates_type"`
}

type Place struct {
	ID          string          `firestore:"id"          json:"id"`
	URL         string          `firestore:"url"         json:"url"`
	PlaceType   string          `firestore:"placeType"   json:"place_type"`
	Name        string          `firestore:"name"        json:"name"`
	FullName    string          `firestore:"fullName"    json:"full_name"`
	CountryCode string          `firestore:"countryCode" json:"country_code"`
	Country     string          `firestore:"country"     json:"country"`
	BoundingBox json.RawMessage `firestore:"boundingBox" json:"bounding_box"`
	// [Unused] Attributes  Attributes
}

// type BoundingBox struct {
// 	Coordinates [][][]float64 `firestore:"coordinates" json:"coordinates"` // [longitude, latitude]
// }
