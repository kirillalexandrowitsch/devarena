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

	if event.Metadata["reward_granted"] != true {
		t.Fatalf("expected reward_granted metadata %t, got %v", true, event.Metadata["reward_granted"])
	}
}

func TestEventMetadataStringReturnsStringValue(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"hero_name": "Ragnar",
		},
	}

	value, exists := event.MetadataString("hero_name")

	if !exists {
		t.Fatal("expected hero_name metadata to exist")
	}

	if value != "Ragnar" {
		t.Fatalf("expected hero_name metadata %q, got %q", "Ragnar", value)
	}
}

func TestEventMetadataStringReturnsFalseForWrongType(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"reward_experience": 30,
		},
	}

	value, exists := event.MetadataString("reward_experience")

	if exists {
		t.Fatal("expected reward_experience metadata not to be a string")
	}

	if value != "" {
		t.Fatalf("expected empty string value, got %q", value)
	}
}

func TestEventMetadataIntReturnsIntValue(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"reward_experience": 30,
		},
	}

	value, exists := event.MetadataInt("reward_experience")

	if !exists {
		t.Fatal("expected reward_experience metadata to exist")
	}

	if value != 30 {
		t.Fatalf("expected reward_experience metadata %d, got %d", 30, value)
	}
}

func TestEventMetadataIntReturnsFalseForWrongType(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"hero_name": "Ragnar",
		},
	}

	value, exists := event.MetadataInt("hero_name")

	if exists {
		t.Fatal("expected hero_name metadata not to be an int")
	}

	if value != 0 {
		t.Fatalf("expected zero int value, got %d", value)
	}
}

func TestEventMetadataAccessorsReturnFalseForMissingMetadata(t *testing.T) {
	event := Event{}

	stringValue, stringExists := event.MetadataString("hero_name")
	if stringExists {
		t.Fatal("expected missing string metadata to return false")
	}

	if stringValue != "" {
		t.Fatalf("expected empty string value, got %q", stringValue)
	}

	intValue, intExists := event.MetadataInt("reward_experience")
	if intExists {
		t.Fatal("expected missing int metadata to return false")
	}

	if intValue != 0 {
		t.Fatalf("expected zero int value, got %d", intValue)
	}
}

func TestEventMetadataTextReturnsStringValue(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"hero_name": "Ragnar",
		},
	}

	value, exists := event.MetadataText("hero_name")

	if !exists {
		t.Fatal("expected hero_name metadata text to exist")
	}

	if value != "Ragnar" {
		t.Fatalf("expected hero_name metadata text %q, got %q", "Ragnar", value)
	}
}

func TestEventMetadataTextReturnsIntValue(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"reward_experience": 30,
		},
	}

	value, exists := event.MetadataText("reward_experience")

	if !exists {
		t.Fatal("expected reward_experience metadata text to exist")
	}

	if value != "30" {
		t.Fatalf("expected reward_experience metadata text %q, got %q", "30", value)
	}
}

func TestEventMetadataTextReturnsBoolValue(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"reward_granted": true,
		},
	}

	value, exists := event.MetadataText("reward_granted")

	if !exists {
		t.Fatal("expected reward_granted metadata text to exist")
	}

	if value != "true" {
		t.Fatalf("expected reward_granted metadata text %q, got %q", "true", value)
	}
}

func TestEventMetadataTextReturnsPayloadValue(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"payload": EventPayload([]byte("battle finished")),
		},
	}

	value, exists := event.MetadataText("payload")

	if !exists {
		t.Fatal("expected payload metadata text to exist")
	}

	if value != "battle finished" {
		t.Fatalf("expected payload metadata text %q, got %q", "battle finished", value)
	}
}

func TestEventMetadataTextReturnsFalseForUnsupportedType(t *testing.T) {
	event := Event{
		Metadata: map[string]any{
			"tags": []string{"reward", "battle"},
		},
	}

	value, exists := event.MetadataText("tags")

	if exists {
		t.Fatal("expected unsupported metadata type to return false")
	}

	if value != "" {
		t.Fatalf("expected empty metadata text, got %q", value)
	}
}

func TestEventMetadataTextReturnsFalseForMissingMetadata(t *testing.T) {
	event := Event{}

	value, exists := event.MetadataText("missing")

	if exists {
		t.Fatal("expected missing metadata text to return false")
	}

	if value != "" {
		t.Fatalf("expected empty metadata text, got %q", value)
	}
}
