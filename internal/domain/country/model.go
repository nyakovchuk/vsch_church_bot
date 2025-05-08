package country

type Country struct {
	ID     int    `json:"id"`
	NameRu string `json:"country_ru"`
	NameEn string `json:"Country_en"`
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
