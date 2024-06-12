package keys

import (
	"encoding/json"
	"errors"
	"strings"
)

type Weapon int

func (c *Weapon) MarshalJSON() ([]byte, error) {
	return json.Marshal(weaponNames[*c])
}

func (c *Weapon) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	s = strings.ToLower(s)
	for i, v := range weaponNames {
		if v == s {
			*c = Weapon(i)
			return nil
		}
	}
	return errors.New("unrecognized weapon key")
}

func (c Weapon) String() string {
	return weaponNames[c]
}

var weaponNames = []string{
	"",
	"akuoumaru",
	"alleyhunter",
	"amenomakageuchi",
	"amosbow",
	"apprenticesnotes",
	"aquasimulacra",
	"aquilafavonia",
	"athousandfloatingdreams",
	"balladoftheboundlessblue",
	"balladofthefjords",
	"beaconofthereedsea",
	"beginnersprotector",
	"blackcliffagate",
	"blackclifflongsword",
	"blackcliffpole",
	"blackcliffslasher",
	"blackcliffwarbow",
	"blacktassel",
	"bloodtaintedgreatsword",
	"calamityqueller",
	"cashflowsupervision",
	"cinnabarspindle",
	"cloudforged",
	"compoundbow",
	"coolsteel",
	"cranesechoingcall",
	"crescentpike",
	"crimsonmoonssemblance",
	"darkironsword",
	"deathmatch",
	"debateclub",
	"dialoguesofthedesertsages",
	"dodocotales",
	"dragonsbane",
	"dragonspinespear",
	"dullblade",
	"elegyfortheend",
	"emeraldorb",
	"endoftheline",
	"engulfinglightning",
	"everlastingmoonglow",
	"eyeofperception",
	"fadingtwilight",
	"favoniuscodex",
	"favoniusgreatsword",
	"favoniuslance",
	"favoniussword",
	"favoniuswarbow",
	"ferrousshadow",
	"festeringdesire",
	"filletblade",
	"finaleofthedeep",
	"fleuvecendreferryman",
	"flowingpurity",
	"forestregalia",
	"freedomsworn",
	"frostbearer",
	"fruitoffulfillment",
	"hakushinring",
	"halberd",
	"hamayumi",
	"harangeppakufutsu",
	"harbingerofdawn",
	"huntersbow",
	"hunterspath",
	"ibispiercer",
	"ironpoint",
	"ironsting",
	"jadefallssplendor",
	"kagotsurubeisshin",
	"kagurasverity",
	"katsuragikirinagamasa",
	"keyofkhajnisut",
	"kingssquire",
	"kitaincrossspear",
	"lightoffoliarincision",
	"lionsroar",
	"lithicblade",
	"lithicspear",
	"lostprayertothesacredwinds",
	"luxurioussealord",
	"magicguide",
	"mailedflower",
	"makhairaaquamarine",
	"mappamare",
	"memoryofdust",
	"messenger",
	"missivewindspear",
	"mistsplitterreforged",
	"mitternachtswaltz",
	"moonpiercer",
	"mouunsmoon",
	"oathsworneye",
	"oldmercspal",
	"otherworldlystory",
	"pocketgrimoire",
	"polarstar",
	"portablepowersaw",
	"predator",
	"primordialjadecutter",
	"primordialjadewingedspear",
	"prospectorsdrill",
	"prototypeamber",
	"prototypearchaic",
	"prototypecrescent",
	"prototyperancour",
	"prototypestarglitter",
	"rainslasher",
	"rangegauge",
	"ravenbow",
	"recurvebow",
	"redhornstonethresher",
	"rightfulreward",
	"royalbow",
	"royalgreatsword",
	"royalgrimoire",
	"royallongsword",
	"royalspear",
	"rust",
	"sacrificialbow",
	"sacrificialfragments",
	"sacrificialgreatsword",
	"sacrificialjade",
	"sacrificialsword",
	"sapwoodblade",
	"scionoftheblazingsun",
	"seasonedhuntersbow",
	"serpentspine",
	"sharpshootersoath",
	"silversword",
	"skyridergreatsword",
	"skyridersword",
	"skywardatlas",
	"skywardblade",
	"skywardharp",
	"skywardpride",
	"skywardspine",
	"slingshot",
	"snowtombedstarsilver",
	"solarpearl",
	"songofbrokenpines",
	"songofstillness",
	"splendoroftranquilwaters",
	"staffofhoma",
	"staffofthescarletsands",
	"summitshaper",
	"swordofdescension",
	"swordofnarzissenkreuz",
	"talkingstick",
	"thealleyflash",
	"thebell",
	"theblacksword",
	"thecatch",
	"thedockhandsassistant",
	"thefirstgreatmagic",
	"theflute",
	"thestringless",
	"theunforged",
	"theviridescenthunt",
	"thewidsith",
	"thrillingtalesofdragonslayers",
	"thunderingpulse",
	"tidalshadow",
	"tomeoftheeternalflow",
	"toukaboushigure",
	"travelershandysword",
	"tulaytullahsremembrance",
	"twinnephrite",
	"ultimateoverlordsmegamagicsword",
	"urakumisugiri",
	"verdict",
	"vortexvanquisher",
	"wanderingevenstar",
	"wastergreatsword",
	"wavebreakersfin",
	"whiteblind",
	"whiteirongreatsword",
	"whitetassel",
	"windblumeode",
	"wineandsong",
	"wolffang",
	"wolfsgravestone",
	"xiphosmoonlight",
}

const (
	NoWeapon Weapon = iota
	Akuoumaru
	AlleyHunter
	AmenomaKageuchi
	AmosBow
	ApprenticesNotes
	AquaSimulacra
	AquilaFavonia
	AThousandFloatingDreams
	BalladOfTheBoundlessBlue
	BalladOfTheFjords
	BeaconOfTheReedSea
	BeginnersProtector
	BlackcliffAgate
	BlackcliffLongsword
	BlackcliffPole
	BlackcliffSlasher
	BlackcliffWarbow
	BlackTassel
	BloodtaintedGreatsword
	CalamityQueller
	CashflowSupervision
	CinnabarSpindle
	Cloudforged
	CompoundBow
	CoolSteel
	CranesEchoingCall
	CrescentPike
	CrimsonMoonsSemblance
	DarkIronSword
	Deathmatch
	DebateClub
	DialoguesOfTheDesertSages
	DodocoTales
	DragonsBane
	DragonspineSpear
	DullBlade
	ElegyForTheEnd
	EmeraldOrb
	EndOfTheLine
	EngulfingLightning
	EverlastingMoonglow
	EyeOfPerception
	FadingTwilight
	FavoniusCodex
	FavoniusGreatsword
	FavoniusLance
	FavoniusSword
	FavoniusWarbow
	FerrousShadow
	FesteringDesire
	FilletBlade
	FinaleOfTheDeep
	FleuveCendreFerryman
	FlowingPurity
	ForestRegalia
	FreedomSworn
	Frostbearer
	FruitOfFulfillment
	HakushinRing
	Halberd
	Hamayumi
	HaranGeppakuFutsu
	HarbingerOfDawn
	HuntersBow
	HuntersPath
	IbisPiercer
	IronPoint
	IronSting
	JadefallsSplendor
	KagotsurubeIsshin
	KagurasVerity
	KatsuragikiriNagamasa
	KeyOfKhajNisut
	KingsSquire
	KitainCrossSpear
	LightOfFoliarIncision
	LionsRoar
	LithicBlade
	LithicSpear
	LostPrayerToTheSacredWinds
	LuxuriousSeaLord
	MagicGuide
	MailedFlower
	MakhairaAquamarine
	MappaMare
	MemoryOfDust
	Messenger
	MissiveWindspear
	MistsplitterReforged
	MitternachtsWaltz
	Moonpiercer
	MouunsMoon
	OathswornEye
	OldMercsPal
	OtherworldlyStory
	PocketGrimoire
	PolarStar
	PortablePowerSaw
	Predator
	PrimordialJadeCutter
	PrimordialJadeWingedSpear
	ProspectorsDrill
	PrototypeAmber
	PrototypeArchaic
	PrototypeCrescent
	PrototypeRancour
	PrototypeStarglitter
	Rainslasher
	RangeGauge
	RavenBow
	RecurveBow
	RedhornStonethresher
	RightfulReward
	RoyalBow
	RoyalGreatsword
	RoyalGrimoire
	RoyalLongsword
	RoyalSpear
	Rust
	SacrificialBow
	SacrificialFragments
	SacrificialGreatsword
	SacrificialJade
	SacrificialSword
	SapwoodBlade
	ScionOfTheBlazingSun
	SeasonedHuntersBow
	SerpentSpine
	SharpshootersOath
	SilverSword
	SkyriderGreatsword
	SkyriderSword
	SkywardAtlas
	SkywardBlade
	SkywardHarp
	SkywardPride
	SkywardSpine
	Slingshot
	SnowTombedStarsilver
	SolarPearl
	SongOfBrokenPines
	SongOfStillness
	SplendorOfTranquilWaters
	StaffOfHoma
	StaffOfTheScarletSands
	SummitShaper
	SwordOfDescension
	SwordOfNarzissenkreuz
	TalkingStick
	TheAlleyFlash
	TheBell
	TheBlackSword
	TheCatch
	TheDockhandsAssistant
	TheFirstGreatMagic
	TheFlute
	TheStringless
	TheUnforged
	TheViridescentHunt
	TheWidsith
	ThrillingTalesOfDragonSlayers
	ThunderingPulse
	TidalShadow
	TomeOfTheEternalFlow
	ToukabouShigure
	TravelersHandySword
	TulaytullahsRemembrance
	TwinNephrite
	UltimateOverlordsMegaMagicSword
	UrakuMisugiri
	Verdict
	VortexVanquisher
	WanderingEvenstar
	WasterGreatsword
	WavebreakersFin
	Whiteblind
	WhiteIronGreatsword
	WhiteTassel
	WindblumeOde
	WineAndSong
	WolfFang
	WolfsGravestone
	XiphosMoonlight
)
