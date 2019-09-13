package dota_api

import (
	"encoding/json"
	"strconv"
)

// http://www.sharonkuo.me/dota2/matchdetails.html info from here.

type GameMode int

const (
	AnyMode = 0

	AllPick GameMode = iota
	CaptainsMode
	RandomDraft
	SingleDraft
	AllRandom
	Intro
	TheDiretide
	ReverseCaptainsMode
	Greeviling
	TutorialMode
	MidOnly
	LeastPlayed
	NewPlayerPool
	CompendiumMatchmaking
	Custom
	CaptainsDraft
	BalancedDraft
	AbilityDraft
	Event
	AllRandomDeathMatch
	PvPMid
	RankedAllPick
)

type SkillBracket uint

const (
	AnySkill SkillBracket = iota
	Normal
	High
	VeryHigh
)

type LeaverStatus uint

const (
	None LeaverStatus = iota
	Disconnected
	DisconnectedTooLong
	Abandoned
	AFK
	NeverConnected
	NeverConnectedTooLong
)

type LobbyType int

const (
	Invalid LobbyType = -1

	PublicMatchMaking LobbyType = iota
	Practice
	Tournament
	Tutorial
	CoopWithBots
	TeamMatch
	SoloQueue
	Ranked
	SoloMidPvP
)

type PlayerSlot uint8

// if Radiant, then the first bit is 0/false; if Dire, then the first bit is 1/true
func (d PlayerSlot) IsDire() bool {
	return d&(1<<7) > 0
}

// if Radiant, then the first bit is 0/false; if Dire, then the first bit is 1/true
func (d PlayerSlot) IsRadiant() bool {
	return !d.IsDire()
}

// The last three bits represent the player's position in the team (decimal 0 - 4).
// If a player's player_slot is 129 in decimal (which is 10000001 in binary),
// then they are the second player on the Dire team.
func (d PlayerSlot) GetPosition() (p uint) {
	p = uint(d & ((1 << 7) - 1))
	return
}

type Team uint

const (
	Radiant Team = iota
	Dire
)

// TODO: add methods that read information from bits
type TowerStatus uint16

// TODO: add methods that read information from bits
type BarracksStatus uint16

type MatchId int64

func CreateMatchId(match []byte) (MatchId, error) {
	var matchId MatchId
	err := json.Unmarshal(match, &matchId)

	return matchId, err
}

func CreateMatchId(match int64) (MatchId, error) {
	var matchId MatchId
	err := json.Unmarshal(match, &matchId)

	return matchId, err
}

func (m MatchId) ToString() string {
	return strconv.FormatInt(int64(m), 10)
}

type MatchSequenceNo int64

func (m MatchSequenceNo) ToString() string {
	return strconv.FormatInt(int64(m), 10)
}

type AccountId uint32
