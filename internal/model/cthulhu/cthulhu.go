package cthulhu

import "go.mongodb.org/mongo-driver/bson/primitive"

type CharacterInfo struct {
	Sex        string `bson:"sex"`
	Occupation string `bson:"occupation"`
	Residence  string `bson:"residence"`
	Birthplace string `bson:"birthplace"`
}

type Characteristic struct {
	Ages        int `bson:"ages"`
	Str         int `bson:"str"`
	Dex         int `bson:"dex"`
	Int         int `bson:"int"`
	Con         int `bson:"con"`
	App         int `bson:"app"`
	Pow         int `bson:"pow"`
	Siz         int `bson:"siz"`
	Edu         int `bson:"edu"`
	StrOrg      int `bson:"str_org"`
	DexOrg      int `bson:"dex_org"`
	IntOrg      int `bson:"int_org"`
	ConOrg      int `bson:"con_org"`
	AppOrg      int `bson:"app_org"`
	PowOrg      int `bson:"pow_org"`
	SizOrg      int `bson:"siz_org"`
	EduOrg      int `bson:"edu_org"`
	Luck        int `bson:"luck"`
	LuckOrg     int `bson:"luck_org"`
	LuckXp      int `bson:"luck_xp"`
	Sanity      int `bson:"sanity"`
	SanityOrg   int `bson:"sanity_org"`
	HitPoints   int `bson:"hit_points"`
	MagicPoints int `bson:"magic_points"`
	MajorWound  int `bson:"major_wound"`
	TempInsane  int `bson:"temp_insane"`
	IndefInsane int `bson:"indef_insane"`
}

type Skill struct {
	Value int `bson:"value"`
	Xp    int `bson:"xp"`
}

type AdditionalSkill struct {
	Name  string `bson:"name"`
	Label string `bson:"label"`
	Value int    `bson:"value"`
	Xp    int    `bson:"xp"`
}

type Weapon struct {
	Skill       string `bson:"skill"`
	Weapon      string `bson:"weapon"`
	Damage      string `bson:"damage"`
	Range       int    `bson:"range"`
	Attacks     int    `bson:"attacks"`
	Ammo        int    `bson:"ammo"`
	Malfunction int    `bson:"malfunction"`
}

type Spell struct {
	Spell    string `bson:"spell"`
	Cost     string `bson:"cost"`
	CastTime string `bson:"cast_time"`
}

type Encounter struct {
	Entity     string `bson:"entity"`
	SanityLoss int    `bson:"sanity_loss"`
	Total      int    `bson:"total"`
}

type Stuff struct {
	Item    string `bson:"item"`
	Weight  string `bson:"weight"`
	Special string `bson:"special"`
}

type CthulhuCharacter struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	OwnerId       string             `bson:"ownerId"`
	Character     Character          `bson:"character"`
	Money         int                `bson:"money"`
	SpendingLevel int                `bson:"spending_level"`
}

type Character struct {
	CharacterId      string            `bson:"characterId"`
	Name             string            `bson:"name"`
	Notes            string            `bson:"notes"`
	Info             CharacterInfo     `bson:"info"`
	Characteristics  Characteristic    `bson:"characteristics"`
	Skills           map[string]Skill  `bson:"skills"`
	AdditionalSkills []AdditionalSkill `bson:"additional_skills"`
	Weapons          []Weapon          `bson:"weapons"`
	Spells           []Spell           `bson:"spells"`
	Encounters       []Encounter       `bson:"encounters"`
	Stuff            []Stuff           `bson:"stuff"`
}
