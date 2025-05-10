package country

type Country struct {
	ID     int    `json:"id"`
	NameRu string `json:"country_ru"`
	NameEn string `json:"country_en"`
	Flag   string
}

type CountryList []Country

type CountryWithChurchesCount struct {
	Country
	ChurchesCount int `json:"churches_count"`
}

type CountryWithChurchesCountList []CountryWithChurchesCount

func ToModels(dtos *[]WithChurchesCountDTO) []CountryWithChurchesCount {
	var models []CountryWithChurchesCount
	for _, dto := range *dtos {
		models = append(models, dto.ToModel())
	}
	return models
}

func (c *CountryWithChurchesCount) ToDtoResponse() DtoResponse {
	return DtoResponse{
		ID:            c.ID,
		NameRu:        c.NameRu,
		NameEn:        c.NameEn,
		Flag:          c.Flag,
		ChurchesCount: c.ChurchesCount,
	}
}

func ToDtoResponses(models *[]CountryWithChurchesCount) []DtoResponse {
	var dtoResponses []DtoResponse
	for _, model := range *models {
		dtoResponses = append(dtoResponses, model.ToDtoResponse())
	}
	return dtoResponses
}
