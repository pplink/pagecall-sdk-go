package pagecall_test

import (
	"testing"

	"github.com/pplink/pagecall-sdk-go/pagecall"
)

const key string = "pagecall_api_key"
const layoutID string = "pagecall_layout_id"

func TestCreateRoom(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	room, err := client.CreateRoom(pagecall.PrivateRoomType, "SDK test", layoutID, false, []string{})

	if err != nil {
		t.Error(err)
	}

	_, err = client.TerminateRoom(room.ID)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateUser(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.CreateUser("sdk", "pplinkian")

	if err != nil {
		t.Error(err)
	}
}

func TestJoinRoom(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	room, err := client.CreateRoom(pagecall.PrivateRoomType, "SDK Test", layoutID, false, []string{})

	if err != nil {
		t.Error(err)
	}

	user, err := client.CreateUser("sdk", "pplinkian")

	if err != nil {
		t.Error(err)
	}

	member, err := client.JoinRoom(room.ID, user.UserID, nil, nil)

	if err != nil {
		t.Error(err)
	}

	client.BuildURLToJoinRoom(member.RoomID, member.AccessToken)
	_, err = client.TerminateRoom(room.ID)

	if err != nil {
		t.Error(err)
	}
}

func TestGetLiveSessions(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.GetLiveSessions("roomID")

	if err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.GetUser("sdk")

	if err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.GetUsers()

	if err != nil {
		t.Error(err)
	}
}

func TestGetMembers(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.GetMembers("roomID", 0, 10)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRoom(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.GetRoom("roomID")

	if err != nil {
		t.Error(err)
	}
}

func TestGetRooms(t *testing.T) {
	client := pagecall.NewPageCallClient(key)
	_, err := client.GetRooms()

	if err != nil {
		t.Error(err)
	}
}
