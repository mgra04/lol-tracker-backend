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
	EndOfGameResult    string         `json:"endOfGameResult"`
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
	// yellow crossed swords ping
	AllInPings int `json:"allInPings"`
	// green flag ping
	AssistMePings   int `json:"assistMePings"`
	Assists         int `json:"assists"`
	BaronKills      int `json:"baronKills"`
	BountyLevel     int `json:"bountyLevel"`
	ChampExperience int `json:"champExperience"`
	ChampLevel      int `json:"champLevel"`
	// Prior to patch 11.4, on Feb 18th, 2021, this field returned invalid championIds.
	// We recommend determining the champion based on the championName field for matches played prior to patch 11.4.
	ChampionID   int    `json:"championId"`
	ChampionName string `json:"championName"`
	// blue generic ping
	CommandPings int `json:"commandPings"`
	// This field is currently only utilized for Kayn's transformations.
	// (Legal values: 0 - None, 1 - Slayer, 2 - Assassin)
	ChampionTransform       int         `json:"championTransform"`
	ConsumablesPurchased    int         `json:"consumablesPurchased"`
	Challenges              *Challenges `json:"challenges"`
	DamageDealtToBuildings  int         `json:"damageDealtToBuildings"`
	DamageDealtToObjectives int         `json:"damageDealtToObjectives"`
	DamageDealtToTurrets    int         `json:"damageDealtToTurrets"`
	DamageSelfMitigated     int         `json:"damageSelfMitigated"`
	Deaths                  int         `json:"deaths"`
	DetectorWardsPlaced     int         `json:"detectorWardsPlaced"`
	DoubleKills             int         `json:"doubleKills"`
	DragonKills             int         `json:"dragonKills"`
	EligibleForProgression  bool        `json:"eligibleForProgression"`
	// 	Yellow questionmark ping
	EnemyMissingPings int `json:"enemyMissingPings"`
	// Red eyeball ping
	EnemyVisionPings          int  `json:"enemyVisionPings"`
	FirstBloodAssist          bool `json:"firstBloodAssist"`
	FirstBloodKill            bool `json:"firstBloodKill"`
	FirstTowerAssist          bool `json:"firstTowerAssist"`
	FirstTowerKill            bool `json:"firstTowerKill"`
	GameEndedInEarlySurrender bool `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender      bool `json:"gameEndedInSurrender"`
	HoldPings                 int  `json:"holdPings"`
	// Tellow circle with horizontal line ping
	GetBackPings int `json:"getBackPings"`
	GoldEarned   int `json:"goldEarned"`
	GoldSpent    int `json:"goldSpent"`
	// Both individualPosition and teamPosition are computed by the game server and are
	// different versions of the most likely position played by a player. The individualPosition
	// is the best guess for which position the player actually played in isolation of
	// anything else. The teamPosition is the best guess for which position the player
	// actually played if we add the constraint that each team must have one top player, one
	// jungle, one middle, etc. Generally the recommendation is to use the teamPosition field
	// over the individualPosition field.
	IndividualPosition          string `json:"individualPosition"`
	InhibitorKills              int    `json:"inhibitorKills"`
	InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
	InhibitorsLost              int    `json:"inhibitorsLost"`
	Item0                       int    `json:"item0"`
	Item1                       int    `json:"item1"`
	Item2                       int    `json:"item2"`
	Item3                       int    `json:"item3"`
	Item4                       int    `json:"item4"`
	Item5                       int    `json:"item5"`
	Item6                       int    `json:"item6"`
	ItemsPurchased              int    `json:"itemsPurchased"`
	KillingSprees               int    `json:"killingSprees"`
	Kills                       int    `json:"kills"`
	Lane                        string `json:"lane"`
	LargestCriticalStrike       int    `json:"largestCriticalStrike"`
	LargestKillingSpree         int    `json:"largestKillingSpree"`
	LargestMultiKill            int    `json:"largestMultiKill"`
	LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
	MagicDamageDealt            int    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken            int    `json:"magicDamageTaken"`
	NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
	// green ward ping
	NeedVisionPings         int `json:"needVisionPings"`
	NexusKills              int `json:"nexusKills"`
	NexusTakedowns          int `json:"nexusTakedowns"`
	NexusLost               int `json:"nexusLost"`
	ObjectivesStolen        int `json:"objectivesStolen"`
	ObjectivesStolenAssists int `json:"objectivesStolenAssists"`
	// blue arrow pointing at ground ping
	OnMyWayPings                   int               `json:"onMyWayPings"`
	ParticipantID                  int               `json:"participantId"`
	PlayerScore0                   int               `json:"playerScore0"`
	PlayerScore1                   int               `json:"playerScore1"`
	PlayerScore2                   int               `json:"playerScore2"`
	PlayerScore3                   int               `json:"playerScore3"`
	PlayerScore4                   int               `json:"playerScore4"`
	PlayerScore5                   int               `json:"playerScore5"`
	PlayerScore6                   int               `json:"playerScore6"`
	PlayerScore7                   int               `json:"playerScore7"`
	PlayerScore8                   int               `json:"playerScore8"`
	PlayerScore9                   int               `json:"playerScore9"`
	PlayerScore10                  int               `json:"playerScore10"`
	PlayerScore11                  int               `json:"playerScore11"`
	PentaKills                     int               `json:"pentaKills"`
	Perks                          *ParticipantPerks `json:"perks"`
	PhysicalDamageDealt            int               `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int               `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int               `json:"physicalDamageTaken"`
	// Arena features
	Placement       int `json:"placement"`
	PlayerAugment1  int `json:"playerAugment1"`
	PlayerAugment2  int `json:"playerAugment2"`
	PlayerAugment3  int `json:"playerAugment3"`
	PlayerAugment4  int `json:"playerAugment4"`
	PlayerSubteamId int `json:"playerSubteamId"`
	// green minion ping
	PushPings              int    `json:"pushPings"`
	ProfileIcon            int    `json:"profileIcon"`
	PUUID                  string `json:"puuid"`
	QuadraKills            int    `json:"quadraKills"`
	RiotIDGameName         string `json:"riotIdGameName"`
	RiotIDName             string `json:"riotIdName"`
	RiotIDTagline          string `json:"riotIdTagline"`
	Role                   string `json:"role"`
	SightWardsBoughtInGame int    `json:"sightWardsBoughtInGame"`
	Spell1Casts            int    `json:"spell1Casts"`
	Spell2Casts            int    `json:"spell2Casts"`
	Spell3Casts            int    `json:"spell3Casts"`
	Spell4Casts            int    `json:"spell4Casts"`
	// Arena features
	SubteamPlacement     int    `json:"subteamPlacement"`
	Summoner1Casts       int    `json:"summoner1Casts"`
	Summoner1ID          int    `json:"summoner1Id"`
	Summoner2Casts       int    `json:"summoner2Casts"`
	Summoner2ID          int    `json:"summoner2Id"`
	SummonerID           string `json:"summonerId"`
	SummonerLevel        int    `json:"summonerLevel"`
	SummonerName         string `json:"summonerName"`
	TeamEarlySurrendered bool   `json:"teamEarlySurrendered"`
	TeamID               int    `json:"teamId"`
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
	TotalAllyJungleMinionsKilled   int    `json:"totalAllyJungleMinionsKilled"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalEnemyJungleMinionsKilled  int    `json:"totalEnemyJungleMinionsKilled"`
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

type Challenges struct {
	AssistStrekCount                          int     `json:"12AssistStreakCount"`
	BaronBuffGoldAdvantageOverThreshold       int     `json:"baronBuffGoldAdvantageOverThreshold"`
	ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
	EarliestBaron                             int     `json:"earliestBaron"`
	EarliestBaronTakedown                     int     `json:"earliestDragonTakedown"`
	EarliestElderDragon                       int     `json:"earliestElderDragon"`
	EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
	FasterSupportQuestCompletion              int     `json:"fasterSupportQuestCompletion"`
	FastestLegendary                          int     `json:"fastestLegendary"`
	HadAfkTeammate                            int     `json:"hadAfkTeammate"`
	HighestChampionDamage                     int     `json:"highestChampionDamage"`
	HighestCrowdControlScore                  int     `json:"highestCrowdControlScore"`
	HighestWardKills                          int     `json:"highestWardKills"`
	JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle"`
	KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler"`
	LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                            int     `json:"legendaryCount"`
	MaxCsAdvantageOnLaneOpponent              float32 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
	MostWardsDestroyedOneSweeper              int     `json:"mostWardsDestroyedOneSweeper"`
	MythicItemUsed                            int     `json:"mythicItemUsed"`
	PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
	SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
	TakedownsFirst25Minutes                   int     `json:"takedownsFirst25Minutes"`
	TeleportTakedowns                         int     `json:"teleportTakedowns"`
	ThirdInhibitorDestroyedTime               int     `json:"thirdInhibitorDestroyedTime"`
	ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
	VisionScoreAdvantageLaneOpponent          float32 `json:"visionScoreAdvantageLaneOpponent"`
	InfernalScalePickup                       int     `json:"InfernalScalePickup"`
	FistBumpParticipation                     int     `json:"fistBumpParticipation"`
	VoidMonsterKill                           int     `json:"voidMonsterKill"`
	AbilityUses                               int     `json:"abilityUses"`
	AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                  float32 `json:"alliedJungleMonsterKills"`
	BaronTakedowns                            int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                                int     `json:"bountyGold"`
	BuffsStolen                               int     `json:"buffsStolen"`
	CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
	ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
	DamagePerMinute                           float32 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage               float32 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                                int     `json:"doubleAces"`
	DragonTakedowns                           int     `json:"dragonTakedowns"`
	LegendaryItemUsed                         []int   `json:"legendaryItemUsed"`
	EffectiveHealAndShielding                 float32 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                   float32 `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
	EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
	FirstTurretKilled                         int     `json:"firstTurretKilled"`
	FirstTurretKilledTime                     float32 `json:"firstTurretKilledTime"`
	FlawlessAces                              int     `json:"flawlessAces"`
	FullTeamTakedown                          int     `json:"fullTeamTakedown"`
	GameLength                                float32 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                             float32 `json:"goldPerMinute"`
	HadOpenNexus                              int     `json:"hadOpenNexus"`
	ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
	InitialBuffCount                          int     `json:"initialBuffCount"`
	InitialCrabCount                          int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                   float32 `json:"jungleCsBefore10Minutes"`
	JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	KDA                                       float32 `json:"kda"`
	KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
	KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
	KillingSprees                             int     `json:"killingSprees"`
	KillParticipation                         float32 `json:"killParticipation"`
	KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
	KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
	KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
	KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
	LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
	LostAnInhibitor                           int     `json:"lostAnInhibitor"`
	MaxKillDeficit                            int     `json:"maxKillDeficit"`
	MejaisFullStackInTime                     int     `json:"mejaisFullStackInTime"`
	MoreEnemyJungleThanOpponent               float32 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
	Multikills                                int     `json:"multikills"`
	MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
	MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
	OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                          int     `json:"outnumberedKills"`
	OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                               int     `json:"perfectGame"`
	PickKillWithAlly                          int     `json:"pickKillWithAlly"`
	PoroExplosions                            int     `json:"poroExplosions"`
	QuickCleanse                              int     `json:"quickCleanse"`
	QuickFirstTurret                          int     `json:"quickFirstTurret"`
	QuickSoloKills                            int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
	SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
	ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
	ShortestTimeToAceFromFirstTakedown        float32 `json:"shortestTimeToAceFromFirstTakedown"`
	SkillshotsDodged                          int     `json:"skillshotsDodged"`
	SkillshotsHit                             int     `json:"skillshotsHit"`
	SnowballsHit                              int     `json:"snowballsHit"`
	SoloBaronKills                            int     `json:"soloBaronKills"`
	SoloKills                                 int     `json:"soloKills"`
	StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
	SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
	TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
	Takedowns                                 int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
	TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                            int     `json:"teamBaronKills"`
	TeamDamagePercentage                      float32 `json:"teamDamagePercentage"`
	TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
	TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
	TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
	TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
	TurretTakedowns                           int     `json:"turretTakedowns"`
	TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
	TwoWardsOneSweeperCount                   int     `json:"twoWardsOneSweeperCount"`
	UnseenRecalls                             int     `json:"unseenRecalls"`
	VisionScorePerMinute                      float32 `json:"visionScorePerMinute"`
	WardsGuarded                              int     `json:"wardsGuarded"`
	WardTakedowns                             int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
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
	Horde      Objective `json:"horde"`
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

type TimelineDTO struct {
	Metadata *MetadataTimeline `json:"metadata"`
	Info     *InfoTimeline     `json:"info"`
}

type MetadataTimeline struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type InfoTimeline struct {
	EndOfGameResult string                 `json:"endOfGameResult"`
	Interval        int64                  `json:"frameInterval"`
	GameId          int64                  `json:"GameId"`
	Participants    []*ParticipantTimeline `json:"participants"`
	Frames          []*MatchFrame          `json:"frames"`
}

type ParticipantTimeline struct {
	ParticipantId int    `json:"participantId"`
	PUUID         string `json:"puuid"`
}

// MatchFrame is a single frame in the timeline of a game
type MatchFrame struct {
	Events            []*MatchEvent                `json:"events"`
	ParticipantFrames map[string]*ParticipantFrame `json:"participantFrames"`
	Timestamp         int                          `json:"timestamp"`
}

type ParticipantFrame struct {
	ChampionStats            *ChampionStats `json:"championStats"`
	CurrentGold              int            `json:"currentGold"`
	DamageStats              *DamageStats   `json:"damageStats"`
	GoldPerSecond            int            `json:"goldPerSecond"`
	JungleMinionsKilled      int            `json:"jungleMinionsKilled"`
	Level                    int            `json:"level"`
	MinionsKilled            int            `json:"minonsKilled"`
	ParticipantId            int            `json:"participantId"`
	Position                 *MatchPosition `json:"position"`
	TimeEnemySpentControlled int            `json:"timeEnemySpentControlled"`
	TotalGold                int            `json:"totalGold"`
	Xp                       int            `json:"xp"`
}

type ChampionStats struct {
	AbilityHaste         int `json:"abilityHaste"`
	AbilityPower         int `json:"abilityPower"`
	Armor                int `json:"armor"`
	ArmorPen             int `json:"armorPen"`
	ArmorPenPercent      int `json:"armorPenPercent"`
	AttackDamage         int `json:"attackDamage"`
	AttackSpeed          int `json:"attackSpeed"`
	BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
	BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
	CcReduction          int `json:"ccReduction"`
	CooldownReduction    int `json:"cooldownReduction"`
	Health               int `json:"health"`
	HealthMax            int `json:"healthMax"`
	HealthRegen          int `json:"healthRegen"`
	Lifesteal            int `json:"lifesteal"`
	MagicPen             int `json:"magicPen"`
	MagicPenPercent      int `json:"magicPenPercent"`
	MagicResist          int `json:"magicResist"`
	MovementSpeed        int `json:"movementSpeed"`
	Omnivamp             int `json:"omnivamp"`
	PhysicalVamp         int `json:"physicalVamp"`
	Power                int `json:"power"`
	PowerMax             int `json:"powerMax"`
	PowerRegen           int `json:"powerRegen"`
	SpellVamp            int `json:"spellVamp"`
}

type DamageStats struct {
	MagicDamageDone               int `json:"magicDamageDone"`
	MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
	MagicDamageTaken              int `json:"magicDamageTaken"`
	PhysicalDamageDone            int `json:"physicalDamageDone"`
	PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
	PhysicalDamageTaken           int `json:"physicalDamageTaken"`
	TotalDamageDone               int `json:"totalDamageDone"`
	TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
	TotalDamageTaken              int `json:"totalDamageTaken"`
	TrueDamageDone                int `json:"trueDamageDone"`
	TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
	TrueDamageTaken               int `json:"trueDamageTaken"`
}

type MatchEvent struct {
	Timestamp     int64  `json:"timestamp"`
	RealTimestamp int64  `json:"realTimestamp"`
	Type          string `json:"type"`
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