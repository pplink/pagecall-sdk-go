package pagecall

import (
	"encoding/json"
	"fmt"
)

type inlineSession struct {
	ID                   string `json:"_id"`
	UserID               string `json:"user_id"`
	MemberID             string `json:"member_id"`
	RoomID               string `json:"room_id"`
	OrganizationID       string `json:"organization_id"`
	ApplicationID        string `json:"application_id"`
	ConnectionID         string `json:"connection_id"`
	ConnectedAt          string `json:"connected_at"`
	StartUsingCanvasAt   string `json:"start_using_canvas_at,omitempty"`
	DisconnectedAt       string `json:"disconnected_at,omitempty"`
	LastPingedAt         string `json:"last_pinged_at"`
	ElaspsedTime         int    `json:"elapsed_time"`
	SubscribedCanvasTime int    `json:"subscribed_canvas_time"`
	SubscribedMediaSize  int    `json:"subscribed_media_size"`
	IPAdress             string `json:"ip_address"`
	UserAgent            string `json:"user_agent"`
	AppVersion           string `json:"app_version"`
	V                    string `json:"_v"`
}

type member struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	UserID         string                 `json:"user_id"`
	RoomID         string                 `json:"room_id"`
	OrganizationID string                 `json:"organization_id"`
	ApplicationID  string                 `json:"application_id"`
	AccessToken    string                 `json:"access_token"`
	IsAnonymous    bool                   `json:"is_anonymous"`
	Options        map[string]interface{} `json:"options"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	Sessions       []inlineSession        `json:"sessions,omitempty"`
}

func (p pageCallClient) GetMembers(roomID string, offset int, limit int) ([]member, error) {
	path := fmt.Sprintf("/rooms/%s/members?offset=%d&limit=%d", roomID, offset, limit)
	body, err := p.request("GET", path, nil)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Members []member `json:"members"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return respBody.Members, nil
}
