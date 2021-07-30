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
	client := pagecall.NewPageCallClient(key)

	temp := fmt.Sprintf("%d", time.Now().Unix())

	newUser, err := client.CreateUser(temp, temp)

	assert.NoError(t, err)

	user, err := client.GetUser(newUser.UserID)

	assert.NoError(t, err)

	users, err := client.GetUsers()

	assert.NoError(t, err)
	assert.NotEqual(t, len(users), 0)

	newRoom, err := client.CreateRoom(pagecall.PrivateRoomType, "SDK Test Room", layoutID, false, []string{})

	assert.NoError(t, err)

	room, err := client.GetRoom(newRoom.ID)

	assert.NoError(t, err)

	rooms, err := client.GetRooms()

	assert.NoError(t, err)
	assert.NotEqual(t, len(rooms), 0)

	member, err := client.JoinRoom(room.ID, user.UserID, nil, nil)

	assert.NoError(t, err)

	members, err := client.GetMembers(room.ID, 0, 10)

	assert.NoError(t, err)
	assert.NotEqual(t, len(members), 0)

	url := client.BuildURLToJoinRoom(member.RoomID, member.AccessToken)

	assert.Equal(t, fmt.Sprintf("%s/%s?access_token=%s", pagecall.AppDomain, member.RoomID, member.AccessToken), url)

	_, err = client.GetLiveSessions(room.ID)

	assert.NoError(t, err)

	_, err = client.TerminateRoom(room.ID)

	assert.NoError(t, err)
}
