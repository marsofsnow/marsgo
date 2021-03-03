// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type IndexReply struct {
	Resp string `json:"resp"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserReply struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	JwtToken
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}

type AddReq struct {
	Book  string `form:"book"`
	Price int64  `form:"price"`
}

type AddResp struct {
	Ok bool `json:"ok"`
}

type CheckReq struct {
	Book string `form:"book"`
}

type CheckResp struct {
	Found bool  `json:"found"`
	Price int64 `json:"price"`
}

type GetSmsCodeReq struct {
	Number string `path:"number"` // 路由path，如/foo/:id
}

type GetSmsCodeResp struct {
	SmsCode string `json:"smsCode"`
}

type AccountAttributes struct {
	SignalingKey    string `json:"signalingKey"` // signalingKey is a randomly generated 32 byte AES key and a 20 byte HMAC-SHA1 MAC key, concatenated together and Base64 encoded.
	FetchesMessages bool   `json:"fetchesMessages default=true"`
	RegistrationId  uint64 `json:"registrationId"` //// registrationId is a 14 bit integer that's randomly generated at client install time. This will be used for clients to detect whether an app has reinstalled and lost their session state.
	Name            string `json:"name"`
	Voice           bool   `json:"voice"`
	Video           bool   `json:"video"`
}

type SimpleDevice struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	LastSeen int64  `json:"lastSeen"`
	Created  int64  `json:"created"`
}

type PreKey struct {
	KeyId     int64  `json:"keyId"`
	PublicKey string `json:"publicKey"`
}

type SignedPreKey struct {
	Prekey    PreKey `json:"prekey"`
	Signature string `json:"signature"`
}

type ConfirmVerificationCodeReq struct {
	VerificationCode  string            `path:"verificationCode"`
	PhoneNumber       string            `json:"phoneNumber"`
	Password          string            `json:"password"`
	SupportsSms       bool              `json:"supportsSms"` // supportsSms indicates whether a client supports SMS as a transport.
	AccountAttributes AccountAttributes `json:"accountAttributes"`
}

type ConfirmVerificationCodeRes struct {
	Account string `json:"account"`
}

type PutAccountAttributesReq struct {
	PhoneNumber       string            `json:"phoneNumber"`
	AccountAttributes AccountAttributes `json:"accountAttributes"`
}

type PutAccountAttributesRes struct {
	Account string `json:"account"`
}

type GetDevicesRes struct {
	Total int            `json:"total"`
	List  []SimpleDevice `json:"list"`
}

type DelDeviceReq struct {
	DeviceId int64 `path:"deviceId"`
}

type DelDeviceRes struct {
	Device SimpleDevice `json:"device"`
}

type PutDeviceReq struct {
	VerificationCode  string            `path:"verificationCode"`
	PhoneNumber       string            `json:"phoneNumber"`
	AccountAttributes AccountAttributes `json:"accountAttributes"`
}

type PutDeviceRes struct {
	Account string `json:"account"`
}

type GetProfileReq struct {
	Number string `path:"number"` // 路由path，如/foo/:id
}

type GetProfileResp struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	IdentityKey string `json:"identityKey"`
}

type GetDirTokenReq struct {
	Token string `path:"token"`
}

type GetDirTokenRes struct {
	Token string `json:"token"`
	Relay string `json:"relay,omitempty"`
	Voice bool   `json:"voice"`
	Video bool   `json:"video"`
}

type GetDirTokensReq struct {
	Tokens []string `json:"tokens"`
}

type GetDirTokensRes struct {
	Total int              `json:"total"`
	List  []GetDirTokenRes `json:"list"`
}

type GetKeysRes struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"` //返回数据库表元素
}

type KeyDevices struct {
	DeviceId       int64        `json:"deviceId"`
	RegistrationId uint64       `json:"registrationId"`
	SignedPreKey   SignedPreKey `json:"signedPreKey"`
	PreKey         PreKey       `json:"preKey"`
}

type GetDeviceKeyReq struct {
	Number   string `json:"number"`
	DeviceId int    `json:"deviceId"`
}

type GetDeviceKeyRes struct {
	IdentityKey string       `json:"identityKey"`
	Devices     []KeyDevices `json:"devices"`
}

type PutKeysReq struct {
	PreKeys      []PreKey     `json:"preKeys"`
	IdentityKey  string       `json:"identityKey"`
	SignedPreKey SignedPreKey `json:"signedPreKey"`
}

type PutKeysResp struct {
	Resp interface{} `json:"resp"`
}

type GetKeysSignedRes struct {
	SignedPreKey SignedPreKey `json:"signedPreKey"`
}

type PutKeysSignedReq struct {
	SignedPreKey SignedPreKey `json:"signedPreKey"`
}

type PutKeysSignedRes struct {
	Resp interface{} `json:"resp"`
}
