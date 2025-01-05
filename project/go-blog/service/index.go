package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetAllIndexInfo(slug string, pageIndex int, pageSize int) (*models.HomeResponse, error) {
	var posts []models.Post
	var postMore []models.PostMore
	var total int
	var err error

	category, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	if slug == "" {
		posts, total, err = dao.GetPostPage(pageIndex, pageSize)
		if err != nil {
			return nil, err
		}
	} else {
		posts, total, err = dao.GetPostPageBySlug(slug, pageIndex, pageSize)
		if err != nil {
			return nil, err
		}
	}

	for _, p := range posts {
		content := []rune(p.Content) // 表示中文字符
		if len(content) > 200 {
			content = content[:200] // 只取200个字符
		}

		pm := models.PostMore{
			Pid:          p.Pid,
			Title:        p.Title,
			Slug:         p.Slug,
			Content:      template.HTML(content), // 内容超长 折叠显示
			CategoryId:   p.CategoryId,
			CategoryName: category[p.CategoryId].Name,
			UserId:       p.UserId,
			UserName:     dao.GetUserNameByUserId(p.UserId),
			ViewCount:    p.ViewCount,
			Type:         p.Type,
			CreateAt:     models.DateDay(p.CreateAt),
			UpdateAt:     models.DateDay(p.UpdateAt),
		}
		postMore = append(postMore, pm)

	}

	// 页面上用到的所有数据需要定义
	// 计算总页数

	pageAll := ((total - 1) / pageSize) + 1
	cc := []int{}
	for i := 0; i < pageAll; i++ {
		cc = append(cc, i+1)
	}

	hr := &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: category,
		Posts:     postMore,
		Total:     total,
		Page:      pageIndex,
		Pages:     cc,
		PageEnd:   pageIndex != pageAll,
	}
	return hr, nil
}
