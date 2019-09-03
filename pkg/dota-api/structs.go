package dota_api

import "encoding/json"

type APIResponse struct {
	Result json.RawMessage `json:"result"`
}

type Error struct {
	ErrorMsg string `json:"error"`
}

type MatchResult struct {
	Error
	Players               []Player
	Season                uint `json:"item_0,omitempty"`
	RadiantWin            bool `json:"radiant_win"`
	Duration              uint
	MatchStart            uint           `json:"start_time"`
	MatchId               uint64         `json:"match_id"`
	MatchSequenceNo       uint           `json:"match_seq_num"`
	TowerStatusRadiant    TowerStatus    `json:"tower_status_radiant"`
	TowerStatusDire       TowerStatus    `json:"tower_status_dire"`
	BarracksStatusRadiant BarracksStatus `json:"barracks_status_radiant"`
	BarracksStatusDire    BarracksStatus `json:"barracks_status_dire"`
	Cluster               uint
	FirstBloodTime        int       `json:"first_blood_time"`
	LobbyType             LobbyType `json:"lobby_type"`
	HumanPlayers          uint      `json:"human_players"`
	LeagueId              uint
	PositiveVotes         uint      `json:"positive_votes"`
	NegativeVotes         uint      `json:"negative_votes"`
	GameMode              GameMode  `json:"game_mode"`
	PicksBans             []PickBan `json:"picks_bans,omitempty"`
}

type Player struct {
	AccountId     uint32     `json:"account_id"`
	PlayerSlot    PlayerSlot `json:"player_slot"`
	HeroId        uint       `json:"hero_id"`
	Item0         uint       `json:"item_0"`
	Item1         uint       `json:"item_1"`
	Item2         uint       `json:"item_2"`
	Item3         uint       `json:"item_3"`
	Item4         uint       `json:"item_4"`
	Item5         uint       `json:"item_5"`
	Kills         uint
	Deaths        uint
	Assists       uint
	LeaverStatus  LeaverStatus `json:"leaver_status"`
	GoldRemaining uint         `json:"gold"`
	LastHits      uint         `json:"last_hits"`
	Denies        uint         `json:"denies"`
	GPM           uint         `json:"gold_per_min"`
	XPM           uint         `json:"xp_per_min"`
	GoldSpent     uint         `json:"gold_spent"`
	HeroDamage    uint         `json:"hero_damage"`
	TowerDamage   uint         `json:"tower_damage"`
	HeroHealing   uint         `json:"hero_healing"`
	Level         uint
	Abilities     []Ability `json:"ability_upgrades"`
	Units         []Unit    `json:"additional_units,omitempty"`
}

type Ability struct {
	Id           uint `json:"ability"`
	TimeUpgraded int  `json:"time"`
	Level        uint
}

type Unit struct {
	Name  string `json:"unitname"`
	Item0 uint   `json:"item_0"`
	Item1 uint   `json:"item_1"`
	Item2 uint   `json:"item_2"`
	Item3 uint   `json:"item_3"`
	Item4 uint   `json:"item_4"`
	Item5 uint   `json:"item_5"`
}

type PickBan struct {
	IsPick bool `json:"is_pick"`
	HeroId uint `json:"hero_id"`
	Team   Team
	Order  uint
}
