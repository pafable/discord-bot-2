package rpg

type Adventure struct {
	Name        string
	RewardValue int
	Rating      AdventureRating
}

type AdventureRating struct {
	Value string
}

func (p *Player) GetHardAdventure(name string, reward int) Adventure {
	return Adventure{
		Name:        name,
		RewardValue: reward,
		Rating:      AdventureRating{"HARD"},
	}
}
