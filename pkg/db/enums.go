package db

type Source uint

const (
	Manabox Source = iota
)

func (s Source) String() string {
	return [...]string{"Manabox"}[s]
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

type BoardCategory uint

const (
	Mainboard BoardCategory = iota
	Sideboard
	Commanders
	Maybeboard
)

func (s BoardCategory) String() string {
	return [...]string{
		"Mainboard",
		"Sideboard",
		"Commanders",
		"Maybeboard",
	}[s]
}
