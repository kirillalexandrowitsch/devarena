package battle

import "fmt"

const RewardGrantedEventType = "battle.reward_granted"

type EventPayload []byte

type Event struct {
	Type     string
	Payload  EventPayload
	Metadata map[string]any
}

func NewRewardEvent(heroName string, reward Reward) Event {
	payload := NewRewardEventPayload(heroName, reward)

	return Event{
		Type:    RewardGrantedEventType,
		Payload: payload,
		Metadata: map[string]any{
			"hero_name":         heroName,
			"reward_item":       reward.Item,
			"reward_experience": reward.Experience,
			"payload_size":      payload.Size(),
			"reward_granted":    true,
		},
	}
}

func (event Event) MetadataString(key string) (string, bool) {
	value, exists := event.Metadata[key]
	if !exists {
		return "", false
	}

	stringValue, ok := value.(string)

	return stringValue, ok
}

func (event Event) MetadataInt(key string) (int, bool) {
	value, exists := event.Metadata[key]
	if !exists {
		return 0, false
	}

	intValue, ok := value.(int)

	return intValue, ok
}

func (event Event) MetadataText(key string) (string, bool) {
	value, exists := event.Metadata[key]
	if !exists {
		return "", false
	}

	switch typedValue := value.(type) {
	case string:
		return typedValue, true
	case int:
		return fmt.Sprintf("%d", typedValue), true
	case bool:
		return fmt.Sprintf("%t", typedValue), true
	case EventPayload:
		return typedValue.String(), true
	default:
		return "", false
	}
}

func NewRewardEventPayload(heroName string, reward Reward) EventPayload {
	message := fmt.Sprintf(
		"hero=%s;experience=%d;item=%s",
		heroName,
		reward.Experience,
		reward.Item,
	)

	return EventPayload([]byte(message))
}

func (payload EventPayload) Size() int {
	return len(payload)
}

func (payload EventPayload) String() string {
	return string(payload)
}
