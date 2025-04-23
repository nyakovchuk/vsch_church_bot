package language

type Language struct {
	ID   int
	Code string
	Name string
}

func ToModels(dtos *[]DtoRepository) []Language {
	var models []Language
	for _, dto := range *dtos {
		models = append(models, dto.ToModel())
	}
	return models
}
