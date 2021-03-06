// Code generated by goctl. DO NOT EDIT.
package types

type IndexReply struct {
	Resp string `json:"resp"`
}

type GetCodeReq struct {
	Transport string `path:"transport"`
	Number    string `path:"number"`
	Client    string `form:"client,optional"`
	Captcha   string `form:"captcha,optional"`
	Challenge string `form:"challenge,optional"`
}

type DeviceCapabilities struct {
	UUID bool `json:"uuid,optional"`
}

type AccountAttributes struct {
	SignalingKey                   string             `json:"signalingKey,optional"`
	FetchesMessages                bool               `json:"fetchesMessages"`
	RegistrationID                 int                `json:"registrationId"`
	Name                           string             `json:"name,optional"`
	Pin                            string             `json:"pin,optional"`
	RegistrationLock               string             `json:"registrationLock,optional"`
	UnidentifiedAccessKey          string             `json:"unidentifiedAccessKey,optional"`
	UnrestrictedUnidentifiedAccess bool               `json:"unrestrictedUnidentifiedAccess,optional"`
	Capabilities                   DeviceCapabilities `json:"capabilities,optional"`
}

type VerifyAccountReq struct {
	VerificationCode               string             `path:"verificationCode"`
	SignalingKey                   string             `json:"signalingKey,optional"`
	FetchesMessages                bool               `json:"fetchesMessages"`
	RegistrationID                 int                `json:"registrationId"`
	Name                           string             `json:"name,optional"`
	Pin                            string             `json:"pin,optional"`
	RegistrationLock               string             `json:"registrationLock,optional"`
	UnidentifiedAccessKey          string             `json:"unidentifiedAccessKey,optional"`
	UnrestrictedUnidentifiedAccess bool               `json:"unrestrictedUnidentifiedAccess,optional"`
	Capabilities                   DeviceCapabilities `json:"capabilities,optional"`
}

type VerifyAccountRes struct {
	UUID string `json:"uuid"`
}

type SetAttributesReq struct {
	SignalingKey                   string             `json:"signalingKey,optional"`
	FetchesMessages                bool               `json:"fetchesMessages"`
	RegistrationID                 int                `json:"registrationId"`
	Name                           string             `json:"name,optional"`
	Pin                            string             `json:"pin,optional"`
	RegistrationLock               string             `json:"registrationLock,optional"`
	UnidentifiedAccessKey          string             `json:"unidentifiedAccessKey,optional"`
	UnrestrictedUnidentifiedAccess bool               `json:"unrestrictedUnidentifiedAccess,optional"`
	Capabilities                   DeviceCapabilities `json:"capabilities,optional"`
}

type DeprecatedPin struct {
	Pin string `json:"pin"`
}

type RegistrationLock struct {
	RegistrationLock string `json:"registrationLock"`
}

type DeviceName struct {
	DeviceName string `json:"deviceName"`
}

type AccountCreationResult struct {
	UUID string `json:"uuid"`
}

type GcmRegistrationID struct {
	GcmRegistrationID string `json:"gcmRegistrationId"`
}

type SetUserName struct {
	Username string `json:"username"`
}

type JwtTokenAdx struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}

type AdxUserLoginReq struct {
	Account                        string              `json:"account"` //eos chain username,保证unique
	Sign                           string              `json:"sign"`    //  eos 用户用自己的私钥对name的签名
	SignalingKey                   string              `json:"signalingKey,optional"`
	FetchesMessages                bool                `json:"fetchesMessages,default=true"`
	RegistrationID                 int                 `json:"registrationId,optional"`
	Pin                            string              `json:"pin,optional"`
	Name                           string              `json:"name,optional"`
	RegistrationLock               string              `json:"registrationLock,optional"`
	UnidentifiedAccessKey          string              `json:"unidentifiedAccessKey,optional"`
	UnrestrictedUnidentifiedAccess bool                `json:"unrestrictedUnidentifiedAccess,optional"`
	Capabilities                   DeviceCapabilitiesx `json:"capabilities,optional"`
}

type DeviceCapabilitiesx struct {
	UUID bool `json:"uuid,optional"`
}

type AdxUserLoginRes struct {
	JwtTokenAdx
	Uuid  string `json:"uuid"`
	IsNew bool   `json:"isNew"`
}

type IncomingMessagex struct {
	Content                   string `json:"content"`
	Type                      int    `json:"type"`
	DestinationDeviceId       int    `json:"destinationDeviceId,default=1"` //发到哪一个设备
	DestinationRegistrationId int    `json:"destinationRegistrationId"`
	Destination               string `json:"destination,optional"`
	Body                      string `json:"body,optional"`
	Relay                     string `json:"relay,optional"`
}

type PutMessagesReq struct {
	Destination string             `path:"destination" json:"destination"`
	Online      bool               `json:"online"`
	Timestamp   int64              `json:"timestamp"`
	Messages    []IncomingMessagex `json:"messages"`
}

type PutMessagesRes struct {
	NeedsSync bool `json:"needsSync"`
}

type OutcomingMessagex struct {
	Id              int64  `json:"id"`
	Cached          bool   `json:"cached"`
	Guid            string `json:"guid"`
	Type            int    `json:"type"`
	Relay           string `json:"relay"`
	Timestamp       int64  `json:"timestamp"`
	Source          string `json:"source"`
	SourceUuid      string `json:"sourceUuid"`
	SourceDevice    int64  `json:"sourceDevice"`
	Message         string `json:"message"`
	Content         string `json:"content"`
	ServerTimestamp int64  `json:"serverTimestamp"`
}

type GetPendingMsgsReq struct {
}

type GetPendingMsgsRes struct {
	List []OutcomingMessagex `json:"messages"`
	More bool                `json:"more"`
}

type Envelope struct {
	Xtype           int    `json:"type"`
	Source          string `json:"source"`
	SourceUuid      string `json:"sourceUuid"`
	SourceDevice    int    `json:"sourceDevice"`
	Relay           string `json:"relay"`
	Timestamp       uint64 `json:"timestamp"`
	LegacyMessage   string `json:"legacyMessage"`
	Content         string `json:"content"`
	ServerGuid      string `json:"guid"`
	ServerTimestamp uint64 `json:"serverTimestamp"`
}

type PubsubMessage struct {
	Xtype   int      `json:"type"`
	Content Envelope `json:"envelop"`
}

type PutProfileKeyReq struct {
	AccountName string `path:"accountName"`
	Profilekey  string `json:"profileKey"`
}

type GetProfileKeyReq struct {
	AccountName string `path:"accountName"`
}

type GetProfileKeyRes struct {
	Profilekey  string `json:"profileKey"`
	Uuid        string `json:"uuid"`
	AccountName string `path:"accountName"`
}

type DeliveryReq struct {
}

type DeliveryRes struct {
	Certificate string `json:"certificate"`
}

type PreKey struct {
	KeyId     int64  `json:"keyId"`
	PublicKey string `json:"publicKey"`
}

type SignedPrekey struct {
	Signature string `json:"signature"`
	KeyId     int64  `json:"keyId"`
	PublicKey string `json:"publicKey"`
}

type PutKeysReq struct {
	IdentityKey  string       `json:"identityKey"`
	SignedPreKey SignedPrekey `json:"signedPreKey"`
	PreKeys      []PreKey     `json:"preKeys"`
}

type GetKeysReqx struct {
	Identifier string `path:"identifier"`
	DeviceId   string `path:"deviceId"`
}

type PreKeyResponseItemx struct {
	DeviceId       int64        `json:"deviceId"`
	RegistrationId int64        `json:"registrationId"`
	PreKey         PreKey       `json:"preKey"`
	SignedPrekey   SignedPrekey `json:"signedPreKey"`
}

type GetKeysRes struct {
	IdentityKey string                `json:"identityKey"`
	Devices     []PreKeyResponseItemx `json:"devices"`
}

type PreKeyCountx struct {
	Count int `json:"count"`
}

type SetNameReq struct {
	Name string `path:"name"`
}

type CreateProfileRequest struct {
	Version    string `json:"version"`
	Name       string `json:"name"`
	Avatar     bool   `json:"avatar"`
	Commitment string `json:"commitment"`
}

type ProfileAvatarUploadAttributes struct {
	Key        string `json:"key"`
	Credential string `json:"credential"`
	Acl        string `json:"acl"`
	Algorithm  string `json:"algorithm"`
	Date       string `json:"date"`
	Policy     string `json:"policy"`
	Signature  string `json:"signature"`
}

type GetProfileByUserName struct {
	UserName string `path:"username"`
}

type UserCapabilities struct {
	Uuid bool `json:"uuid"`
	Gv2  bool `json:"gv2"`
}

type Profile struct {
	IdentityKey                    string           `json:"identityKey"`
	Name                           string           `json:"name"`
	Avatar                         string           `json:"avatar"`
	UnidentifiedAccess             string           `json:"unidentifiedAccess"`
	UnrestrictedUnidentifiedAccess bool             `json:"unrestrictedUnidentifiedAccess"`
	Capabilities                   UserCapabilities `json:"capabilities"`
	UserName                       string           `json:"userName"`
	Uuid                           string           `json:"uuid"`
	Credential                     string           `json:"credential"`
}

type GetProfileByUUID struct {
	Uuid    string `path:"uuid"`
	Version string `path:"version"`
}

type GetProfileByUUIDCredentia struct {
	Uuid              string `path:"uuid"`
	Version           string `path:"version"`
	CredentialRequest string `path:"credentialRequest"`
}

type ProvisioningMessage struct {
	Destination string `path:"destination"`
	Body        string `json:"body"`
}

type ExternalServiceCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"` // UserName:UserName加盐哈系
}

type ClientContact struct {
	Token    string `json:"token"`
	Voice    bool   `json:"voice,optional"`
	Video    bool   `json:"video,optional"`
	Relay    string `json:"relay,optional"`
	Inactive bool   `json:"inactive,optional"`
}

type ClientContactTokens struct {
	Contacts []string `json:"contacts"`
}

type ClientContacts struct {
	Contacts []ClientContact `json:"contacts"`
}

type GetTokenPresenceReq struct {
	Token string `path:"token"`
}

type DeviceInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastSeen int64  `json:"lastSeen"`
	Created  int64  `json:"created"`
}

type DeviceInfoList struct {
	Devices []DeviceInfo `json:"devices"`
}

type DelDeviceReq struct {
	DeviceId string `path:"device_id"`
}

type VerificationCode struct {
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"`
	PushCode  string `json:"pushCode,optional"`
}
