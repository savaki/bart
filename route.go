package main

type RouteInfo struct {
	Uri      string  `xml:"uri" json:"uri"`
	SchedNum int     `xml:"sched_num" json:"sched_num"`
	Routes   []Route `xml:"routes>route" json:"routes"`
	Message  string  `xml:"message" json:"message,omitempty"`
}

type Route struct {
	Name         string   `xml:"name" json:"name"`
	Abbr         string   `xml:"abbr" json:"abbr"`
	RouteID      string   `xml:"routeID" json:"routeID"`
	Number       int      `xml:"number" json:"number"`
	Origin       string   `xml:"origin" json:"origin"`
	Destination  string   `xml:"destination" json:"destination"`
	Direction    string   `xml:"direction" json:"direction"`
	Color        string   `xml:"color" json:"color"`
	Holidays     int      `xml:"holidays" json:"holidays"`
	StationCount int      `xml:"num_stns" json:"num_stns"`
	Stations     []string `xml:"config>station" json:"stations"`
}
