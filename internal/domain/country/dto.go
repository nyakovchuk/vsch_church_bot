package country

type DtoRepository struct {
	ID     int    `db:"id"`
	NameRu string `db:"country_ru"`
	NameEn string `db:"country_en"`
}

func (dto DtoRepository) ToModel() Country {
	return Country{
		ID:     dto.ID,
		NameRu: dto.NameRu,
		NameEn: dto.NameEn,
	}
}

type WithChurchesCountDTO struct {
	DtoRepository
	ChurchesCount int `db:"churches_count"`
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
