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

func (cwd ChurchWithDistance) ToDtoResponse() DtoResponse {
	return DtoResponse{
		Name:       cwd.Church.NameRu,
		Alias:      cwd.Church.Alias,
		Address:    cwd.Church.AddressRu,
		Latitude:   cwd.Church.Coordinates.Latitude,
		Longitude:  cwd.Church.Coordinates.Longitude,
		Confession: cwd.Church.Confession.Name,
		Distance:   cwd.Distance,
	}
}

func ToModels(dtos *[]DtoRepository) []Church {
	var models []Church
	for _, dto := range *dtos {
		models = append(models, dto.ToModel())
	}
	return models
}
