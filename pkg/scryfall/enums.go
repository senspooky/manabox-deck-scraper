package scryfall

type Language uint

const (
	English Language = iota
	Spanish
	French
	German
	Italian
	Portuguese
	Japanese
	Korean
	Russian
	SimplifiedChinese
	TraditionalChinese
	Hebrew
	Latin
	AncientGreek
	Arabic
	Sanskrit
	Phyrexian
)

func (l Language) String() string {
	return [...]string{
		"en",
		"es",
		"fr",
		"de",
		"it",
		"pt",
		"ja",
		"ko",
		"ru",
		"zhs",
		"nht",
		"he",
		"la",
		"grc",
		"ar",
		"sa",
		"ph",
	}[l]
}

func (l Language) Friendly() string {
	return [...]string{
		"English",
		"Spanish",
		"French",
		"German",
		"Italian",
		"Portuguese",
		"Japanese",
		"Korean",
		"Russian",
		"Simplified Chinese",
		"Traditional Chinese",
		"Hebrew",
		"Latin",
		"Ancient Greek",
		"Arabic",
		"Sanskirt",
		"Phyrexian",
	}[l]
}

func (l Language) PrintedCode() string {
	return [...]string{
		"en",
		"es",
		"fr",
		"de",
		"it",
		"pt",
		"ja",
		"ko",
		"ru",
		"cs",
		"ct",
		"", // These languages have not been printed on cards with lang. codes
		"",
		"",
		"",
		"",
		"ph",
	}[l]
}

type Layout uint

const (
	Normal Layout = iota
	Split
	Flip
	Transform
	ModalDFC
	Meld
	Leveler
	Class
	Saga
	Adventure
	Mutate
	Prototype
	Battle
	Planar
	Scheme
	Vanguard
	Token
	DoubleFacedToken
	Emblem
	Augment
	Host
	ArtSeries
	ReversibleCard
)

func (l Layout) String() string {
	return [...]string{
		"normal",
		"split",
		"flip",
		"transform",
		"modal_dfc",
		"meld",
		"leveler",
		"class",
		"saga",
		"adventure",
		"mutate",
		"prototype",
		"battle",
		"planar",
		"scheme",
		"vanguard",
		"token",
		"double_faced_token",
		"emblem",
		"augment",
		"host",
		"art_series",
		"reversible_card",
	}[l]
}

func (l Layout) Friendly() string {
	return [...]string{
		"Normal",
		"Split",
		"Flip",
		"Transform",
		"Modal DFC",
		"Meld",
		"Leveler",
		"Class",
		"Saga",
		"Adventure",
		"Mutate",
		"Prototype",
		"Battle",
		"Planar",
		"Scheme",
		"Vanguard",
		"Token",
		"Double-Faced Token",
		"Emblem",
		"Augment",
		"Host",
		"Art Series",
		"Reversible Card",
	}[l]
}

type Color uint

const (
	White Color = iota
	Blue
	Black
	Red
	Green
)

func (c Color) String() string {
	return [...]string{
		"W",
		"U",
		"B",
		"R",
		"G",
	}[c]
}

func (c Color) Friendly() string {
	return [...]string{
		"White",
		"Blue",
		"Black",
		"Red",
		"Green",
	}[c]
}
