package country

import (
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
)

type DtoRepository struct {
	ID     int    `db:"id"`
	NameRu string `db:"country_ru"`
	NameEn string `db:"country_en"`
}

type WithChurchesCountDTO struct {
	DtoRepository
	ChurchesCount int `db:"churches_count"`
}

type DtoResponse struct {
	ID            int
	NameRu        string
	NameEn        string
	Flag          string
	ChurchesCount int
}

func (dto DtoRepository) ToModel() Country {
	return Country{
		ID:     dto.ID,
		NameRu: dto.NameRu,
		NameEn: dto.NameEn,
	}
}

func (dto WithChurchesCountDTO) ToModel() CountryWithChurchesCount {
	return CountryWithChurchesCount{
		Country: Country{
			ID:     dto.ID,
			NameRu: dto.NameRu,
			NameEn: dto.NameEn,
		},
		ChurchesCount: dto.ChurchesCount,
	}
}

func (d DtoResponse) FlagWithName(langCode string) string {
	if langCode == i18n.LangCodeRu {
		return fmt.Sprintf("%s %s", d.Flag, d.NameRu)
	}

	return fmt.Sprintf("%s %s", d.Flag, d.NameEn)
}
