package church

type DtoRepository struct {
	ID             int     `db:"id"`
	NameEn         string  `db:"name_en"`
	NameRu         string  `db:"name_ru"`
	Alias          string  `db:"alias"`
	CountryRu      string  `db:"country_ru"`
	CountryId      int     `db:"country_id"`
	StateId        int     `db:"state_id"`
	CityId         int     `db:"city_id"`
	AddressRu      string  `db:"address_ru"`
	Latitude       float64 `db:"latitude"`
	Longitude      float64 `db:"longitude"`
	ConfessionId   int     `db:"confession_id"`
	ConfessionName string  `db:"confession_name"`
}
