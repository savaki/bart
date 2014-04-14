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
			value := &Station{}
			err := xml.Unmarshal([]byte(content), value)
			So(err, ShouldBeNil)

			Convey("Then I expect the meta data to be set", func() {
				So(value.Uri, ShouldEqual, "http://api.bart.gov/api/etd.aspx?cmd=etd&orig=RICH")
				So(value.DateString, ShouldEqual, "03/30/2011")
				So(value.TimeString, ShouldEqual, "02:43:27 PM PDT")
				So(value.Name, ShouldEqual, "Richmond")
				So(value.Abbr, ShouldEqual, "RICH")
			})

			Convey("Then I expect the etds to be returned", func() {
				etd = value.ETD[0]
				So(len(value.ETD), ShouldEqual, 2)

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
}
