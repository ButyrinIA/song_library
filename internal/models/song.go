package models

// Song представляет песню
// swagger:model
type Song struct {
	ID          uint   `db:"id" json:"id" `
	Group       string `db:"group" json:"group"`
	Song        string `db:"song" json:"song"`
	Text        string `db:"text" json:"text,omitempty"`
	Link        string `db:"link" json:"link,omitempty"`
	ReleaseDate string `db:"releaseDate" json:"releaseDate,omitempty"`
	APIFetched  bool   `db:"api_fetched" json:"api_fetched"`
}
