package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetPostDetailByPid(pid int) (*models.PostRes, error) {
	var post = &models.PostRes{}
	p, err := dao.GetPostDetailByPid(pid)
	if err != nil {
		return nil, err
	}
	post.Viewer = config.Cfg.Viewer
	post.SystemConfig = config.Cfg.SystemConfig
	post.Article = models.PostMore{
		Pid:          p.Pid,
		Title:        p.Title,
		Slug:         p.Slug,
		Content:      template.HTML(p.Content),
		CategoryId:   p.CategoryId,
		CategoryName: dao.GetCategoryNameByCategoryId(p.CategoryId),
		UserId:       p.UserId,
		UserName:     dao.GetUserNameByUserId(p.UserId),
		ViewCount:    p.ViewCount,
		Type:         p.Type,
		CreateAt:     models.DateDay(p.CreateAt),
		UpdateAt:     models.DateDay(p.UpdateAt),
	}

	return post, nil
}

func GetWriting() (writeing *models.WriteingResponse, err error) {
	writeing = &models.WriteingResponse{}
	writeing.Title = config.Cfg.Viewer.Title
	writeing.CdnURL = config.Cfg.SystemConfig.CdnURL
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	writeing.Categorys = categorys
	return writeing, nil
}

func GetPigeonhole() (p *models.PigeonholeRes, err error) {
	p = &models.PigeonholeRes{}
	p.Viewer = config.Cfg.Viewer
	p.SystemConfig = config.Cfg.SystemConfig
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	p.Categorys = categorys
	posts, err := dao.GetPostAll()
	if err != nil {
		return nil, err
	}

	// 创建一个map实例
	linesMap := make(map[string][]models.Post)
	for _, post := range posts {
		month := post.CreateAt.Format("2006-01-02")
		linesMap[month] = append(linesMap[month], post)
	}
	p.Lines = linesMap

	return p, nil
}

func SearchPost(val string) (resp []models.SearchResp, err error) {
	resp, err = dao.SearchPost(val)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
