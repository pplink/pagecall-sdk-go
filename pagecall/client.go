package pagecall

import (
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
	request(method string, path string, payload io.Reader) ([]byte, error)

	/*
		Create an user.
	*/
	CreateUser(id string, name string) (*user, error)

	/*
		Create a room.
	*/
	CreateRoom(roomType string, name string, layoutID string, isDistinct bool, userIDs []string) (*room, error)

	/*
		Get user's information.
	*/
	GetUser(userID string) (*user, error)

	/*
		Get a list of all users.
	*/
	GetUsers() ([]user, error)

	/*
		Get room's information.
	*/
	GetRoom(roomID string) (*room, error)

	/*
		Get a list of all rooms.
	*/
	GetRooms() ([]room, error)

	/*
		Get a list of all members in the room.
	*/
	GetMembers(roomID string, offset int, limit int) ([]member, error)

	/*
		Get a list of member sessions currently in the room.
	*/
	GetLiveSessions(roomID string) ([]session, error)

	/*
		Add an user to the room.
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
