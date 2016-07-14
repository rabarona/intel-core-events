package main

import (
	"os"
	"encoding/json"
	"fmt"
)

// Implements routines for reading core events
// from CPUFamily_core.json and creates a new list
// of r+unmask+event for further monitoring

type Event struct{
	EventCode string `json:"EventCode"`
	UMask string `json:"UMask"`
	EventName string `json:"EventName"`
	BriefDescription string `json:"BriefDescription"`
	PublicDescription string `json:"PublicDescription"`
	Counter int `json:"Counter"`
	CounterHTOff int `json:"CounterHTOff"`
	SampleAfterValue int `json:"SampleAfterValue"`
	MSRIndex int `json:"MSRIndex"`
	MSRValue int `json:"MSRValue"`
	TakenAlone int `json:"TakenAlone"`
	CounterMask int `json:"CounterMask"`
	Invert int `json:"Invert"`
	AnyThread int `json:"AnyThread"`
	EdgeDetect int `json:"EdgeDetect"`
	PEBS int `json:"PEBS"`
	Data_LA int `json:"Data_LA"`
	L1_Hit_Indication int `json:"L1_Hit_Indication"`
	Errata int `json:"Errata"`
	Offcore int `json:"Offcore"`
}

func main() {
	const (
		eventListFileLocation = "/home/rabarona/Go/src/user/telemetry/data/Haswell_core_V21.json"
	)

	eventListFile, err := os.Open(eventListFileLocation)
	if err != nil {
		fmt.Print("openenig events file", err.Error())
	}

	data := make([]byte, 100)
	eventListFile.Read(data)
	events := make([]Event, 10)


	if err = json.Unmarshal(data, &events); err != nil {
		fmt.Print("parsing events file", err.Error())
	}

	fmt.Print(events)
	eventListFile.Close()
	return
}
