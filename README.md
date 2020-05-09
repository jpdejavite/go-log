# Go log

A golang lib to log with coi and metadata


## Usage

```
package main

import (
	"github.com/jpdejavite/go-log/pkg/log"
)

// Must have fields exported to appear in log
type Metadata struct {
	RequestID int
	UserData  UserMetadata
}

// Must have fields exported to appear in log
type UserMetadata struct {
	ID   int
	Name string
}

func main() {

	meta := Metadata{
		RequestID: 2423,
		UserData: UserMetadata{
			ID:   123123,
			Name: "My Name",
		},
	}
	tag := "MY_TAG"
	message := "erro in this line"
	coi := log.GenerateCoi(nil)
	log.Debug(tag, message, meta, coi)
	log.Error(tag, message, meta, coi)
	log.Info(tag, message, meta, coi)
	log.Fatal(tag, message, meta, coi)
	log.Warn(tag, message, meta, coi)
}

```
