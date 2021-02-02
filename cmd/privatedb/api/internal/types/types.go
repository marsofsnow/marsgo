// Code generated by goctl. DO NOT EDIT.
package types

type AddmsgReply struct {
	Ok bool `json:"ok"` //是否成功
}

type AddmsgReq struct {
	Pair     string `json:"pair"`     // 交易对
	Dealid   int64  `json:"dealid"`   //订单id
	Sender   string `json:"sender"`   //发送者
	Receiver string `json:"receiver"` //接收者
	Content  string `json:"content"`  //消息内容
}

type GetPayAccountReply struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

type GetPayAccountReq struct {
	Username    string `json:"username,optional"`   //用户名
	PaymentType uint32 `json:"pmt_type,optional"`   //支付方式id
	IsActived   uint8  `json:"is_actived,optional"` //是否激活
}

type GetPaymentTypeReply struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

type GetmsgsReply struct {
	Total    int         `json:"total"`
	PageSize int         `json:"pageSize"`
	Data     interface{} `json:"list"`
}

type GetmsgsReq struct {
	PageIndex int    `json:"pageIndex"` //页数
	PageSize  int    `json:"pageSize"`  //叶大小
	Pair      string `json:"pair"`
	Dealid    int64  `json:"dealid"`
	Status    int    `json:"status,optional"` //1.未读2.已读 选填，没有查所有
}

type IndexReply struct {
	Resp string `json:"resp"`
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type UpsertPayAccountReply struct {
	Ok bool `json:"ok"` //是否成功
}

type UpsertPayAccountReq struct {
	Id            uint32 `json:"pmt_id"`                 //=0插入 ！=0 update
	Username      string `json:"username"`               //用户名
	PaymentTypeId uint32 `json:"pmt_type"`               //支付方式id
	IsActived     uint8  `json:"is_actived,options=1|2"` //是否激活
	Detail        string `json:"ciphertext"`             //帐号详情
}

type UserReply struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	JwtToken
}