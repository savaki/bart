package main

type Stations struct {
	Uri      string    `xml:"uri" json:"uri"`
	Stations []Station `xml:"stations>station" json:"stations"`
}
