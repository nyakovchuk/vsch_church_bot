package language

type DtoRepository struct {
	ID   int    `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}

func (dto DtoRepository) ToModel() Language {
	return Language{
		ID:   dto.ID,
		Code: dto.Code,
		Name: dto.Name,
	}
}
