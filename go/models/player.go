package models

import (
	"github.com/google/uuid"
)

// Player Positions.
type PlayerPosition int

const (
	PLAYER_POS_GOALIE PlayerPosition = iota
	PLAYER_POS_DEFENDER
	PLAYER_POS_MIDFIELDER
	PLAYER_POS_FORWARD
	PLAYER_POS_ANY
	PLAYER_POS_END
)

// Streaks a player can go on.
type PlayerStreak int

const (
	PLAYER_STREAK_COLD PlayerStreak = iota - 1
	PLAYER_STREAK_NONE
	PLAYER_STREAK_HOT
)

// Cards in different cups are counted separately for players;
// for each league or cup the cards are stored in such a struct.
type PlayerCard struct {
	leagueId uuid.UUID
	yellow   int
	red      int
}

// Goals and games in different leagues and cups are counted separately for players.
type PlayerGamesGoals struct {
	leagueId uuid.UUID
	games    int
	goals    int
	shots    int
}

type PlayerInjury int

const (
	PLAYER_INJURY_NONE PlayerInjury = iota
	PLAYER_INJURY_CONCUSSION
	PLAYER_INJURY_PULLED_MUSCLE
	PLAYER_INJURY_HAMSTRING
	PLAYER_INJURY_GROIN
	PLAYER_INJURY_FRAC_ANKLE
	PLAYER_INJURY_RIB
	PLAYER_INJURY_LEG
	PLAYER_INJURY_BROK_ANKLE
	PLAYER_INJURY_ARM
	PLAYER_INJURY_SHOULDER
	PLAYER_INJURY_LIGAMENT
	PLAYER_INJURY_CAREER_STOP
	PLAYER_INJURY_END
)

// Enum for different player data.
type PlayerValue int

const (
	PLAYER_VALUE_GAMES PlayerValue = iota
	PLAYER_VALUE_GOALS
	PLAYER_VALUE_SHOTS
	PLAYER_VALUE_CARD_YELLOW
	PLAYER_VALUE_CARD_RED
	PLAYER_VALUE_END
)

// Enumeration for the yellow/red card status during the live game.
type PlayerCardStatus int

const (
	PLAYER_CARD_STATUS_NONE PlayerCardStatus = iota
	PLAYER_CARD_STATUS_YELLOW
	PLAYER_CARD_STATUS_RED
)

// Representation of a player.
// see #PlayerAttributes
type Player struct {
	name            string
	position        PlayerPosition
	currentPosition PlayerPosition
	health          PlayerInjury
	recovery        int              // Weeks until the player gets healthy.
	id              int              // Id of the player within the team.
	value           int              // Value of the player.
	wage            int              // Wage of the player.
	offers          int              // Number of times the player received a contract offer.
	streak          int              // The streak the player is on.
	cardStatus      PlayerCardStatus // The card status of the player during a live game.
	skill           float64          // Skill. Between 0 and a constant (specified in the constants file).
	currentSkill    float64
	talent          float64            // Talent. The peak ability (which isn't always reached).
	etal            []float64          // Estimated talent (the user never sees the actual talent). Depends on scout quality.
	fitness         float64            // Fitness. Between 0 and 1.
	lsu             float64            // Last skill update. Number of weeks since the player skill was last updated.
	age             int                // Age in years.
	peakAge         int                // Age at which the player reaches his peak ability.
	peakRegion      int                // Region around the peak age during which the player's ability is at the peak (in years).
	contract        int                // The years until the player's contract expires.
	streak_prob     float64            // This number determines how probable it is that a player goes on a hot/cold streak. Between -1 and 1.
	streak_count    int                // How many weeks the streak goes (or how long a new streak may not begin if the value is negative).
	participation   bool               // Whether the player participated in the team's last match.
	gamesGoals      []PlayerGamesGoals // Array of games and goals; one item per league and cup.
	cards           []PlayerCard       // Array of cards; one item per league and cup.
	career          []PlayerValue      // Career goals, games etc.
	team            *Team
}

// Enum for player attributes that can be shown in a player list.
type PlayerListAttributeValue int

const (
	PLAYER_LIST_ATTRIBUTE_NAME PlayerListAttributeValue = iota
	PLAYER_LIST_ATTRIBUTE_CPOS
	PLAYER_LIST_ATTRIBUTE_POS
	PLAYER_LIST_ATTRIBUTE_CSKILL
	PLAYER_LIST_ATTRIBUTE_SKILL
	PLAYER_LIST_ATTRIBUTE_FITNESS
	PLAYER_LIST_ATTRIBUTE_GAMES
	PLAYER_LIST_ATTRIBUTE_SHOTS
	PLAYER_LIST_ATTRIBUTE_GOALS
	PLAYER_LIST_ATTRIBUTE_STATUS
	PLAYER_LIST_ATTRIBUTE_CARDS
	PLAYER_LIST_ATTRIBUTE_AGE
	PLAYER_LIST_ATTRIBUTE_ETAL
	PLAYER_LIST_ATTRIBUTE_VALUE
	PLAYER_LIST_ATTRIBUTE_WAGE
	PLAYER_LIST_ATTRIBUTE_CONTRACT
	PLAYER_LIST_ATTRIBUTE_TEAM
	PLAYER_LIST_ATTRIBUTE_LEAGUE_CUP
	PLAYER_LIST_ATTRIBUTE_END
)

type PlayerInfoAttributeValue int

const (
	PLAYER_INFO_ATTRIBUTE_NAME PlayerInfoAttributeValue = iota
	PLAYER_INFO_ATTRIBUTE_POS
	PLAYER_INFO_ATTRIBUTE_CPOS
	PLAYER_INFO_ATTRIBUTE_SKILL
	PLAYER_INFO_ATTRIBUTE_CSKILL
	PLAYER_INFO_ATTRIBUTE_FITNESS
	PLAYER_INFO_ATTRIBUTE_ETAL
	PLAYER_INFO_ATTRIBUTE_AGE
	PLAYER_INFO_ATTRIBUTE_HEALTH
	PLAYER_INFO_ATTRIBUTE_VALUE
	PLAYER_INFO_ATTRIBUTE_WAGE
	PLAYER_INFO_ATTRIBUTE_CONTRACT
	PLAYER_INFO_ATTRIBUTE_GAMES_GOALS
	PLAYER_INFO_ATTRIBUTE_YELLOW_CARDS
	PLAYER_INFO_ATTRIBUTE_BANNED
	PLAYER_INFO_ATTRIBUTE_STREAK
	PLAYER_INFO_ATTRIBUTE_CAREER
	PLAYER_INFO_ATTRIBUTE_OFFERS
	PLAYER_INFO_ATTRIBUTE_END
)

// A struct telling us which player attributes to show in a player list.
// @see #PlayerListAttributeValue
type PlayerListAttribute struct {
	onOff []bool
}
