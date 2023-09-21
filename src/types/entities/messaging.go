package entities

type Member struct {
	UserUUID string `json:"userUuid"`
	RoomUUID string `json:"roomUuid"`
}

type Room struct {
	UUID          string     `json:"uuid"`
	Members       []*Member  `json:"members" `
	Messages      []*Message `json:"messages" `
	CreatedAtNano float64    `json:"createdAtNano"`
}

type Message struct {
	UUID          string    `json:"uuid"`
	UserUUID      string    `json:"userUuid" `
	RoomUUID      string    `json:"roomUuid" `
	MessageText   string    `json:"messageText" `
	CreatedAtNano float64   `json:"createdAtNano"`
	MessageStatus string    `json:"messageStatus"`
	SeenBy        []*SeenBy `json:"seenBy"`
}

type SeenBy struct {
	MessageUUID string `json:"messageUuid"`
	UserUUID    string `json:"userUuid"`
}
