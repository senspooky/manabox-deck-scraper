package manabox

// Deck
type Deck struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Format              int    `json:"format"`
	EditDateUTC         int64  `json:"editDateUTC"`
	Colors              string `json:"colors"`
	CoverImageURL       string `json:"coverImageUrl"`
	CoverImageCvID      int    `json:"coverImageCvId"`
	CoverImageFaceIndex int    `json:"coverImageFaceIndex"`
	Cards               []Card `json:"cards"`
}
type Cardmarket struct {
	Value float64 `json:"value"`
	Link  string  `json:"link"`
}
type TcgPlayer struct {
	Value float64 `json:"value"`
	Link  string  `json:"link"`
}
type CardKingdom struct {
	Value float64 `json:"value"`
	Link  string  `json:"link"`
}
type Cardhoarder struct {
	Value float64 `json:"value"`
	Link  string  `json:"link"`
}
type CardPricing struct {
	Cardmarket  Cardmarket  `json:"cardmarket"`
	TcgPlayer   TcgPlayer   `json:"tcgPlayer"`
	CardKingdom CardKingdom `json:"cardKingdom"`
	Cardhoarder Cardhoarder `json:"cardhoarder"`
}
type Faces struct {
	ManaCost interface{} `json:"manaCost"`
}
type Images struct {
	ImageURLSmall  string `json:"imageUrlSmall"`
	ImageURLNormal string `json:"imageUrlNormal"`
}
type Card struct {
	CvID                 int         `json:"cvId"`
	ScryfallID           string      `json:"scryfallId"`
	Name                 string      `json:"name"`
	Quantity             int         `json:"quantity"`
	BoardCategory        int         `json:"boardCategory"`
	Variant              int         `json:"variant"`
	Type                 int         `json:"type"`
	SetID                string      `json:"setId"`
	SetName              string      `json:"setName"`
	SetIconID            string      `json:"setIconId"`
	CardPricing          CardPricing `json:"cardPricing"`
	Rarity               int         `json:"rarity"`
	CollectorNumberGroup int         `json:"collectorNumberGroup"`
	ColorIdentity        string      `json:"colorIdentity"`
	ManaValue            float64     `json:"manaValue"`
	Layout               int         `json:"layout"`
	CardhoarderID        int         `json:"cardhoarderId,omitempty"`
	Faces                []Faces     `json:"faces"`
	Images               []Images    `json:"images"`
}

// Format "enum"
type Format uint

const (
	NoFormat Format = iota
	Standard
	Modern
	Legacy
	Commander
	Brawl
	Pioneer
	Vintage
	Pauper
	DuelCommander
	Frontier
	Oathbreaker
	Historic
	Penny
	Premodern
	OldSchool
	Conquest
	TinyLeaders
	PauperCommander
	Explorer
	HistoricBrawl
	Gladiator
	CanadianHighlander
)

func (f Format) String() string {
	return [...]string{
		"No Format",
		"Standard",
		"Modern",
		"Legacy",
		"Commander",
		"Brawl",
		"Pioneer",
		"Vintage",
		"Pauper",
		"Duel Commander",
		"Frontier",
		"Oathbreaker",
		"Historic",
		"Penny Dreadful",
		"Premodern",
		"Old School",
		"Conquest",
		"Tiny Leaders",
		"Pauper Commander",
		"Explorer",
		"Historic Brawl",
		"Gladiator",
		"Canadian Highlander",
	}[f]
}
