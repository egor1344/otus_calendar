package calendar

import (
	"github.com/golang/protobuf/ptypes"
	"testing"
)

func TestCalendar_AddEvent(t *testing.T) {
	calendar := Calendar{}
	calendar.Init()

	newEvent := Event{Id: 1, Date: ptypes.TimestampNow(), Title: "test", Description: "Description"}
	err := calendar.AddEvent(&newEvent)
	if err != nil {
		t.Error(err)
	}
	if calendar.idSeq != 1 {
		t.Error("Id sequence in calendar, dont increment")
	}
	if len(calendar.Events) != 1 {
		t.Error("Event dont append in calendar")
	}
	newEvent = Event{Id: 2, Date: ptypes.TimestampNow(), Title: "test", Description: "Description"}
	err = calendar.AddEvent(&newEvent)
	if err != nil {
		t.Error(err)
	}
	if calendar.idSeq != 2 {
		t.Error("Id sequence in calendar, dont increment")
	}
	if len(calendar.Events) != 2 {
		t.Error("Event dont append in calendar")
	}
	err = calendar.AddEvent(&newEvent)
	if err == nil {
		t.Error("Add event with already exists id")
	}
}

func TestCalendar_UpdateEvent(t *testing.T) {
	calendar := Calendar{}
	calendar.Init()
	newEvent := Event{Id: 1, Date: ptypes.TimestampNow(), Title: "test", Description: "Description"}
	err := calendar.AddEvent(&newEvent)
	if err != nil {
		t.Error(err)
	}
	if calendar.idSeq != 1 {
		t.Error("Id sequence in calendar, dont increment")
	}
	if len(calendar.Events) != 1 {
		t.Error("Event dont append in calendar")
	}
	newEvent = Event{Id: 2, Date: ptypes.TimestampNow(), Title: "test2", Description: "Description"}
	err = calendar.UpdateEvent(1, &newEvent)
	if err != nil {
		t.Error(err)
	}
	if len(calendar.Events) != 1 {
		t.Error("Event append in calendar when update func")
	}
	if calendar.Events[1].Title != "test2" {
		t.Error("Wrong update event in calendar")
	}

}

func TestCalendar_DeleteEvent(t *testing.T) {
	calendar := Calendar{}
	calendar.Init()
	newEvent := Event{Id: 1, Date: ptypes.TimestampNow(), Title: "test2", Description: "Description"}
	err := calendar.AddEvent(&newEvent)
	if err != nil {
		t.Error(err)
	}
	if calendar.idSeq != 1 {
		t.Error("Id sequence in calendar, dont increment")
	}
	if len(calendar.Events) != 1 {
		t.Error("Event dont append in calendar")
	}
	err = calendar.DeleteEvent(newEvent.Id)
	if err != nil {
		t.Error(err)
	}
	if len(calendar.Events) != 0 {
		t.Error("Event dont delete event in calendar")
	}

}
