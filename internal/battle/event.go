package battle

import "fmt"

type EventPayload []byte

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
