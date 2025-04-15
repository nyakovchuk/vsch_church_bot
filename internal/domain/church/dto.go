package church

type DtoRepository struct {
	ID             int     `db:"id"`
	NameEn         string  `db:"name_en"`
	NameRu         string  `db:"name_ru"`
	Alias          string  `db:"alias"`
	CountryRu      string  `db:"country_ru"`
	CountryId      int     `db:"country_id"`
	StateId        int     `db:"state_id"`
	CityId         int     `db:"city_id"`
	AddressRu      string  `db:"address_ru"`
	Latitude       float64 `db:"latitude"`
	Longitude      float64 `db:"longitude"`
	ConfessionId   int     `db:"confession_id"`
	ConfessionName string  `db:"confession_name"`
}

type DtoResponse struct {
	Name       string
	Alias      string
	Address    string
	Latitude   float64
	Longitude  float64
	Confession string
	Distance   float64
}

func (dto DtoRepository) ToModel() Church {
	return Church{
		ID:        dto.ID,
		NameEn:    dto.NameEn,
		NameRu:    dto.NameRu,
		Alias:     dto.Alias,
		AddressRu: dto.AddressRu,
		Location: Location{
			CountryRu: dto.CountryRu,
			CountryId: dto.CountryId,
			StateId:   dto.StateId,
			CityId:    dto.CityId,
		},
		Coordinates: Coordinates{
			Latitude:  dto.Latitude,
			Longitude: dto.Longitude,
		},
		Confession: Confession{
			ID:   dto.ConfessionId,
			Name: dto.ConfessionName,
		},
	}
}
