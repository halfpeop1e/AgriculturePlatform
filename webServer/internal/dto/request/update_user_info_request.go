package request

import "mime/multipart"

type UpdateUserInfoRequest struct {
	Nickname string                `form:"nickname" json:"nickname"`
	Bio      string                `form:"bio" json:"bio"`
	Avatar   *multipart.FileHeader `form:"avatar" json:"avatar"`
	Location string                `form:"location" json:"location"`
	Phone    string                `form:"phone" json:"phone"`
	Address  string                `form:"address" json:"address"`
	Tags     []string              `form:"tags" json:"tags"`
}
