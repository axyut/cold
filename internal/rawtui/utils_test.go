package rawtui_test

import (
	"testing"
	"time"

	"github.com/axyut/playgo/internal/rawtui"
)

func TestFmtDuration(t *testing.T) {

	samples := map[time.Duration]string{
		time.Second * 5:                "00:05",
		time.Hour * 2:                  "02:00:00",
		time.Minute*4 + time.Second*15: "04:15",
		time.Minute * 0:                "00:00",
		time.Millisecond * 5:           "00:00",
	}

	for k, v := range samples {

		got := rawtui.FmtDuration(k)

		if got != v {
			t.Errorf("fmtDuration(%s); Expected %s got %s", k, v, got)
		}

	}

}
