package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type FlightRecord struct {
	CreatedAt       string `json:"created_at"`
	CreatedStr      string `json:"created_str"`
	FlightMissionId string `json:"flight_mission_id"`
	Id              string `json:"id"`
	LocalId         string `json:"local_id"`
	Message         string `json:"message"`
	RecordAt        string `json:"record_at"`
}
func (fr FlightRecord)String()string{
	return fmt.Sprintf(`{CreatedAt:"%v" CreatedStr:"%v" FlightMissionId:"%v" Id:"%v" LocalId:"%v" Message:"%v" RecordAt:"%v"}`,fr.CreatedAt,fr.CreatedStr,fr.FlightMissionId,fr.Id,fr.LocalId,fr.Message,fr.RecordAt)
}
func main() {
	data := ParseData(ReadFile("624670.json"))
	max := int64(-1)
	min := int64(-1)
	for _, v := range data {
		//fmt.Println(v)
		tmp := ParseTimestamp(v.CreatedStr)
		if max == -1 {
			max = tmp
		} else if max < tmp {
			max = tmp
		}
		if min == -1 {
			min = tmp
		} else if min > tmp {
			min = tmp
		}
	}
	fmt.Printf("%0.2f minutes\n", float64(max-min)/60000000000.0)
}
func ParseTimestamp(timestamp string) int64 {

	if strings.Count(timestamp, ":") > 2 {
		timestamp = regexp.MustCompile(":(\\d+)$").ReplaceAllString(timestamp, ".$1")
	}

	t, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		panic(err)
	}
	return t.UnixNano()
}
func ReadFile(filepath string) []byte {
	var tmp bytes.Buffer
	fp, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	tmp.ReadFrom(fp)
	return tmp.Bytes()
}
func ParseData(data []byte) (frs []FlightRecord) {
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&frs)
	if err != nil {
		panic(err)
	}
	return frs
}
