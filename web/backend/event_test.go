package backend

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEvent(t *testing.T) {
	events := EventList{Items: []Event{
		{Name: "foo", Type: EventType("foo"), Metrics: EventMetrics{
			Count: 1,
		}},
		{Name: "bar", Type: EventType("bar"), Metrics: EventMetrics{
			Count: 1,
		}},
	}}
	raw, err := json.Marshal(events)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(raw))
}
