package runequest

type Info struct {
	Sex        string `bson:"sex"`
	Species    string `bson:"species"`
	Clan       string `bson:"clan"`
	Age        string `bson:"age"`
	Culture    string `bson:"culture"`
	Religion   string `bson:"religion"`
	Parent     string `bson:"parent"`
	Occupation string `bson:"occupation"`
	Reputation int    `bson:"reputation"`
}

type Characteristic struct {
	Str          int `bson:"str"`
	Con          int `bson:"con"`
	Siz          int `bson:"siz"`
	Int          int `bson:"int"`
	Dex          int `bson:"dex"`
	Pow          int `bson:"pow"`
	Cha          int `bson:"cha"`
	StrMax       int `bson:"str_max"`
	ConMax       int `bson:"con_max"`
	SizMax       int `bson:"siz_max"`
	IntMax       int `bson:"int_max"`
	DexMax       int `bson:"dex_max"`
	PowMax       int `bson:"pow_max"`
	ChaMax       int `bson:"cha_max"`
	StrOrg       int `bson:"str_org"`
	ConOrg       int `bson:"con_org"`
	SizOrg       int `bson:"siz_org"`
	IntOrg       int `bson:"int_org"`
	DexOrg       int `bson:"dex_org"`
	PowOrg       int `bson:"pow_org"`
	ChaOrg       int `bson:"cha_org"`
	MaxPowForGain int `bson:"maxPowForGain"`
	PowXpRolls    int `bson:"powXpRolls"`
	// Other fields...
}

type Skill struct {
	Skill string `bson:"skill"`
	Value string `bson:"value"`
	Xp    int    `bson:"xp"`
}

// Define additional structs for Spirits, WeaponSkills, Weapons, Spells, Armor, Stuff, HitPoints...

type RuneQuestCharacter struct {
	CharacterId string         `bson:"characterId"`
	OwnerId     string         `bson:"ownerId"`
	Character   Character      `bson:"character"`
	// Additional fields for Skills, Spirits, Weapons, etc.
}

// Define a struct for the embedded Character document
type Character struct {
	Name           string            `bson:"name"`
	Notes          string            `bson:"notes"`
	Xp             int               `bson:"xp"`
	Info           Info              `bson:"info"`
	Characteristics Characteristic   `bson:"characteristics"`
	Skills         []Skill           `bson:"skills"`
	// Additional fields for Spirits, Weapons, Spells, etc.
}
