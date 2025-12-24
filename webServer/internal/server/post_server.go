package gorm

import (
	"fmt"
	"go-agriculture/internal/config"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"go-agriculture/internal/util"
	"log"
	"strconv"
	"time"
)

type postServerType struct{}

var PostServer = &postServerType{}

// GetPostList 获取帖子列表（对应前端 /posts）
func (p *postServerType) GetPostList() (string, []respond.Post, int) {
	var posts []model.Post

	// 查询数据，按创建时间倒序
	if res := dao.GormDB.Order("created_at DESC").Find(&posts); res.Error != nil {
		log.Printf("查询帖子列表失败: %v", res.Error)
		return "查询失败", nil, -1
	}

	// 转换数据
	postList := make([]respond.Post, 0, len(posts))
	for _, post := range posts {
		newPost := p.convertPostToRespond(post, "")
		postList = append(postList, newPost)
	}

	return "获取成功", postList, 0
}

// GetPostById 获取单个帖子详情（对应前端 /posts/:id）
func (p *postServerType) GetPostById(id string, userId string) (string, *respond.Post, int) {
	var post model.Post

	// 查询帖子
	if res := dao.GormDB.Where("uuid = ?", id).First(&post); res.Error != nil {
		log.Printf("查询帖子失败: %v", res.Error)
		return "帖子不存在", nil, -1
	}

	postData := p.convertPostToRespond(post, userId)
	return "获取成功", &postData, 0
}

// CreatePost 创建帖子（对应前端 POST /posts）
func (p *postServerType) CreatePost(req request.CreatePostRequest, userId string) (string, *respond.Post, int) {
	// 验证用户ID是否匹配
	if req.Author.Id != userId {
		return "用户信息不匹配", nil, -1
	}

	// 检查用户是否存在
	var user model.User
	if res := dao.GormDB.Where("uuid = ?", userId).First(&user); res.Error != nil {
		log.Printf("查询用户失败: %v", res.Error)
		return "用户不存在", nil, -1
	}

	tempAvatar := "http://" + fmt.Sprint(config.GetConfig().MainConfig.ServerIP) + ":" + fmt.Sprint(config.GetConfig().MainConfig.Port) + "/static/avatars/" + user.Avatar

	post := model.Post{
		Uuid:         util.GenerateUUID(),
		Title:        req.Title,
		Content:      req.Content,
		AuthorId:     userId,
		AuthorName:   req.Author.Nickname,
		AuthorAvatar: tempAvatar,
		LikeCount:    0,
		ReplyCount:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if res := dao.GormDB.Create(&post); res.Error != nil {
		log.Printf("创建帖子失败: %v", res.Error)
		return "创建帖子失败", nil, -1
	}

	postData := p.convertPostToRespond(post, userId)
	return "创建成功", &postData, 0
}

// AddReply 添加回复（对应前端 POST /posts/:postId/replies）
func (p *postServerType) AddReply(postId string, req request.CreateReplyRequest, userId string) (string, *respond.Reply, int) {
	// 检查帖子是否存在
	var post model.Post
	if res := dao.GormDB.Where("uuid = ?", postId).First(&post); res.Error != nil {
		log.Printf("查询帖子失败: %v", res.Error)
		return "帖子不存在", nil, -1
	}

	var user model.User

	if res := dao.GormDB.Where("uuid = ?", userId).First(&user); res.Error != nil {

		log.Printf("查询用户失败: %v", res.Error)

		return "用户不存在", nil, -1

	}
	tempAvatar := "http://" + fmt.Sprint(config.GetConfig().MainConfig.ServerIP) + ":" + fmt.Sprint(config.GetConfig().MainConfig.Port) + "/static/avatars/" + user.Avatar

	// 创建回复
	reply := model.PostReply{
		PostId:       postId,
		Content:      req.Content,
		AuthorId:     userId,
		AuthorName:   req.Author.Nickname,
		AuthorAvatar: tempAvatar,
		LikeCount:    0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if res := dao.GormDB.Create(&reply); res.Error != nil {
		log.Printf("创建回复失败: %v", res.Error)
		return "创建回复失败", nil, -1
	}

	// 更新帖子的回复数
	post.ReplyCount += 1
	dao.GormDB.Save(&post)

	replyData := p.convertReplyToRespond(reply, userId)
	return "回复成功", &replyData, 0
}

// TogglePostLike 切换帖子点赞状态（对应前端 POST /posts/:postId/like）
func (p *postServerType) TogglePostLike(postId string, req request.ToggleLikeRequest, userId string) (string, int) {
	// 检查帖子是否存在
	var post model.Post
	if res := dao.GormDB.Where("uuid = ?", postId).First(&post); res.Error != nil {
		log.Printf("查询帖子失败: %v", res.Error)
		return "帖子不存在", -1
	}

	// 检查当前点赞状态
	var count int64
	dao.GormDB.Model(&model.PostLike{}).
		Where("post_id = ? AND user_id = ?", postId, userId).
		Count(&count)
	currentLiked := count > 0

	if currentLiked == req.Liked {
		return "状态未改变", -1
	}

	if req.Liked {
		// 添加点赞
		like := model.PostLike{
			PostId:    postId,
			UserId:    userId,
			CreatedAt: time.Now(),
		}
		dao.GormDB.Create(&like)
		post.LikeCount += 1
	} else {
		// 取消点赞
		var like model.PostLike
		dao.GormDB.Where("post_id = ? AND user_id = ?", postId, userId).First(&like)
		dao.GormDB.Delete(&like)
		if post.LikeCount > 0 {
			post.LikeCount -= 1
		}
	}

	dao.GormDB.Save(&post)
	return "操作成功", 0
}

// ToggleReplyLike 切换回复点赞状态（对应前端 POST /posts/:postId/replies/:replyId/like）
func (p *postServerType) ToggleReplyLike(postId string, replyId string, req request.ToggleLikeRequest, userId string) (string, int) {
	// 将回复ID转换为int64
	replyIdInt, err := parseReplyId(replyId)
	if err != nil {
		return "回复ID格式错误", -1
	}

	// 检查回复是否存在
	var reply model.PostReply
	if res := dao.GormDB.Where("id = ? AND post_id = ?", replyIdInt, postId).First(&reply); res.Error != nil {
		log.Printf("查询回复失败: %v", res.Error)
		return "回复不存在", -1
	}

	// 检查当前点赞状态
	var count int64
	dao.GormDB.Model(&model.PostReplyLike{}).
		Where("reply_id = ? AND user_id = ?", replyIdInt, userId).
		Count(&count)
	currentLiked := count > 0

	if currentLiked == req.Liked {
		return "状态未改变", -1
	}

	if req.Liked {
		// 添加点赞
		like := model.PostReplyLike{
			ReplyId:   replyIdInt,
			UserId:    userId,
			CreatedAt: time.Now(),
		}
		dao.GormDB.Create(&like)
		reply.LikeCount += 1
	} else {
		// 取消点赞
		var like model.PostReplyLike
		dao.GormDB.Where("reply_id = ? AND user_id = ?", replyIdInt, userId).First(&like)
		dao.GormDB.Delete(&like)
		if reply.LikeCount > 0 {
			reply.LikeCount -= 1
		}
	}

	dao.GormDB.Save(&reply)
	return "操作成功", 0
}

// convertPostToRespond 将Post model转换为Post respond
func (p *postServerType) convertPostToRespond(post model.Post, userId string) respond.Post {
	// 处理头像路径
	avatar := post.AuthorAvatar

	// 查询当前用户是否点赞该帖子
	var liked bool
	if userId != "" {
		var count int64
		dao.GormDB.Model(&model.PostLike{}).
			Where("post_id = ? AND user_id = ?", post.Uuid, userId).
			Count(&count)
		liked = count > 0
	}

	// 查询回复列表
	var replies []model.PostReply
	dao.GormDB.Where("post_id = ?", post.Uuid).
		Order("created_at ASC").
		Find(&replies)

	// 转换回复数据
	replyList := make([]respond.Reply, 0, len(replies))
	for _, reply := range replies {
		replyList = append(replyList, p.convertReplyToRespond(reply, userId))
	}

	return respond.Post{
		Id:      post.Uuid,
		Title:   post.Title,
		Content: post.Content,
		Author: respond.Author{
			Id:       post.AuthorId,
			Nickname: post.AuthorName,
			Avatar:   avatar,
		},
		Likes:     post.LikeCount,
		Liked:     liked,
		Replies:   replyList,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	}
}

// convertReplyToRespond 将PostReply model转换为Reply respond
func (p *postServerType) convertReplyToRespond(reply model.PostReply, userId string) respond.Reply {
	// 处理头像路径
	avatar := reply.AuthorAvatar

	// 查询当前用户是否点赞该回复
	var liked bool
	if userId != "" {
		var count int64
		dao.GormDB.Model(&model.PostReplyLike{}).
			Where("reply_id = ? AND user_id = ?", reply.Id, userId).
			Count(&count)
		liked = count > 0
	}

	return respond.Reply{
		Id:      fmt.Sprintf("%d", reply.Id),
		Content: reply.Content,
		Author: respond.Author{
			Id:       reply.AuthorId,
			Nickname: reply.AuthorName,
			Avatar:   avatar,
		},
		Likes:     reply.LikeCount,
		Liked:     liked,
		CreatedAt: reply.CreatedAt.Format(time.RFC3339),
	}
}

// parseReplyId 解析回复ID（支持数字和字符串格式）
func parseReplyId(replyId string) (int64, error) {
	// 前端可能发送 "reply_xxx" 格式的ID，需要提取数字部分
	return strconv.ParseInt(replyId, 10, 64)
}
