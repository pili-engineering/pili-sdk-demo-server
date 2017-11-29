# pili-sdk-demo-server
七牛直播/连麦 SDK 官方 demo 的后台业务服务器代码

## 编译运行
参考 [Deploy.md](./Deploy.md)

***

## 接入
* 对外域名：
* 授权：
Authorization：base64(user:password)
* 路径中加冒号的为参数，需要传递相应的参数请求

***

## 登陆
```
POST /pili/v1/login
Content-Type: application/json

{
	"name":<name>,
	"password":<password>
}
```
* `<Name>`:用户名，必填
* `<Password>`:密码，必填

返回

```
{
    "code": 200,
    "message": "ok"
}
```

***

## 创建用户
```
POST /pili/v1/user/new
Content-Type: application/json
Authorization:<adminAuthorization>

{
	"name":<name>,
	"password":<password>,
	"room":<room>
}
```
* `<AdminAuthorization>`:admin鉴权，必填
* `<Name>`:用户名，必填
* `<Password>`:密码，必填
* `<Room>`:连麦房间，必填

返回

```
{
    "code": 200,
    "name": <name>,
    "room": <room>
}
```

***

## 查询用户
```
GET /pili/v1/user/query/:name

```

返回

```
{
    "code": 200,
    "name": <name>,
    "room": <room>
}
```

***

## 更新用户
```
POST /pili/v1/user/update/:name
Authorization:<authorization>
Content-Type: application/json

{
	"password":<password>,
}

```
* `<Password>`:修改后的密码

返回

```
{
    "code": 200,
    "message": "ok"
}
```

***

## 删除用户
```
POST /pili/v1/user/delete/:name
Authorization:<authorization>
Content-Type: application/json

```
**authorization中的用户和需要删除的name不匹配是无法删除的**

返回

```
{
    "code": 200,
    "message": "ok"
}
```

***

## 创建房间
```
POST /pili/v1/room/new
Authorization:<authorization>
Content-Type: application/json

{
	"room":<roomId>,
	"user":<user>,
	"max":<max>
}
```
* `<Room>`:房间名，必填
* `<User>`:房间所属用户，必填
* `<Max>`:允许连麦的最大人数，如果不填，默认为99

返回

```
{
	"code":200,
    "room": <room>
}
```

***

## 查询房间状态
```
GET /pili/v1/room/query/:id

```
返回

```
{
	"code":200,
	"room":<room>,
	"ownerId":<ownerId>,
	"userMax":<userMax>,
	"status":<status>,
}
```

***

## 删除房间
```
POST /pili/v1/room/delete/:id
Authorization:<authorization>
Content-Type: application/json

```
 返回

```
{
	"code":200,
	"message":"ok",
}
```

***

## 获取RoomToken
```
POST /pili/v1/room/token
Authorization:<authorization>
Content-Type: application/json

{
	"room":<room>,
	"user":<user>,
	"version":<version>
}
```
* `<Room>`:房间名，必填
* `<User>`:房间拥有人，必填
* `<Version>`:连麦版本号，字符串类型，默认为空，可填1.0或2.0

返回

```
<accesskey>:<sign>:<encodedSign>
```

***

## 获取推流地址
```
POST /pili/v1/stream/:id
Authorization:<authorization>

```

返回

```
{
    "code": 200,
    "url": <rtmp_publish_url>
}
```

***

## 获取播放地址
```
GET /pili/v1/stream/query/:id
Authorization:<authorization>

```
返回

```
{
    "code": 200,
    "hdl": <hdl_play_url>,
    "hls": <hls_play_url>,
    "rtmp": <rtmp_play_url>
}
```

***

## 错误码
* 403 : no authorized


