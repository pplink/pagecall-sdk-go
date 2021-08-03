package pagecall

import (
	"encoding/json"
	"fmt"
)

type session struct {
	ID                  string `json:"id"`
	UserID              string `json:"user_id"`
	MemberID            string `json:"member_id"`
	RoomID              string `json:"room_id"`
	ApplicationID       string `json:"application_id"`
	OrganizationID      string `json:"organization_id"`
	SubscribedMediaSize int    `json:"subscribed_media_size"`
	ConnectionID        string `json:"connection_id"`
	ConnectedAt         string `json:"connected_at"`
	StartUsingCanvasAt  string `json:"start_using_canvas_at"`
	LastPingedAt        string `json:"last_pinged_at"`
	IPAdress            string `json:"ip_address"`
	UserAgent           string `json:"user_agent"`
	AppVersion          string `json:"app_version"`
}

func (p pageCallClient) GetLiveSessions(roomID string, offset int, limit int) ([]session, error) {
	path := fmt.Sprintf("/rooms/%s/sessions?is_connecting=true&offset=%d&limit=%d", roomID, offset, limit)
	body, err := p.request("GET", path, nil)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Sessions []session `json:"sessions"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return respBody.Sessions, nil
}
