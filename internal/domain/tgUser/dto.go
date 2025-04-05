package tgUser

type DtoRepository struct {
	TgID         int64  `db:"tg_id"`
	Username     string `db:"username"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	LanguageCode string `db:"language_code"`
	IsBot        bool   `db:"is_bot"`
	IsPremium    bool   `db:"is_premium"`
}

func ModelToDto(tu TgUser) DtoRepository {
	return DtoRepository{
		TgID:         tu.TgID,
		Username:     tu.Username,
		FirstName:    tu.FirstName,
		LastName:     tu.LastName,
		LanguageCode: tu.LanguageCode,
		IsBot:        tu.IsBot,
		IsPremium:    tu.IsPremium,
	}
}
