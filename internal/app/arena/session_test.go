package arena

import (
	"testing"
	"time"
)

func TestSessionReportDurationUsesFinishedAtWhenSet(t *testing.T) {
	startedAt := time.Date(2026, time.June, 10, 12, 0, 0, 0, time.UTC)
	finishedAt := startedAt.Add(2 * time.Second)

	report := sessionReport{
		startedAt:  startedAt,
		finishedAt: finishedAt,
	}

	if report.duration() != 2*time.Second {
		t.Fatalf("expected duration %s, got %s", 2*time.Second, report.duration())
	}
}

func TestSessionReportDurationUsesCurrentTimeWhenNotFinished(t *testing.T) {
	report := sessionReport{
		startedAt: time.Now().Add(-time.Second),
	}

	if report.duration() <= 0 {
		t.Fatal("expected positive duration")
	}
}
