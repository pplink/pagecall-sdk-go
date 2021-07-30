package pagecall

import (
	"encoding/json"
	"fmt"
)

type session struct {
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
	StartUsingCanvasAt   string `json:"start_using_canvas_at"`
	ID                   string `json:"id"`
}

func (p pageCallClient) GetLiveSessions(roomID string) ([]session, error) {
	path := fmt.Sprintf("/rooms/%s/sessions", roomID)
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
