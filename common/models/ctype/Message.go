package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// MessageType 聊天信息类型
type MessageType struct {
	// 消息类型(0: 文本类型，1: 图片消息，2: 视频消息，3: 文件消息，4: 语音消息，5: 语言通话，6: 视频通话，7: 撤回消息，8: 回复消息，9: 引用消息)
	Type int8 `gorm:"column:type; comment:'消息类型'" json:"type"`
	// 文本消息
	TextMessage *TextMessage `gorm:"column:text_message;comment:'文本消息'" json:"text_message"`
	// 图片消息
	ImageMessage *ImageMessage `gorm:"column:image_message;comment:'图片消息'" json:"image_message"`
	// 视频消息
	VideoMessage *VideoMessage `gorm:"column:video_message;comment:'视频消息'" json:"video_message"`
	// 文件消息
	FileMessage *FileMessage `gorm:"column:file_message;comment:'文件消息'" json:"file_message"`
	// 语音消息
	AudioMessage *AudioMessage `gorm:"column:audio_message;comment:'语音消息'" json:"audio_message"`
	// 语言通话
	VoiceCall *VoiceCall `gorm:"column:voice_call;comment:'语言通话'" json:"voice_call"`
	// 视频通话
	VideoCall *VideoCall `gorm:"column:video_call;comment:'视频通话'" json:"video_call"`
	// 撤回消息
	RecallMessage *RecallMessage `gorm:"column:recall_message;comment:'撤回消息'" json:"recall_message"`
	// 回复消息
	ReplyMessage *ReplyMessage `gorm:"column:reply_message;comment:'回复消息'" json:"reply_message"`
	// 引用消息
	QuoteMessage *QuoteMessage `gorm:"column:quote_message;comment:'引用消息'" json:"quote_message"`
	// @消息,群聊才有
	AtMessage *AtMessage `gorm:"column:at_message;comment:'@消息'" json:"at_message"`
}

// Scan 从数据库中检索数据时使用Scan方法
func (c *MessageType) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 将数据准备好以存储到数据库中
func (c MessageType) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

// TextMessage 是文本消息
type TextMessage struct {
	Context string `gorm:"column:content; comment:'消息内容'" json:"content"`
}

// ImageMessage 是图片消息
type ImageMessage struct {
	Title    string `gorm:"column:title; comment:'图片标题'" json:"title"`
	ImageUrl string `gorm:"column:image_url; comment:'图片链接'" json:"image_url"`
}

// VideoMessage 是视频消息
type VideoMessage struct {
	Title         string `gorm:"column:title; comment:'视频标题'" json:"title"`
	VideoUrl      string `gorm:"column:video_url; comment:'视频链接'" json:"video_url"`
	VideoDuration int32  `gorm:"column:video_duration; comment:'视频时长'" json:"video_duration"`
}

// FileMessage 是文件消息
type FileMessage struct {
	FileName string `gorm:"column:file_name; comment:'文件名'" json:"file_name"`
	FileUrl  string `gorm:"column:file_url; comment:'文件链接'" json:"file_url"`
	FileSize int64  `gorm:"column:file_size; comment:'文件大小'" json:"file_size"`
	FileType string `gorm:"column:file_type; comment:'文件类型'" json:"file_type"`
}

// AudioMessage 是音频消息
type AudioMessage struct {
	AudioUrl string `gorm:"column:audio_url; comment:'音频链接'" json:"audio_url"`
	Duration int64  `gorm:"column:duration; comment:'音频时长'" json:"duration"`
}

// VoiceCall 是语音通话
type VoiceCall struct {
	CallDuration int32     `gorm:"column:call_duration; comment:'通话时长'" json:"call_duration"`
	StartTime    time.Time `gorm:"column:start_time; comment:'通话开始时间'" json:"start_time"`
	EndTime      time.Time `gorm:"column:end_time; comment:'通话结束时间'" json:"end_time"`
	EndReason    int8      `gorm:"column:end_reason;not null;comment:'通话结束原因(0:正常挂断, 1:对方已拒接, 2:对方不在线, 3:网络不可达, 4:对方已取消, 5:未知错误)'"; json:"end_reason"`
}

// VideoCall 是视频通话
type VideoCall struct {
	CallDuration int32     `gorm:"column:call_duration; comment:'通话时长'" json:"call_duration"`
	StartTime    time.Time `gorm:"column:start_time; comment:'通话开始时间'" json:"start_time"`
	EndTime      time.Time `gorm:"column:end_time; comment:'通话结束时间'" json:"end_time"`
}

// RecallMessage 是撤回消息
type RecallMessage struct {
	OriginMessageID    int8   `gorm:"column:origin_message_id; comment:'原始消息的ID';not null" json:"origin_message_id"`
	OriginMessage      string `gorm:"column:origin_message; comment:'原始消息的内容';not null" json:"origin_message"`
	WithdrawalReminder string `gorm:"column:withdrawal_reminder; comment:'撤回消息的提示';not null" json:"withdrawal_reminder"`
}

// ReplyMessage 是回复消息
type ReplyMessage struct {
	ReplyToMessageID int8        `gorm:"column:reply_to_message_id; comment:'回复的消息ID';not null" json:"reply_to_message_id"`
	ReplyComment     string      `gorm:"column:reply_comment; comment:'回复的消息内容';not null" json:"reply_comment"`
	ReplyMessage     MessageType `gorm:"column:reply_message; comment:'回复的消息类型';not null" json:"reply_message"`
}

// QuoteMessage 是引用消息
type QuoteMessage struct {
	QuoteID      int8        `gorm:"column:quote_id; comment:'引用的消息 ID';not null" json:"quote_id"`
	QuoteComment string      `gorm:"column:quote_comment; comment:'引用的消息内容';not null" json:"quote_comment"`
	QuoteMessage MessageType `gorm:"column:quote_message; comment:'引用的消息类型';not null" json:"quote_message"`
}

// AtMessage 是@消息
type AtMessage struct {
	AtUserID  int8        `gorm:"column:at_user_id; comment:'@的用户ID';not null" json:"at_user_id"`
	AtComment string      `gorm:"column:at_comment; comment:'@的消息内容';not null" json:"at_comment"`
	AtMessage MessageType `gorm:"column:at_message; comment:'@的消息类型';not null" json:"at_message"`
}
