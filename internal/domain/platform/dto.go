package platform

type DtoRepository struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (dto DtoRepository) ToModel() Platform {
	return Platform{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
