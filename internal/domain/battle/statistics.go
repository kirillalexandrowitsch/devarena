package battle

type Statistics struct {
	Wins   int
	Losses int
}

func (s Statistics) TotalBattles() int {
	return s.Wins + s.Losses
}

func (s Statistics) WinRatePercent() float64 {
	totalBattles := s.TotalBattles()
	if totalBattles == 0 {
		return 0
	}

	return float64(s.Wins) / float64(totalBattles) * 100
}
