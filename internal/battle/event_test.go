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

func TestNewRewardEvent(t *testing.T) {
	reward := Reward{
		Experience: 30,
		Item:       "Goblin Dagger",
	}

	event := NewRewardEvent("Ragnar", reward)

	if event.Type != RewardGrantedEventType {
		t.Fatalf("expected event type %q, got %q", RewardGrantedEventType, event.Type)
	}

	if event.Payload.String() != "hero=Ragnar;experience=30;item=Goblin Dagger" {
		t.Fatalf("unexpected event payload %q", event.Payload.String())
	}

	if event.Metadata["hero_name"] != "Ragnar" {
		t.Fatalf("expected hero_name metadata %q, got %v", "Ragnar", event.Metadata["hero_name"])
	}

	if event.Metadata["reward_item"] != "Goblin Dagger" {
		t.Fatalf("expected reward_item metadata %q, got %v", "Goblin Dagger", event.Metadata["reward_item"])
	}

	if event.Metadata["reward_experience"] != 30 {
		t.Fatalf("expected reward_experience metadata %d, got %v", 30, event.Metadata["reward_experience"])
	}

	if event.Metadata["payload_size"] != event.Payload.Size() {
		t.Fatalf("expected payload_size metadata %d, got %v", event.Payload.Size(), event.Metadata["payload_size"])
	}
}
