package calendar

import (
	"github.com/egor1344/otus_home/calendar/models/event"
	"github.com/pkg/errors"
	"sync"
)

// Структура события
type Event event.Event

// Структура календаря
type Calendar struct {
	idSeq  int64 // Счетчик id(костыль пока бд не подключим)
	Events map[int64]*Event
}

// Инициализация каледаря
func (c *Calendar) Init() {
	c.Events = make(map[int64]*Event)
}

// Добавление события в календаре
func (c *Calendar) AddEvent(event *Event) error {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := c.Events[event.Id]
	if ok {
		return errors.New("Event with this id has already been created")
	}
	c.Events[event.Id] = event
	if event.Id > c.idSeq {
		c.idSeq = event.Id
	}
	return nil
}

// Обновление события в календаре
func (c *Calendar) UpdateEvent(eventId int64, newEvent *Event) error {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	if eventId != 0 {
		_, ok := c.Events[eventId]
		if !ok {
			return errors.New("Event with id not created")
		}
		c.Events[eventId] = newEvent
		return nil
	}
	_, ok := c.Events[newEvent.Id]
	if !ok {
		return errors.New("Event with id not created")
	}
	c.Events[newEvent.Id] = newEvent
	return nil
}

// Удаление события в календаря
func (c *Calendar) DeleteEvent(eventId int64) error {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	delete(c.Events, eventId)
	return nil
}
