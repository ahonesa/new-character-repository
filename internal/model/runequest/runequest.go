package runequest

import "go.mongodb.org/mongo-driver/bson/primitive"

type RuneQuestCharacter struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CharacterId string             `bson:"characterId"`
	OwnerId     string             `bson:"ownerId"`
	Character   CharacterDetails   `bson:"character"`
}

type CharacterDetails struct {
	Name            string          `bson:"name"`
	Notes           string          `bson:"notes"`
	XP              int             `bson:"xp"`
	Info            CharacterInfo   `bson:"info"`
	Characteristics Characteristics `bson:"characteristics"`
	Skills          []Skill         `bson:"skills"`
	Spirits         []Spirit        `bson:"spirits"`
	WeaponSkills    []WeaponSkill   `bson:"weaponskills"`
	Weapons         []Weapon        `bson:"weapons"`
	FreeInt         string          `bson:"freeint"`
	Spells          []Spell         `bson:"spells"`
	Armor           []ArmorPiece    `bson:"armor"`
	Stuff           []Stuff         `bson:"stuff"`
	HitPoints       HitPoints       `bson:"hitpoints"`
	HidesOfLand     int             `bson:"hidesOfLand"`
	FlocksOfHerd    int             `bson:"flocksOfHerd"`
	Money           int             `bson:"money"`
}

type CharacterInfo struct {
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

type Characteristics struct {
	Str           int    `bson:"str"`
	Con           int    `bson:"con"`
	Siz           int    `bson:"siz"`
	Int           int    `bson:"int"`
	Dex           int    `bson:"dex"`
	Pow           int    `bson:"pow"`
	Cha           int    `bson:"cha"`
	StrMax        int    `bson:"str_max"`
	ConMax        int    `bson:"con_max"`
	SizMax        int    `bson:"siz_max"`
	IntMax        int    `bson:"int_max"`
	DexMax        int    `bson:"dex_max"`
	PowMax        int    `bson:"pow_max"`
	ChaMax        int    `bson:"cha_max"`
	StrOrg        int    `bson:"str_org"`
	ConOrg        int    `bson:"con_org"`
	SizOrg        int    `bson:"siz_org"`
	IntOrg        int    `bson:"int_org"`
	DexOrg        int    `bson:"dex_org"`
	PowOrg        int    `bson:"pow_org"`
	ChaOrg        int    `bson:"cha_org"`
	MaxPowForGain int    `bson:"maxPowForGain"`
	PowXpRolls    int    `bson:"powXpRolls"`
	Rp1Current    int    `bson:"rp1Current"`
	Rp1Name       string `bson:"rp1Name"`
	Rp1Total      int    `bson:"rp1Total"`
	Rp2Current    int    `bson:"rp2Current"`
	Rp2Name       string `bson:"rp2Name"`
	Rp2Total      int    `bson:"rp2Total"`
	Rp3Current    int    `bson:"rp3Current"`
	Rp3Name       string `bson:"rp3Name"`
	Rp3Total      int    `bson:"rp3Total"`
	Rp4Current    int    `bson:"rp4Current"`
	Rp4Name       string `bson:"rp4Name"`
	Rp4Total      int    `bson:"rp4Total"`
	HeroPoints    int    `bson:"heroPoints"`
}

type Skill struct {
	Skill string `bson:"skill"`
	Value string `bson:"value"`
	XP    int    `bson:"xp"`
}

type Spirit struct {
	Name    string `bson:"name"`
	Species string `bson:"species"`
	Stats   string `bson:"stats"`
	Spells  string `bson:"spells"`
	Notes   string `bson:"notes"`
}

type WeaponSkill struct {
	Skill string `bson:"skill"`
	Value string `bson:"value"`
	XP    int    `bson:"xp"`
}

type Weapon struct {
	Skill      string `bson:"skill"`
	Weapon     string `bson:"weapon"`
	SR         string `bson:"sr"`
	Damage     string `bson:"damage"`
	Armor      string `bson:"armor"`
	WeaponType string `bson:"weaponType"`
}

type Spell struct {
	SpellType string `bson:"spelltype"`
	Spell     string `bson:"spell"`
	Rank      string `bson:"rank"`
	Value     string `bson:"value"`
}

type ArmorPiece struct {
	ArmorType string `bson:"armorType"`
	Head      string `bson:"head"`
	Chest     string `bson:"chest"`
	Abdomen   string `bson:"abdomen"`
	RH        string `bson:"rh"`
	LH        string `bson:"lh"`
	RL        string `bson:"rl"`
	LL        string `bson:"ll"`
}

type Stuff struct {
	Item    string `bson:"item"`
	Weight  string `bson:"weight"`
	Special string `bson:"special"`
}

type HitPoints struct {
	Base    int `bson:"base"`
	Head    int `bson:"head"`
	Chest   int `bson:"chest"`
	Abdomen int `bson:"abdomen"`
	RArm    int `bson:"rarm"`
	LArm    int `bson:"larm"`
	RLeg    int `bson:"rleg"`
	LLeg    int `bson:"lleg"`
}
