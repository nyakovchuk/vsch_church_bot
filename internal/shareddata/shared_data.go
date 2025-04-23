package shareddata

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/language"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
)

type Data struct {
	Churches  []church.Church
	Platform  platform.Platform
	Languages []language.Language
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
		fmt.Println("error getting churches", err)
	}

	return Data{
		Churches:  churches,
		Platform:  platform,
		Languages: languages,
	}
}
