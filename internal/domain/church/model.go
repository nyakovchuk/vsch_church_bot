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

type ChurchWithDistance struct {
	Church   *Church
	Distance float64
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
	ID     int
	NameRu string
	NameEn string
}

func (c Church) ToDtoResponse() DtoResponse {
	return DtoResponse{
		NameRU:       c.NameRu,
		NameEN:       c.NameEn,
		Alias:        c.Alias,
		Latitude:     c.Coordinates.Latitude,
		Longitude:    c.Coordinates.Longitude,
		ConfessionRu: c.Confession.NameRu,
		ConfessionEn: c.Confession.NameEn,
	}
}

func (cwd ChurchWithDistance) ToDtoResponse() DtoResponse {
	return DtoResponse{
		NameRU:       cwd.Church.NameRu,
		NameEN:       cwd.Church.NameEn,
		Alias:        cwd.Church.Alias,
		Latitude:     cwd.Church.Coordinates.Latitude,
		Longitude:    cwd.Church.Coordinates.Longitude,
		ConfessionRu: cwd.Church.Confession.NameRu,
		ConfessionEn: cwd.Church.Confession.NameEn,
		Distance:     cwd.Distance,
	}
}

func ToModels(dtos *[]DtoRepository) []Church {
	var models []Church
	for _, dto := range *dtos {
		models = append(models, dto.ToModel())
	}
	return models
}
