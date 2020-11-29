package main

type Pokemon struct {
	Id                     int
	Height                 int
	Weight                 int
	Name                   string
	Order                  int
	BaseExperience         int    `json:"base_experience"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Abilities              []Abilities
	Forms                  []NameUrl
	GameIndices            []GameIndices `json:"game_indices"`
	Types                  []Types
	Stats                  []Stats
	Sprites                Sprites
	Species                NameUrl
	Moves                  []Moves
}

type VersionGroupDetails struct {
	LevelLearnedAt  int     `json:"level_learned_at"`
	MoveLearnMethod NameUrl `json:"move_learn_method"`
	VersionGroup    NameUrl `json:"version_group"`
}

type Moves struct {
	Move                NameUrl
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type NameUrl struct {
	Name string
	Url  string
}

type GenerationImage struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
	BackShiny        string `json:"back_shiny"`
	BackGray         string `json:"back_gray"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
	FrontGray        string `json:"front_gray"`
	Animated         *GenerationImage
}

type GenerationI struct {
	RedBlue GenerationImage `json:"red-blue"`
	Yellow  GenerationImage
}
type GenerationII struct {
	Crystal GenerationImage
	Gold    GenerationImage
	Silver  GenerationImage
}
type GenerationIII struct {
	Emerald          GenerationImage
	FireredLeafgreen GenerationImage `json:"firered-leafgreen"`
	RubySapphire     GenerationImage `json:"ruby-sapphire"`
}
type GenerationIV struct {
	Platinum            GenerationImage
	DiamondPearl        GenerationImage `json:"diamond-pearl"`
	HeartgoldSoulsilver GenerationImage `json:"heartgold-soulsilver"`
}
type GenerationV struct {
	BlackWhite GenerationImage `json:"black-white"`
}
type GenerationVI struct {
	OmegarubyAlphasapphire GenerationImage `json:"omegaruby-alphasapphire"`
	XY                     GenerationImage `json:"x-y"`
}
type GenerationVII struct {
	Icons             GenerationImage
	UltraSunUltraMoon GenerationImage `json:"ultra-sun-ultra-moon"`
}
type GenerationVIII struct {
	Icons GenerationImage
}

type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationII   GenerationII   `json:"generation-ii"`
	GenerationIII  GenerationIII  `json:"generation-iii"`
	GenerationIV   GenerationIV   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVI   GenerationVI   `json:"generation-vi"`
	GenerationVII  GenerationVII  `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  string `json:"front_female"`
}

type OfficialArtWork struct {
	FrontDefault string `json:"front_default"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	OfficialArtWork OfficialArtWork `json:"official-artwork"`
}

type Sprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
	Versions         Versions
}

type Stats struct {
	BaseStat int `json:"base_stat"`
	Effort   int
	Stat     NameUrl
}

type Types struct {
	Slot int
	Type NameUrl
}

type Abilities struct {
	Slot     int
	IsHidden bool `json:"is_hidden"`
	Ability  NameUrl
}

type GameIndices struct {
	GameIndex int `json:"game_index"`
	Version   NameUrl
}
