package event

import (
	"os/user"

	"github.com/dezhishen/satori-model-go/pkg/channel"
	"github.com/dezhishen/satori-model-go/pkg/guild"
	"github.com/dezhishen/satori-model-go/pkg/guildmember"
	"github.com/dezhishen/satori-model-go/pkg/guildrole"
	"github.com/dezhishen/satori-model-go/pkg/login"
	"github.com/dezhishen/satori-model-go/pkg/message"
)

type Op struct {
	Op uint8 `json:"op,omitempty"`
}

type EventOp struct {
	*Op
	Body *Event `json:"body"`
}

type Event struct {
	Id        int64                    `json:"id"`                 //事件 ID
	Type      string                   `json:"type"`               //事件类型
	Platform  string                   `json:"platform"`           //接收者的平台名称
	SelfId    string                   `json:"self_id"`            //接收者的平台账号
	Timestamp int64                    `json:"timestamp"`          //事件的时间戳
	Channel   *channel.Channel         `json:"channel,omitempty"`  //事件所属的频道
	Guild     *guild.Guild             `json:"guild,omitempty"`    //事件所属的群组
	Login     *login.Login             `json:"login,omitempty"`    //事件的登录信息
	Member    *guildmember.GuildMember `json:"member,omitempty"`   //事件的目标成员
	Message   *message.Message         `json:"message,omitempty"`  //事件的消息
	Operator  *user.User               `json:"operator,omitempty"` //事件的操作者
	Role      *guildrole.GuildRole     `json:"role,omitempty"`     //事件的目标角色
	User      *user.User               `json:"user,omitempty"`     //事件的目标用户
}

type PingOpInfo struct {
	*Op
}

type PongOpInfo struct {
	*Op
}

type IdentifyOpInfo struct {
	*Op
	Body *Identify `json:"body,omitempty"`
}

type Identify struct {
	Token    string `json:"token"`
	Sequence int64  `json:"sequence,omitempty"`
}

type ReadyOpInfo struct {
	*Op
	Body *Ready `json:"body"`
}

type Ready struct {
	Logins []login.Login `json:"logins,omitempty"`
}
