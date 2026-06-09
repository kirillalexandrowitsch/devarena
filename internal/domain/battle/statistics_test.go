package battle

import "testing"

func TestStatisticsTotalBattles(t *testing.T) {
	statistics := Statistics{
		Wins:   3,
		Losses: 2,
	}

	got := statistics.TotalBattles()
	want := 5

	if got != want {
		t.Fatalf("expected total battles %d, got %d", want, got)
	}
}

func TestStatisticsWinRatePercent(t *testing.T) {
	statistics := Statistics{
		Wins:   2,
		Losses: 1,
	}

	got := statistics.WinRatePercent()
	want := 66.66666666666666

	if got != want {
		t.Fatalf("expected win rate %f, got %f", want, got)
	}
}

func TestStatisticsWinRatePercentWithoutBattles(t *testing.T) {
	statistics := Statistics{}

	got := statistics.WinRatePercent()
	want := 0.0

	if got != want {
		t.Fatalf("expected win rate %f, got %f", want, got)
	}
}
