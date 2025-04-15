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

func (c Church) ToDtoResponse() DtoResponse {
	return DtoResponse{
		Name:       c.NameRu,
		Alias:      c.Alias,
		Address:    c.AddressRu,
		Latitude:   c.Coordinates.Latitude,
		Longitude:  c.Coordinates.Longitude,
		Confession: c.Confession.Name,
	}
}

func ToModels(dtos *[]DtoRepository) []Church {
	var models []Church
	for _, dto := range *dtos {
		models = append(models, dto.ToModel())
	}
	return models
}
