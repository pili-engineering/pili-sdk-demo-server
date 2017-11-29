package cli

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	pili2 "github.com/pili-engineering/pili-sdk-go.v2/pili"
	"time"
)

type CreateRoomResp struct {
	Room string `json:"room_name"`
}

type DeleteRoomResp struct {
}

type RoomStatusResp struct {
	Room        string `json:"room_name"`
	OwnerUserID string `json:"owner_id"`
	UserMax     int    `json:"user_max"`
	Status      int    `json:"room_status"`
}

type RoomAccessPolicy struct {
	Room     string `json:"room_name"`
	User     string `json:"user_id"`
	Perm     string `json:"perm"`
	Version  string `json:"version"`
	ExpireAt int64  `json:"expire_at"`
}

type CreateRoomArgs struct {
	User    string `json:"owner_id"`
	Room    string `json:"room_name"`
	UserMax int    `json:"user_max"`
}

func RoomStatus(mac *pili2.MAC, room string) (status RoomStatusResp, err error) {
	client := pili2.New(mac, nil)
	err = client.Call(&status, "GET", "http://rtc.qiniuapi.com/v1/rooms/"+room)
	return
}

func RoomCreate(mac *pili2.MAC, room, user string, max int) (ret CreateRoomResp, err error) {
	client := pili2.New(mac, nil)
	err = client.CallWithJSON(&ret, "POST", "http://rtc.qiniuapi.com/v1/rooms", CreateRoomArgs{Room: room, User: user, UserMax: max})
	return
}

func RoomDelete(mac *pili2.MAC, room string) (ret DeleteRoomResp, err error) {
	client := pili2.New(mac, nil)
	err = client.Call(&ret, "DELETE", "http://rtc.qiniuapi.com/v1/rooms/"+room)
	return
}

func CreateToken(mac *pili2.MAC, room, user, version string) string {
	policy := RoomAccessPolicy{
		Room:     room,
		User:     user,
		Perm:     "user",
		Version:  version,
		ExpireAt: time.Now().Add(time.Hour * 24).UnixNano(),
	}
	b, _ := json.Marshal(policy)
	return signWithData(b, mac)
}

func signWithData(b []byte, mac *pili2.MAC) (token string) {

	blen := base64.URLEncoding.EncodedLen(len(b))

	key := mac.AccessKey
	nkey := len(key)
	ret := make([]byte, nkey+30+blen)

	base64.URLEncoding.Encode(ret[nkey+30:], b)

	h := hmac.New(sha1.New, mac.SecretKey)
	h.Write(ret[nkey+30:])
	digest := h.Sum(nil)

	copy(ret, key)
	ret[nkey] = ':'
	base64.URLEncoding.Encode(ret[nkey+1:], digest)
	ret[nkey+29] = ':'

	return string(ret)
}
