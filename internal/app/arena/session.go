package arena

import (
	"fmt"
	"io"
	"os"
	"time"
)

type sessionReport struct {
	startedAt  time.Time
	finishedAt time.Time
}

func startSessionReport() *sessionReport {
	return &sessionReport{
		startedAt: time.Now(),
	}
}

func (report *sessionReport) finish() {
	report.finishedAt = time.Now()
}

func (report sessionReport) duration() time.Duration {
	if report.finishedAt.IsZero() {
		return time.Since(report.startedAt)
	}

	return report.finishedAt.Sub(report.startedAt)
}

func (report sessionReport) print() {
	report.printTo(os.Stdout)
}

func (report sessionReport) printTo(output io.Writer) {
	fmt.Fprintln(output, "Arena session duration:", report.duration())
}
