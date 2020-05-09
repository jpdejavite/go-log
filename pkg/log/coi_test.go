package log_test

import (
	"testing"
	"time"

	"github.com/jpdejavite/go-log/pkg/log"
)

func TestGenerateCoiDefaultConfig(t *testing.T) {
	got := log.GenerateCoi(nil)
	if got == "" || len(got) != 10 {
		t.Errorf("GenerateCoi default config error: %s", got)
	}
}

func TestGenerateCoiDefaultConfigRandoness(t *testing.T) {
	coi1 := log.GenerateCoi(nil)
	time.Sleep(time.Second)
	coi2 := log.GenerateCoi(nil)
	if coi1 == coi2 {
		t.Errorf("GenerateCoi randomness not working: %s", coi1)
	}
}

func TestGenerateCoiCustomConfig(t *testing.T) {
	coi := log.GenerateCoi(&log.CoiConfig{
		Charset: "A",
		Length:  20,
	})
	if coi != "AAAAAAAAAAAAAAAAAAAA" {
		t.Errorf("GenerateCoi custom config not working: %s", coi)
	}
}
