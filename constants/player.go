package constants

type PlayerType string

var (
	// DefaultPlayerAmount is the default amount in dollars a player is worth
	DefaultPlayerAmount = 1000000

	PlayerMinAge = 18
	PlayerMaxAge = 40

	PlayerTypeGoalKeeper PlayerType = "goalkeeper"
	PlayerTypeDefender   PlayerType = "defender"
	PlayerTypeMidfielder PlayerType = "midfielder"
	PlayerTypeAttacker   PlayerType = "attacker"

	DefaultTeamPlayerMapping = map[PlayerType]int{
		PlayerTypeGoalKeeper: 3,
		PlayerTypeDefender:   6,
		PlayerTypeMidfielder: 6,
		PlayerTypeAttacker:   5,
	}

	PlayerMinPercentageIncrease = 10
	PlayerMaxPercentageIncrease = 100
)
