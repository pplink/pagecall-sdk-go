package pagecall

import (
	"encoding/json"
	"fmt"
)

type inlineSession struct {
	ID                   string `json:"_id"`
	ElaspsedTime         int    `json:"elapsed_time"`
	SubscribedCanvasTime int    `json:"subscribed_canvas_time"` // Deprecated
	SubscribedMediaSize  int    `json:"subscribed_media_size"`
	LastPingedAt         string `json:"last_pinged_at"`
	RoomID               string `json:"room_id"`
	UserID               string `json:"user_id"`
	MemberID             string `json:"member_id"`
	ApplicationID        string `json:"application_id"`
	OrganizationID       string `json:"organization_id"`
	IPAdress             string `json:"ip_address"`
	UserAgent            string `json:"user_agent"`
	ConnectionID         string `json:"connection_id"`
	ConnectedAt          string `json:"connected_at"`
	AppVersion           string `json:"app_version"`
	StartUsingCanvasAt   string `json:"start_using_canvas_at,omitempty"`
	V                    string `json:"_v"`
	DisconnectedAt       string `json:"disconnected_at,omitempty"`
}

type member struct {
	ApplicationID  string                 `json:"application_id"`
	UserID         string                 `json:"user_id"`
	AccessToken    string                 `json:"access_token"`
	CreatedAt      string                 `json:"created_at"`
	Name           string                 `json:"name"`
	OrganizationID string                 `json:"organization_id"`
	UpdatedAt      string                 `json:"updated_at"`
	ID             string                 `json:"id"`
	RoomID         string                 `json:"room_id"`
	IsAnonymous    bool                   `json:"is_anonymous"`
	Options        map[string]interface{} `json:"options"`
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
