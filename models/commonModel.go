package models


type DeviceInfo struct {
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	OsVersion  string `json:"osVersion"`
	Platform   string `json:"platform"`
	Resolution string `json:"resolution"`
}
type ImageList struct {

	PartCode      int `json:"partCode"`
	ImageUrl      string `json:"imageUrl"`

}

