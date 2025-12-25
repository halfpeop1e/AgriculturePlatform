package gorm

import (
	"encoding/json"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
)

type knowledgeServerType struct{}

var KnowledgeServer = &knowledgeServerType{}

func (k *knowledgeServerType) GetKnowledgeList() (string, *[]respond.KnowledgeRespond, int) {
	var knowledges []respond.KnowledgeRespond
	// 查询所有知识
	var knowledgeModels []model.Knowledge
	if res := dao.GormDB.Find(&knowledgeModels); res.Error != nil {
		return "查询失败", nil, -1
	}
	for _, knowledge := range knowledgeModels {
		var temptags []string
		json.Unmarshal([]byte(knowledge.Tags), &temptags)
		knowledges = append(knowledges, respond.KnowledgeRespond{
			Id:      knowledge.Id,
			Title:   knowledge.Title,
			Excerpt: knowledge.Excerpt,
			Content: knowledge.Content,
			Author:  knowledge.Author,
			Date:    knowledge.Date,
			Tags:    temptags,
		})
	}
	return "获取成功", &knowledges, 0
}

func (k *knowledgeServerType) AddKnowledge(req request.KnowledgeRequest) (string, int) {
	if len(req.Title) > 200 || len(req.Excerpt) > 200 || len(req.Content) > 2000 || len(req.Author) > 20 {
		return "参数错误", -1
	}
	tags, _ := json.Marshal(req.Tags)
	newKnowledge := model.Knowledge{
		Title:   req.Title,
		Excerpt: req.Excerpt,
		Content: req.Content,
		Tags:    string(tags),
		Author:  req.Author,
		Date:    req.Date,
	}

	if res := dao.GormDB.Create(&newKnowledge); res.Error != nil {
		return "添加失败", -1
	}

	return "添加成功", 0
}
