package main

import (
	"encoding/xml"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRouteInfo(t *testing.T) {
	var route Route

	Convey("Given an xml routeinfo response", t, func() {
		content := `<?xml version="1.0" encoding="utf-8"?>
		<root><uri><![CDATA[http://api.bart.gov/api/route.aspx?cmd=routeinfo&route=6]]></uri><sched_num>34</sched_num>
		<routes><route><name>Daly City - Fremont</name><abbr>DALY-FRMT</abbr><routeID>ROUTE 6</routeID><number>6</number>
		<origin>DALY</origin><destination>FRMT</destination><direction></direction><color>#339933</color><holidays>0</holidays>
		<num_stns>19</num_stns>
		<config><station>DALY</station><station>BALB</station><station>GLEN</station><station>24TH</station><station>16TH</station><station>CIVC</station><station>POWL</station><station>MONT</station><station>EMBR</station><station>WOAK</station><station>LAKE</station><station>FTVL</station><station>COLS</station><station>SANL</station><station>BAYF</station><station>HAYW</station><station>SHAY</station><station>UCTY</station><station>FRMT</station></config></route></routes><message /></root>
		`

		Convey("When I unmarshall it", func() {
			value := &RouteInfo{}
			err := xml.Unmarshal([]byte(content), value)
			So(err, ShouldBeNil)

			Convey("Then I expect the meta data to be set", func() {
				So(value.Uri, ShouldEqual, "http://api.bart.gov/api/route.aspx?cmd=routeinfo&route=6")
				So(value.SchedNum, ShouldEqual, 34)
			})

			Convey("Then I expect the routes to be returned", func() {
				route = value.Routes[0]
				So(len(value.Routes), ShouldEqual, 1)

				So(route.Name, ShouldEqual, "Daly City - Fremont")
				So(route.Abbr, ShouldEqual, "DALY-FRMT")
				So(route.RouteID, ShouldEqual, "ROUTE 6")
				So(route.Origin, ShouldEqual, "DALY")
				So(route.Destination, ShouldEqual, "FRMT")
				So(route.Direction, ShouldEqual, "")
				So(route.Color, ShouldEqual, "#339933")
				So(route.Holidays, ShouldEqual, 0)
				So(route.StationCount, ShouldEqual, 19)
				So(route.Number, ShouldEqual, 6)
			})

			Convey("Then I expect the stations to be returned", func() {
				stations := []string{"DALY", "BALB", "GLEN", "24TH", "16TH", "CIVC", "POWL", "MONT", "EMBR", "WOAK", "LAKE", "FTVL", "COLS", "SANL", "BAYF", "HAYW", "SHAY", "UCTY", "FRMT"}
				So(value.Routes[0].Stations, ShouldResemble, stations)
			})
		})
	})
}
