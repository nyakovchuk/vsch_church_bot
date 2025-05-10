package shareddata

import (
	"context"
	"errors"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/country"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/language"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
)

type Data struct {
	Churches  []church.Church
	Platform  platform.Platform
	Languages []language.Language
	Countries []country.CountryWithChurchesCount
}

func New(ctx context.Context, cfg *config.Config, services *service.Service) Data {
	churches, err := services.Church.GetAll(ctx)
	if err != nil {
		fmt.Println("error getting churches", err)
	}

	platform, err := services.Platform.GetByName(ctx, cfg.Platform)
	if err != nil {
		fmt.Println("error getting platform", err)
	}

	languages, err := services.Language.GetAll(ctx)
	if err != nil {
		fmt.Println("error getting languages", err)
	}

	countries, err := services.Country.FetchCountryChurchesStats(ctx)
	if err != nil {
		if errors.Is(err, apperrors.ErrGetFlags) {
			fmt.Println("error getting country flags", err)
		} else {
			fmt.Println("error getting countries", err)
		}
	}

	return Data{
		Churches:  churches,
		Platform:  platform,
		Languages: languages,
		Countries: countries,
	}
}
