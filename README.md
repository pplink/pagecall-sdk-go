# PageCall SDK for Go

pagecall-sdk-go is the official PageCall SDK for the Go programming language.

## Installation
```go get -u github.com/pplink/pagecall-sdk-go```

## Prerequisite
- PageCall API Key
- PageCall Layout ID

## Example
```
package main

import (
	"fmt"

	"github.com/pplink/pagecall-sdk-go/pagecall"
)

func main() {
	const apiKey string = "pagecall_api_key"
	const layoutID string = "pagecall_layout_id"

	client := pagecall.NewPageCallClient(apiKey)

	user, err := client.CreateUser("userID", "userName")

	if err != nil {
		panic(err)
	}

	room, err := client.CreateRoom(pagecall.PrivateRoomType, "PageCall SDK Test Room", layoutID)

	if err != nil {
		panic(err)
	}

	member, err := client.JoinRoom(room.ID, user.UserID, nil, nil)

	if err != nil {
		panic(err)
	}

	url := client.BuildURLToJoinRoom(member.RoomID, member.AccessToken)

	fmt.Println(url)
}

```

## Reference

- [PageCall](https://pagecall.net/)
- [PageCall Developer Console](https://console.pagecall.net/)
- [PageCall API Document](https://docs.pagecall.net/)

## License
[Apache-2.0](./LICENSE)
