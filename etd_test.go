package main

import (
	"encoding/xml"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStationInfo(t *testing.T) {
	var etd ETD

	Convey("Given an xml routeinfo response", t, func() {
		content := `<?xml version="1.0" encoding="utf-8" ?> 
<root>
  <uri><![CDATA[http://api.bart.gov/api/etd.aspx?cmd=etd&orig=RICH]]></uri>
  <date>03/30/2011</date> 
  <time>02:43:27 PM PDT</time> 
  <station>
    <name>Richmond</name> 
    <abbr>RICH</abbr> 
    <etd>
      <destination>Fremont</destination> 
      <abbreviation>FRMT</abbreviation> 
      <estimate>
        <minutes>5</minutes> 
        <platform>2</platform> 
        <direction>South</direction> 
        <length>6</length> 
        <color>ORANGE</color> 
        <hexcolor>#ff9933</hexcolor> 
        <bikeflag>1</bikeflag> 
      </estimate>
      <estimate>
        <minutes>20</minutes> 
        <platform>2</platform> 
        <direction>South</direction> 
        <length>6</length> 
        <color>ORANGE</color> 
        <hexcolor>#ff9933</hexcolor> 
        <bikeflag>1</bikeflag> 
      </estimate>
    </etd>
    <etd>
      <destination>Millbrae</destination> 
      <abbreviation>MLBR</abbreviation> 
      <estimate>
        <minutes>Leaving</minutes> 
        <platform>2</platform> 
        <direction>South</direction> 
        <length>10</length> 
        <color>RED</color> 
        <hexcolor>#ff0000</hexcolor> 
        <bikeflag>1</bikeflag> 
      </estimate>
    </etd>
  </station>
  <message /> 
</root>`

		Convey("When I unmarshall it", func() {
			value := &StationInfo{}
			err := xml.Unmarshal([]byte(content), value)
			So(err, ShouldBeNil)

			Convey("Then I expect the meta data to be set", func() {
				So(value.Uri, ShouldEqual, "http://api.bart.gov/api/etd.aspx?cmd=etd&orig=RICH")
				So(value.DateString, ShouldEqual, "03/30/2011")
				So(value.TimeString, ShouldEqual, "02:43:27 PM PDT")
			})

			Convey("And I expect the station info to be set", func() {
				station := value.Station

				So(station.Name, ShouldEqual, "Richmond")
				So(station.Abbr, ShouldEqual, "RICH")

				Convey("Then I expect the etds to be returned", func() {
					etd = station.ETD[0]
					So(len(station.ETD), ShouldEqual, 2)

					So(etd.Destination, ShouldEqual, "Fremont")
					So(etd.Abbr, ShouldEqual, "FRMT")
					So(len(etd.Estimates), ShouldEqual, 2)

					estimate := etd.Estimates[0]
					So(estimate.Minutes, ShouldEqual, "5")
					So(estimate.Platform, ShouldEqual, 2)
					So(estimate.Direction, ShouldEqual, "South")
					So(estimate.Length, ShouldEqual, 6)
					So(estimate.Color, ShouldEqual, "ORANGE")
					So(estimate.HexColor, ShouldEqual, "#ff9933")
					So(estimate.BikeFlag, ShouldBeTrue)
				})
			})
		})
	})

	Convey("Given stations content", t, func() {
		content := `<?xml version="1.0" encoding="utf-8"?>
	<root>
	<uri><![CDATA[http://api.bart.gov/api/stn.aspx?cmd=stninfo&orig=24th]]></uri>
	<stations>
		<station>
			<name>24th St. Mission</name>
			<abbr>24TH</abbr>
			<gtfs_latitude>37.752254</gtfs_latitude>
			<gtfs_longitude>-122.418466</gtfs_longitude>
			<address>2800 Mission Street</address>
			<city>San Francisco</city>
			<county>sanfrancisco</county>
			<state>CA</state>
			<zipcode>94110</zipcode>
			<north_routes><route>ROUTE 2</route><route> ROUTE 6</route><route> ROUTE 8</route><route> ROUTE 12</route></north_routes>
			<south_routes><route>ROUTE 1</route><route> ROUTE 5</route><route> ROUTE 7</route><route> ROUTE 11</route></south_routes>
			<north_platforms><platform>2</platform></north_platforms>
			<south_platforms><platform>1</platform></south_platforms>
			<platform_info>Always check destination signs and listen for departure announcements.</platform_info>
			<intro><![CDATA["The Mission" refers to the San Francisco de Asis Mission, also known as Mission Dolores, which was founded 1776. Today the neighborhood is host to an eclectic mix of restaurants, markets, performance spaces, shops, and nightspots.]]></intro>
			<cross_street><![CDATA[Nearby Cross: 24th St.]]></cross_street>
			<food><![CDATA[Nearby restaurant reviews from <a rel="external" href="http://www.yelp.com/search?find_desc=Restaurant+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>]]></food>
			<shopping><![CDATA[Local shopping from <a rel="external" href="http://www.yelp.com/search?find_desc=Shopping+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>]]></shopping>
			<attraction><![CDATA[More station area attractions from <a rel="external" href="http://www.yelp.com/search?find_desc=+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>]]></attraction>
			<link><![CDATA[http://www.bart.gov/stations/24TH/index.aspx]]></link>
		</station>
	</stations>
	<message></message>
</root>
`

		Convey("When I unmarshall it", func() {
			value := &Stations{}
			err := xml.Unmarshal([]byte(content), value)
			So(err, ShouldBeNil)

			Convey("Then I expect the meta data to be set", func() {
				So(value.Uri, ShouldEqual, "http://api.bart.gov/api/stn.aspx?cmd=stninfo&orig=24th")
			})

			Convey("And I expect the station info to be set", func() {
				station := value.Stations[0]

				So(station.Name, ShouldEqual, "24th St. Mission")
				So(station.Abbr, ShouldEqual, "24TH")

				Convey("Then I expect detailed station info to be returned", func() {
					So(station.Lat, ShouldEqual, "37.752254")
					So(station.Long, ShouldEqual, "-122.418466")
					So(station.Address, ShouldEqual, "2800 Mission Street")
					So(station.City, ShouldEqual, "San Francisco")
					So(station.County, ShouldEqual, "sanfrancisco")
					So(station.State, ShouldEqual, "CA")
					So(station.Zip, ShouldEqual, "94110")
					So(station.NorthRoutes, ShouldContain, "ROUTE 2")
					So(station.SouthRoutes, ShouldContain, "ROUTE 1")
					So(station.NorthPlatforms, ShouldContain, 2)
					So(station.SouthPlatforms, ShouldContain, 1)
					So(station.PlatformInfo, ShouldContainSubstring, "Always check")
					So(station.Intro, ShouldContainSubstring, "The Mission")
					So(station.CrossStreet, ShouldContainSubstring, "Nearby Cross")
					So(station.Food, ShouldContainSubstring, "Nearby restaurant reviews")
					So(station.Shopping, ShouldContainSubstring, "Local shopping")
					So(station.Attraction, ShouldContainSubstring, "More station area attractions")
					So(station.Link, ShouldContainSubstring, "http://www.bart.gov/stations/24TH/index.aspx")
				})
			})
		})
	})
}
