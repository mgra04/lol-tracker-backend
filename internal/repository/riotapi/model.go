package riotapi

type AccountDTO struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type ChampionMasteryDTO struct {
	Puuid                    string                `json:"puuid"`
	PointsUntilNextLv        int                   `json:"championPointsUntilNextLevel"`
	PointsSinceLastLv        int                   `json:"championPointsSinceLastLevel"`
	ChestGranted             bool                  `json:"chestGranted"`
	ChampionId               int                   `json:"championId"`
	LastPlayTime             int64                 `json:"lastPlayTime"`
	ChampionLevel            int                   `json:"championLevel"`
	ChampionPoints           int                   `json:"championPoints"`
	MarkRequiredForNextLevel int                   `json:"markRequiredForNextLevel"`
	ChampionSeasonMilestone  int                   `json:"championSeasonMilestone"`
	TokensEarned             int                   `json:"tokenEarned"`
	MilestoneGrades          []string              `json:"milestoneGrades"`
	NextSeasonMilestone      *NextSeasonMilestones `json:"nextSeasonMilestone"`
}

type NextSeasonMilestones struct {
	RequireGradeCounts map[string]int `json:"requireGradeCounts"`
	RewardMarks        int            `json:"rewardMarks"`
	Bonus              bool           `json:"bonus"`
	RewardConfig       *RewardConfig  `json:"rewardConfig"`
}

type RewardConfig struct {
	RewardValue   string `json:"rewardValue"`
	RewardType    string `json:"rewardType"`
	MaximumReward int    `json:"maximumReward"`
}

type LeagueEntryDTO struct {
	LeagueId     string      `json:"leagueId"`
	Puuid        string      `json:"puuid"`
	QueueType    string      `json:"queueType"`
	Tier         string      `json:"tier"`
	Rank         string      `json:"rank"`
	LeaguePoints int         `json:"leaguePoints"`
	Wins         int         `json:"wins"`
	Losses       int         `json:"losses"`
	HotStreak    bool        `json:"hotStreak"`
	Veteran      bool        `json:"veteran"`
	FreshBlood   bool        `json:"freshBlood"`
	Inactive     bool        `json:"inactive"`
	MiniSeries   *MiniSeries `json:"miniSeries"`
}

type MiniSeries struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type MatchDTO struct {
	Metadata *MatchMetadata `json:"metadata"`
	Info     *MatchInfo     `json:"info"`
}

type MatchMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	GameCreation       int64          `json:"gameCreation"`
	GameDuration       int            `json:"gameDuration"`
	GameEndTimestamp   int64          `json:"gameEndTimestamp"`
	GameID             int64          `json:"gameId"`
	GameMode           string         `json:"gameMode"`
	GameName           string         `json:"gameName"`
	GameStartTimestamp int64          `json:"gameStartTimestamp"`
	GameType           string         `json:"gameType"`
	GameVersion        string         `json:"gameVersion"`
	MapID              int            `json:"mapId"`
	Participants       []*Participant `json:"participants"`
	PlatformID         string         `json:"platformId"`
	QueueID            int            `json:"queueId"`
	Teams              []*Team        `json:"teams"`
	TournamentCode     string         `json:"tournamentCode"`
}

type Participant struct {
	Assists         int `json:"assists"`
	BaronKills      int `json:"baronKills"`
	BountyLevel     int `json:"bountyLevel"`
	ChampExperience int `json:"champExperience"`
	ChampLevel      int `json:"champLevel"`
	// Prior to patch 11.4, on Feb 18th, 2021, this field returned invalid championIds.
	// We recommend determining the champion based on the championName field for matches played prior to patch 11.4.
	ChampionID   int    `json:"championId"`
	ChampionName string `json:"championName"`
	// This field is currently only utilized for Kayn's transformations.
	// (Legal values: 0 - None, 1 - Slayer, 2 - Assassin)
	ChampionTransform         int  `json:"championTransform"`
	ConsumablesPurchased      int  `json:"consumablesPurchased"`
	DamageDealtToBuildings    int  `json:"damageDealtToBuildings"`
	DamageDealtToObjectives   int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets      int  `json:"damageDealtToTurrets"`
	DamageSelfMitigated       int  `json:"damageSelfMitigated"`
	Deaths                    int  `json:"deaths"`
	DetectorWardsPlaced       int  `json:"detectorWardsPlaced"`
	DoubleKills               int  `json:"doubleKills"`
	DragonKills               int  `json:"dragonKills"`
	FirstBloodAssist          bool `json:"firstBloodAssist"`
	FirstBloodKill            bool `json:"firstBloodKill"`
	FirstTowerAssist          bool `json:"firstTowerAssist"`
	FirstTowerKill            bool `json:"firstTowerKill"`
	GameEndedInEarlySurrender bool `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender      bool `json:"gameEndedInSurrender"`
	GoldEarned                int  `json:"goldEarned"`
	GoldSpent                 int  `json:"goldSpent"`
	// Both individualPosition and teamPosition are computed by the game server and are
	// different versions of the most likely position played by a player. The individualPosition
	// is the best guess for which position the player actually played in isolation of
	// anything else. The teamPosition is the best guess for which position the player
	// actually played if we add the constraint that each team must have one top player, one
	// jungle, one middle, etc. Generally the recommendation is to use the teamPosition field
	// over the individualPosition field.
	IndividualPosition             string            `json:"individualPosition"`
	InhibitorKills                 int               `json:"inhibitorKills"`
	InhibitorTakedowns             int               `json:"inhibitorTakedowns"`
	InhibitorsLost                 int               `json:"inhibitorsLost"`
	Item0                          int               `json:"item0"`
	Item1                          int               `json:"item1"`
	Item2                          int               `json:"item2"`
	Item3                          int               `json:"item3"`
	Item4                          int               `json:"item4"`
	Item5                          int               `json:"item5"`
	Item6                          int               `json:"item6"`
	ItemsPurchased                 int               `json:"itemsPurchased"`
	KillingSprees                  int               `json:"killingSprees"`
	Kills                          int               `json:"kills"`
	Lane                           string            `json:"lane"`
	LargestCriticalStrike          int               `json:"largestCriticalStrike"`
	LargestKillingSpree            int               `json:"largestKillingSpree"`
	LargestMultiKill               int               `json:"largestMultiKill"`
	LongestTimeSpentLiving         int               `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int               `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int               `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int               `json:"magicDamageTaken"`
	NeutralMinionsKilled           int               `json:"neutralMinionsKilled"`
	NexusKills                     int               `json:"nexusKills"`
	NexusLost                      int               `json:"nexusLost"`
	NexusTakedowns                 int               `json:"nexusTakedowns"`
	ObjectivesStolen               int               `json:"objectivesStolen"`
	ObjectivesStolenAssists        int               `json:"objectivesStolenAssists"`
	ParticipantID                  int               `json:"participantId"`
	PentaKills                     int               `json:"pentaKills"`
	Perks                          *ParticipantPerks `json:"perks"`
	PhysicalDamageDealt            int               `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int               `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int               `json:"physicalDamageTaken"`
	ProfileIcon                    int               `json:"profileIcon"`
	PUUID                          string            `json:"puuid"`
	QuadraKills                    int               `json:"quadraKills"`
	RiotIDGameName                 string            `json:"riotIdGameName"`
	RiotIDName                     string            `json:"riotIdName"`
	RiotIDTagline                  string            `json:"riotIdTagline"`
	Role                           string            `json:"role"`
	SightWardsBoughtInGame         int               `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int               `json:"spell1Casts"`
	Spell2Casts                    int               `json:"spell2Casts"`
	Spell3Casts                    int               `json:"spell3Casts"`
	Spell4Casts                    int               `json:"spell4Casts"`
	Summoner1Casts                 int               `json:"summoner1Casts"`
	Summoner1ID                    int               `json:"summoner1Id"`
	Summoner2Casts                 int               `json:"summoner2Casts"`
	Summoner2ID                    int               `json:"summoner2Id"`
	SummonerID                     string            `json:"summonerId"`
	SummonerLevel                  int               `json:"summonerLevel"`
	SummonerName                   string            `json:"summonerName"`
	TeamEarlySurrendered           bool              `json:"teamEarlySurrendered"`
	TeamID                         int               `json:"teamId"`
	// Both individualPosition and teamPosition are computed by the game server and are
	// different versions of the most likely position played by a player. The individualPosition
	// is the best guess for which position the player actually played in isolation of
	// anything else. The teamPosition is the best guess for which position the player
	// actually played if we add the constraint that each team must have one top player, one
	// jungle, one middle, etc. Generally the recommendation is to use the teamPosition field
	// over the individualPosition field.
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TripleKills                    int    `json:"tripleKills"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlaced"`
	Win                            bool   `json:"win"`
}

type Team struct {
	Bans       []*TeamBan `json:"bans"`
	Objectives Objectives `json:"objectives"`
	TeamID     int        `json:"teamId"`
	Win        bool       `json:"win"`
}

type TeamBan struct {
	// Turn during which the champion was banned.
	PickTurn int `json:"pickTurn"`
	// Banned championId.
	ChampionID int `json:"championId"`
}

type ParticipantPerks struct {
	StatPerks *StatPerks `json:"statPerks"`
	Styles    []Styles   `json:"styles"`
}

type Styles struct {
	Description string       `json:"description"`
	Selections  []Selections `json:"selections"`
	Style       int          `json:"style"`
}

type Selections struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

type StatPerks struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type Objective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

type Objectives struct {
	Baron      Objective `json:"baron"`
	Champion   Objective `json:"champion"`
	Dragon     Objective `json:"dragon"`
	Inhibitor  Objective `json:"inhibitor"`
	RiftHerald Objective `json:"riftHerald"`
	Tower      Objective `json:"tower"`
}

type PlatformDataDTO struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Locales      []string `json:"locales"`
	Maintenances []Status `json:"maintenances"`
	Incidents    []Status `json:"incidents"`
}

type Status struct {
	Id                int       `json:"id"`
	MaintenanceStatus string    `json:"maintenance_status"`
	IncidentSeverity  string    `json:"incident_severity"`
	Titles            []Content `json:"titles"`
	Updates           []*Update `json:"updates"`
	CreatedAt         string    `json:"created_at"`
	ArchiveAt         string    `json:"archive_at"`
	UpdatedAt         string    `json:"updated_at"`
	Platforms         []string  `json:"platforms"`
}

type Content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type Update struct {
	Id               int       `json:"id"`
	Author           string    `json:"author"`
	Publish          bool      `json:"publish"`
	PublishLocations []string  `json:"publish_locations"`
	Translations     []Content `json:"translations"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
}

// MatchTimelineDTO contains timeline frames for a match
type MatchTimelineDTO struct {
	Frames   []*MatchFrame `json:"frames"`
	Interval int           `json:"frameInterval"`
}

// MatchFrame is a single frame in the timeline of a game
type MatchFrame struct {
	Timestamp         int                          `json:"timestamp"`
	ParticipantFrames map[string]*ParticipantFrame `json:"participantFrames"`
	Events            []*MatchEvent                `json:"events"`
}

type ParticipantFrame struct {
	TotalGold           int            `json:"totalGold"`
	TeamScore           int            `json:"teamScore"`
	ParticipantID       int            `json:"participantId"`
	Level               int            `json:"level"`
	CurrentGold         int            `json:"currentGold"`
	MinionsKilled       int            `json:"minionsKilled"`
	DominionScore       int            `json:"dominionScore"`
	Position            *MatchPosition `json:"position"`
	XP                  int            `json:"xp"`
	JungleMinionsKilled int            `json:"jungleMinionsKilled"`
}

// MatchEventType is the type of an event
type MatchEventType string

// All legal value for match event types
const (
	MatchEventTypeChampionKill     MatchEventType = "CHAMPION_KILL"
	MatchEventTypeWardPlaced       MatchEventType = "WARD_PLACED"
	MatchEventTypeWardKill         MatchEventType = "WARD_KILL"
	MatchEventTypeBuildingKill     MatchEventType = "BUILDING_KILL"
	MatchEventTypeEliteMonsterKill MatchEventType = "ELITE_MONSTER_KILL"
	MatchEventTypeItemPurchased    MatchEventType = "ITEM_PURCHASED"
	MatchEventTypeItemSold         MatchEventType = "ITEM_SOLD"
	MatchEventTypeItemDestroyed    MatchEventType = "ITEM_DESTROYED"
	MatchEventTypeItemUndo         MatchEventType = "ITEM_UNDO"
	MatchEventTypeSkillLevelUp     MatchEventType = "SKILL_LEVEL_UP"
	MatchEventTypeAscendedEvent    MatchEventType = "ASCENDED_EVENT"
	MatchEventTypeCapturePoint     MatchEventType = "CAPTURE_POINT"
	MatchEventTypePoroKingSummon   MatchEventType = "PORO_KING_SUMMON"
)

var (
	// MatchEventTypes is a list of all available match events
	MatchEventTypes = []MatchEventType{
		MatchEventTypeChampionKill,
		MatchEventTypeWardPlaced,
		MatchEventTypeWardKill,
		MatchEventTypeBuildingKill,
		MatchEventTypeEliteMonsterKill,
		MatchEventTypeItemPurchased,
		MatchEventTypeItemSold,
		MatchEventTypeItemDestroyed,
		MatchEventTypeItemUndo,
		MatchEventTypeSkillLevelUp,
		MatchEventTypeAscendedEvent,
		MatchEventTypeCapturePoint,
		MatchEventTypePoroKingSummon,
	}
)

// MatchEvent is an event in a match at a certain timestamp
type MatchEvent struct {
	EventType               string          `json:"eventType"`
	TowerType               string          `json:"towerType"`
	TeamID                  int             `json:"teamId"`
	AscendedType            string          `json:"ascendedType"`
	KillerID                int             `json:"killerId"`
	LevelUpType             string          `json:"levelUpType"`
	PointCaptured           string          `json:"pointCaptured"`
	AssistingParticipantIDs []int           `json:"assistingParticipantIds"`
	WardType                string          `json:"wardType"`
	MonsterType             string          `json:"monsterType"`
	Type                    *MatchEventType `json:"type"`
	SkillSlot               int             `json:"skillSlot"`
	VictimID                int             `json:"victimId"`
	Timestamp               int             `json:"timestamp"`
	AfterID                 int             `json:"afterId"`
	MonsterSubType          string          `json:"monsterSubType"`
	LaneType                string          `json:"laneType"`
	ItemID                  int             `json:"itemId"`
	ParticipantID           int             `json:"participantId"`
	BuildingType            string          `json:"buildingType"`
	CreatorID               int             `json:"creatorId"`
	Position                *MatchPosition  `json:"position"`
	BeforeID                int             `json:"beforeId"`
}

// MatchPosition is a position on the map in a game
type MatchPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SummonerDTO struct {
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}