package constants

var (
	// DefaultPlayerAmount is the default amount in dollars a player is worth
	DefaultPlayerAmount = 1000000

	PlayerMinAge = 18
	PlayerMaxAge = 40

	PlayerTypeGoalKeeper = "goalkeeper"
	PlayerTypeDefender   = "defender"
	PlayerTypeMidfielder = "midfielder"
	PlayerTypeAttacker   = "attacker"

	DefaultTeamPlayerMapping = map[string]int{
		PlayerTypeGoalKeeper: 3,
		PlayerTypeDefender:   6,
		PlayerTypeMidfielder: 6,
		PlayerTypeAttacker:   5,
	}
)
