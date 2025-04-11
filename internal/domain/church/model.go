package church

type Church struct {
	ID          int
	NameEn      string
	NameRu      string
	Alias       string
	AddressRu   string
	Location    Location
	Coordinates Coordinates
	Confession  Confession
}

type Location struct {
	CountryRu string
	CountryId int
	StateId   int
	CityId    int
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Confession struct {
	ID   int
	Name string
}

func (c Church) ToTelegramDto() DtoTelegram {
	return DtoTelegram{
		Name:       c.NameRu,
		Alias:      c.Alias,
		Address:    c.AddressRu,
		Latitude:   c.Coordinates.Latitude,
		Longitude:  c.Coordinates.Longitude,
		Confession: c.Confession.Name,
	}
}

func ToModel(dto DtoRepository) Church {
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

func ToModels(dtos *[]DtoRepository) []Church {
	var models []Church
	for _, dto := range *dtos {
		models = append(models, ToModel(dto))
	}
	return models
}
