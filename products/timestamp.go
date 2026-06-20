package products

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "Unknown"
	}

	const layout = "2006/01"
	return ts.Format(layout)
}

func toTimestamp(v any) (ts timestamp) {
	var t int
	switch v := v.(type) {
	case string:
		t, _ = strconv.Atoi(v)
	case int:
		t = v
	case float64:
		t = int(v)
	case float32:
		t = int(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
