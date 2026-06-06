package battle

import "testing"

func TestNewRewardEventPayload(t *testing.T) {
	reward := Reward{
		Experience: 30,
		Item:       "Goblin Dagger",
	}

	payload := NewRewardEventPayload("Ragnar", reward)

	want := "hero=Ragnar;experience=30;item=Goblin Dagger"

	if payload.String() != want {
		t.Fatalf("expected payload %q, got %q", want, payload.String())
	}
}

func TestEventPayloadSizeCountsBytes(t *testing.T) {
	payload := EventPayload([]byte("Герой"))

	got := payload.Size()
	want := 10

	if got != want {
		t.Fatalf("expected payload size %d bytes, got %d", want, got)
	}
}
