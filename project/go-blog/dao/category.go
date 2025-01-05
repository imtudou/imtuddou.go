package dao

import (
	"go-blog/models"
)

// 获取类型
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		return nil, err
	}

	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		if err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt); err != nil {
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

func GetCategoryNameByCategoryId(categoryId int) (name string) {
	row := DB.QueryRow("select name  from blog_category where cid =?", categoryId)
	if row.Err() != nil {
		return ""
	}
	err := row.Scan(&name)
	if err != nil {
		return ""
	}
	return name
}
