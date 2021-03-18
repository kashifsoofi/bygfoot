package models

import (
	"github.com/google/uuid"
)

// @see team_return_league_cup_value_int()
type LeagueCupValue int

const (
	LEAGUE_CUP_VALUE_NAME LeagueCupValue = iota
	LEAGUE_CUP_VALUE_SHORT_NAME
	LEAGUE_CUP_VALUE_SID
	LEAGUE_CUP_VALUE_SYMBOL
	LEAGUE_CUP_VALUE_ID
	LEAGUE_CUP_VALUE_FIRST_WEEK
	LEAGUE_CUP_VALUE_LAST_WEEK
	LEAGUE_CUP_VALUE_WEEK_GAP
	LEAGUE_CUP_VALUE_YELLOW_RED
	LEAGUE_CUP_VALUE_AVERAGE_SKILL
	LEAGUE_CUP_VALUE_AVERAGE_CAPACITY
	LEAGUE_CUP_VALUE_SKILL_DIFF
	LEAGUE_CUP_VALUE_END
)

// Some team attributes.
type TeamAttribute int

const (
	TEAM_ATTRIBUTE_STYLE TeamAttribute = iota
	TEAM_ATTRIBUTE_BOOST
	TEAM_ATTRIBUTE_END
)

// The stadium of a team.
type Stadium struct {
	name                string
	capacity            int     // How many people fit in. Default: -1 (depends on league).
	average_attendance  int     // How many people watched on average. Default: 0.
	possible_attendance int     // How many people would've watched if every game had been sold out. We need this only to compute the average attendance in percentage of the capacity. Default: 0.
	games               int     // Number of games. Default: 0.
	safety              float64 // Safety percentage between 0 and 100. Default: randomized.
	ticket_price        float64
}

// Structure representing a team.
// @see Player
type Team struct {
	name           string
	symbol         string
	namesFile      string // File the team takes the player names from.
	defFile        string
	strategySid    string    // The sid of the strategy if it's a CPU team.
	leagueId       uuid.UUID // Numerical id of the league or cup the team belongs to.
	id             uuid.UUID // Id of the team.
	structure      int       // Playing structure. @see team_assign_playing_structure()
	style          int       // Playing style. @see team_assign_playing_style()
	boost          int       // Whether player boost or anti-boost is switched on.
	average_talent float64   // Average talent of the players at generation.
	luck           float64   // A value that influences scoring chances etc. If > 1, the team's lucky, if < 1, it's unlucky. Only used for users' teams.
	stadium        *Stadium
	players        []Player
}
