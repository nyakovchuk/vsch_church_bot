package tguser

type TgUser struct {
	ID           int64
	TgID         int64
	Username     string
	FirstName    string
	LastName     string
	LanguageCode string
	IsBot        bool
	IsPremium    bool
}
