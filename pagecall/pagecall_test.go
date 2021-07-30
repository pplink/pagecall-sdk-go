package pagecall_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pplink/pagecall-sdk-go/pagecall"
	"github.com/stretchr/testify/assert"
)

const key string = "pagecall_api_key"
const layoutID string = "pagecall_layout_id"

func TestPageCallSDK(t *testing.T) {
	userID := fmt.Sprintf("%d", time.Now().Unix())
	userName := fmt.Sprintf("%d", time.Now().Unix())
	roomName := fmt.Sprintf("%d", time.Now().Unix())

	client := pagecall.NewPageCallClient(key)

	newUser, err := client.CreateUser(userID, userName)

	assert.NoError(t, err)
	assert.Equal(t, newUser.UserID, userID)

	user, err := client.GetUser(newUser.UserID)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, newUser.ID)

	users, err := client.GetUsers()

	assert.NoError(t, err)
	assert.NotEqual(t, len(users), 0)

	newRoom, err := client.CreateRoom(pagecall.PrivateRoomType, roomName, layoutID, false, []string{})

	assert.NoError(t, err)
	assert.Equal(t, newRoom.Name, roomName)

	room, err := client.GetRoom(newRoom.ID)

	assert.NoError(t, err)
	assert.Equal(t, room.ID, newRoom.ID)

	rooms, err := client.GetRooms()

	assert.NoError(t, err)
	assert.NotEqual(t, len(rooms), 0)

	member, err := client.JoinRoom(room.ID, user.UserID, nil, nil)

	assert.NoError(t, err)
	assert.Equal(t, member.UserID, user.UserID)
	assert.Equal(t, member.RoomID, room.ID)

	members, err := client.GetMembers(room.ID, 0, 10)

	assert.NoError(t, err)
	assert.NotEqual(t, len(members), 0)

	url := client.BuildURLToJoinRoom(member.RoomID, member.AccessToken)

	assert.Equal(t, fmt.Sprintf("%s/%s?access_token=%s", pagecall.AppDomain, member.RoomID, member.AccessToken), url)

	_, err = client.GetLiveSessions(room.ID)

	assert.NoError(t, err)

	terminatedRoom, err := client.TerminateRoom(room.ID)

	assert.NoError(t, err)
	assert.Equal(t, terminatedRoom.IsTerminated, true)
}
