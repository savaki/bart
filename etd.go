package main

type Station struct {
	Uri        string `xml:"uri" json:"uri"`
	DateString string `xml:"date" json:"date"`
	TimeString string `xml:"time" json:"time"`
	Name       string `xml:"station>name" json:"name"`
	Abbr       string `xml:"station>abbr" json:"abbr"`
	ETD        []ETD  `xml:"station>etd" json:"etd"`
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
