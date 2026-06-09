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
		},
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
