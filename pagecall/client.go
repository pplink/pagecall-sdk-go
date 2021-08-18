package pagecall

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const OpenRoomType string = "open"
const PublicRoomType string = "public"
const PrivateRoomType string = "private"

const AppDomain string = "https://app.pagecall.net"
const ApiDomain string = "https://api.pagecall.net"
const ApiVersion string = "v1"

type PageCallClient interface {
	/*
		Create an user.
	*/
	CreateUser(id string, name string) (*user, error)

	/*
		Create a room.
	*/
	CreateRoom(roomType string, name string, layoutID string) (*room, error)

	/*
		Get user's information.
	*/
	GetUser(userID string) (*user, error)

	/*
		Get a list of users.
	*/
	GetUsers(offset int, limit int) ([]user, error)

	/*
		Get room's information.
	*/
	GetRoom(roomID string) (*room, error)

	/*
		Get a list of rooms.
	*/
	GetRooms(offset int, limit int) ([]room, error)

	/*
		Get a list of members in the room.
	*/
	GetMembers(roomID string, offset int, limit int) ([]member, error)

	/*
		Get a list of member sessions currently in the room.
	*/
	GetLiveSessions(roomID string, offset int, limit int) ([]session, error)

	/*
		Add an user to the room.

		If layoutID is set, override room layout.
	*/
	JoinRoom(roomID string, userID string, layoutID *string, options map[string]interface{}) (*member, error)

	/*
		Create an URL to access the room.
	*/
	BuildURLToJoinRoom(roomID string, accessToken string) string

	/*
		Terminate the room.
	*/
	TerminateRoom(roomID string) (*room, error)

	/*
		Using the post action to sessions feature, the according script can be executed within the connected session client.
	*/
	PostActionToSessions(sessionIDs []string, script string) error
}

type pageCallClient struct {
	key string
}

func NewPageCallClient(key string) PageCallClient {
	return &pageCallClient{key}
}

func (p pageCallClient) request(method string, path string, payload io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s%s", ApiDomain, ApiVersion, path)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	authorization := fmt.Sprintf("Bearer %s", p.key)

	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	return body, nil
}

func (p pageCallClient) BuildURLToJoinRoom(roomID string, accessToken string) string {
	return fmt.Sprintf("%s/%s?access_token=%s", AppDomain, roomID, accessToken)
}

func (p pageCallClient) PostActionToSessions(sessionIDs []string, script string) error {
	reqBody := make(map[string]interface{})
	reqBody["type"] = "run_script"
	reqBody["session_ids"] = sessionIDs
	reqBody["script"] = script

	ubytes, _ := json.Marshal(reqBody)
	payload := bytes.NewBuffer(ubytes)

	_, err := p.request("POST", "/post_action_to_sessions", payload)

	if err != nil {
		return err
	}

	return nil
}
