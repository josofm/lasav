package lasav

type Card struct {
	Legalities   []Legality        `json:"legalities"`
	Name         string            `json:"name"`
	Multiverseid string            `json:"multiverseid"`
	Supertypes   []string          `json:"supertypes"`
	ForeignNames []ForeignCardName `json:"foreignNames"`
}

type Legality struct {
	Format   string `json:"format"`
	Legality string `json:"legality"`
}

type ForeignCardName struct {
	Name         string `json:"name"`
	Language     string `json:"language"`
	MultiverseId uint   `json:"multiverseid"`
}
