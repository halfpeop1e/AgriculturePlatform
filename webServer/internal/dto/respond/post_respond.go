package respond

// Author 作者信息（匹配前端 Author 接口）
type Author struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar,omitempty"`
}

// Reply 回复信息（匹配前端 Reply 接口）
type Reply struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	Author    Author `json:"author"`
	Likes     int    `json:"likes,omitempty"`
	Liked     bool   `json:"liked,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

// Post 帖子信息（匹配前端 Post 接口）
type Post struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Author    Author  `json:"author"`
	Likes     int     `json:"likes,omitempty"`
	Liked     bool    `json:"liked,omitempty"`
	Replies   []Reply `json:"replies,omitempty"`
	CreatedAt string  `json:"createdAt,omitempty"`
}
