package main

import (
	"fmt"

	"github.com/matrix-org/gomatrixserverlib"
)

var str string = `{"body":"video_5e1f43c.mp4","info":{"duration":30466.666666666664,"thumbnail_url":"mxc:\/\/localhost\/e76b28e344ae8c388cc19e42fb7eb73d397b1931001a0fd62f26f199d1b3cd1a","mimetype":"video\/mp4","w":480,"h":270,"thumbnail_info":{"size":0,"mimetype":"image\/jpeg","w":0,"h":0}},"url":"mxc:\/\/localhost\/55e9366d7a68d7970b8071145d631745d36231a4f2eeccab366a61f8ef521626","msgtype":"m.video"}`

var bytes = []byte(str)

var v6 gomatrixserverlib.RoomVersion = "6"
var v5 gomatrixserverlib.RoomVersion = "5"

func logTestCase(str string) {
	fmt.Printf("\n")
	fmt.Printf(str)
	fmt.Printf("\n")
}

func logResult(json []byte, err error) {
	fmt.Printf("\n")
	fmt.Printf("JSON:")
	fmt.Printf("\n")
	fmt.Printf("%v", json)
	fmt.Printf("\n")
	fmt.Printf("Error:")
	fmt.Printf("\n")
	fmt.Printf("%v", err)
}

func validateCanonicalJSON() {
	logTestCase("Validating CanonicalJSON()...")
	var json, err = gomatrixserverlib.CanonicalJSON(bytes)
	logResult(json, err)
}

func validateEnforcedJSONRoomV6() {
	logTestCase("Validating EnforcedCanonicalJSON() with RoomVersion 6...")
	var json, err = gomatrixserverlib.EnforcedCanonicalJSON(bytes, v6)
	logResult(json, err)
}

func validateEnforcedJSONRoomV5() {
	logTestCase("Validating EnforcedCanonicalJSON() with RoomVersion 5...")
	var json, err = gomatrixserverlib.EnforcedCanonicalJSON(bytes, v5)
	logResult(json, err)
}

func main() {
	validateCanonicalJSON()
	validateEnforcedJSONRoomV6()
	validateEnforcedJSONRoomV5()
}
