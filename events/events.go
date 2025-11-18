package events

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/ilonaoken/app/valid"
)

type Event struct {
	Id      string
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	if !valid.IsTitleValidate(title) {
		return Event{}, errors.New("не валидная строка")
	}
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("неверный формат даты")
	}
	return Event{
		Id:      uuid.New().String(),
		Title:   title,
		StartAt: t,
	}, nil
}
