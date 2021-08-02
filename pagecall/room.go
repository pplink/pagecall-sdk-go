package pagecall

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type room struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	RoomType                 string   `json:"type"`
	LayoutID                 string   `json:"layout_id"`
	ApplicationID            string   `json:"application_id"`
	OrganizationID           string   `json:"organization_id"`
	IsDistinct               bool     `json:"is_distinct"`
	ThumbnailURL             string   `json:"thumbnail_url"`
	LiveTime                 int      `json:"live_time"`
	LiveTimeSectionStartedAt string   `json:"live_time_section_started_at"`
	IsRecurring              bool     `json:"is_recurring"`
	IsTerminated             bool     `json:"is_terminated"`
	TerminatedAt             string   `json:"terminated_at"`
	CreatedAt                string   `json:"created_at"`
	UpdatedAt                string   `json:"updated_at"`
	DistinctUserIDs          []string `json:"distinct_user_ids"`
	InitialPages             []string `json:"initial_pages"`
	Members                  []member `json:"members"`
}

func (p pageCallClient) CreateRoom(roomType string, name string, layoutID string) (*room, error) {
	reqBody := make(map[string]interface{})
	reqBody["type"] = roomType
	reqBody["name"] = name
	reqBody["layout_id"] = layoutID

	ubytes, _ := json.Marshal(reqBody)
	payload := bytes.NewBuffer(ubytes)

	body, err := p.request("POST", "/rooms", payload)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Room room `json:"room"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return &respBody.Room, nil
}

func (p pageCallClient) JoinRoom(roomID string, userID string, layoutID *string, options map[string]interface{}) (*member, error) {
	reqBody := make(map[string]interface{})
	reqBody["user_id"] = userID

	if layoutID != nil {
		reqBody["layout_id"] = *layoutID
	}

	if options != nil {
		reqBody["options"] = options
	}

	ubytes, _ := json.Marshal(reqBody)
	payload := bytes.NewBuffer(ubytes)

	path := fmt.Sprintf("/rooms/%s/members", roomID)
	body, err := p.request("POST", path, payload)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Member member `json:"member"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return &respBody.Member, nil
}

func (p pageCallClient) TerminateRoom(roomID string) (*room, error) {
	reqBody := make(map[string]interface{})
	reqBody["is_terminated"] = true

	ubytes, _ := json.Marshal(reqBody)
	payload := bytes.NewBuffer(ubytes)

	path := fmt.Sprintf("/rooms/%s", roomID)
	body, err := p.request("PUT", path, payload)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Room room `json:"room"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return &respBody.Room, nil
}

func (p pageCallClient) GetRoom(roomID string) (*room, error) {
	path := fmt.Sprintf("/rooms/%s", roomID)
	body, err := p.request("GET", path, nil)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Room room `json:"room"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return &respBody.Room, nil
}

func (p pageCallClient) GetRooms(offset int, limit int) ([]room, error) {
	path := fmt.Sprintf("/rooms?offset=%d&limit=%d", offset, limit)
	body, err := p.request("GET", path, nil)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Rooms []room `json:"rooms"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return respBody.Rooms, nil
}
