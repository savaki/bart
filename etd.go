package main

type StationInfo struct {
	Uri        string  `xml:"uri" json:"uri"`
	DateString string  `xml:"date" json:"date"`
	TimeString string  `xml:"time" json:"time"`
	Station    Station `xml:"station" json:"station"`
}

type Station struct {
	Name           string   `xml:"name" json:"name"`
	Abbr           string   `xml:"abbr" json:"abbr"`
	ETD            []ETD    `xml:"etd" json:"etd,omitempty"`
	Lat            string   `xml:"gtfs_latitude" json:"lat,omitempty"`
	Long           string   `xml:"gtfs_longitude" json:"long,omitempty"`
	Address        string   `xml:"address" json:"address,omitempty"`
	City           string   `xml:"city" json:"city,omitempty"`
	County         string   `xml:"county" json:"county,omitempty"`
	State          string   `xml:"state" json:"state,omitempty"`
	Zip            string   `xml:"zipcode" json:"zipcode,omitempty"`
	NorthRoutes    []string `xml:"north_routes>route" json:"north_routes,omitempty"`
	SouthRoutes    []string `xml:"south_routes>route" json:"south_routes,omitempty"`
	NorthPlatforms []int    `xml:"north_platforms>platform" json:"north_platforms,omitempty"`
	SouthPlatforms []int    `xml:"south_platforms>platform" json:"south_platforms,omitempty"`
	PlatformInfo   string   `xml:"platform_info" json:"platform_info,omitempty"`
	Intro          string   `xml:"intro" json:"intro,omitempty"`
	CrossStreet    string   `xml:"cross_street" json:"cross_street,omitempty"`
	Food           string   `xml:"food" json:"food,omitempty"`
	Shopping       string   `xml:"shopping" json:"shopping,omitempty"`
	Attraction     string   `xml:"attraction" json:"attraction,omitempty"`
	Link           string   `xml:"link" json:"link,omitempty"`
}

type ETD struct {
	Destination string     `xml:"destination" json:"destination"`
	Abbr        string     `xml:"abbreviation" json:"abbreviation"`
	Estimates   []Estimate `xml:"estimate" json:"estimate"`
}

type Estimate struct {
	Minutes   string `xml:"minutes" json:"minutes"`
	Platform  int    `xml:"platform" json:"platform"`
	Direction string `xml:"direction" json:"direction"`
	Length    int    `xml:"length" json:"length"`
	Color     string `xml:"color" json:"color"`
	HexColor  string `xml:"hexcolor" json:"hexcolor"`
	BikeFlag  bool   `xml:"bikeflag" json:"bikeflag"`
}
