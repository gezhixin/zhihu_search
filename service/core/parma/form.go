package parma

const (
	ExclusivityLoginErrorCode = 9999999
)

type ClientInfo struct {
	DeviceId      string `json:"devId" form:"devId"`
	UserId        string `json:"uid" form:"uid"`
	DeviceName    string `json:"devName" form:"devName"`
	Plat          string `json:"plat" form:"plat"`
	ClientVersion string `json:"cv" form:"cv"`
	AppSource     string `json:"src" form:"src"`
	SessionId     string `json:"sid" form:"sid"`
}
