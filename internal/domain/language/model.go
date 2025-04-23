package language

type Language struct {
	ID   int
	Code string
	Name string
}

type LanguageList []Language

func ToModels(dtos *[]DtoRepository) []Language {
	var models []Language
	for _, dto := range *dtos {
		models = append(models, dto.ToModel())
	}
	return models
}

func (l LanguageList) ByCode(code string) *Language {
	for _, lang := range l {
		if lang.Code == code {
			return &lang
		}
	}
	return nil
}

func (l LanguageList) ByID(id int) *Language {
	for _, lang := range l {
		if lang.ID == id {
			return &lang
		}
	}
	return nil
}
