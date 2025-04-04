package user

import "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"

type User struct {
	ID          int
	TgID        int
	Coordinates model.Coordinates
	LangId      *int
	Radius      int
}
