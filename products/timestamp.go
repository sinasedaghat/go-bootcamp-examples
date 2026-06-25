package products

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts *timestamp) String() string {
	if ts == nil || ts.IsZero() {
		return "unknown"
	}

	const layout = "2006/01"
	return ts.Format(layout)
}

func toTimestamp(v any) (ts *timestamp) {
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
	default:
		t = 0
	}

	ts = &timestamp{Time: time.Unix(int64(t), 0)}
	return ts
}

func (ts *timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(ts.Unix(), 10)), nil
	// if ts.IsZero() {
	// 	return []byte("nill"), nil
	// } else {
	// 	return []byte(strconv.FormatInt(ts.Unix(), 10)), nil
	// }
}

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	ts = toTimestamp(string(data))
	return nil
}
