package dao

import "go-blog/models"

func GetPostPage(pageIndex, pageSize int) (posts []models.Post, count int, err error) {

	rows, err := DB.Query("select  * from blog_post limit ?,?", (pageIndex-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		posts = append(posts, post)
	}

	row := DB.QueryRow("select count(1) from blog_post")
	if row.Err() != nil {
		return nil, 0, row.Err()
	}
	row.Scan(&count)

	return posts, count, nil
}

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select  * from blog_post")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostDetailByPid(pid int) (*models.Post, error) {
	row := DB.QueryRow("select  * from blog_post where pid =?", pid)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var post = &models.Post{}
	row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	return post, nil
}

func GetPostCount() (count int) {
	row := DB.QueryRow("select count(1) from blog_post")
	if row.Err() != nil {
		return 0
	}
	row.Scan(&count)
	return
}

func GetPostCountByCID(cid int) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where category_id=?", cid)
	if row.Err() != nil {
		return 0
	}
	row.Scan(&count)
	return
}

func GetPostPageByCategoryID(cid, pageIndex, pageSize int) ([]models.Post, error) {
	rows, err := DB.Query("select  * from blog_post where category_id = ? limit ?,?", cid, (pageIndex-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBySlug(slug string, pageIndex, pageSize int) (posts []models.Post, count int, err error) {
	rows, err := DB.Query("select  * from blog_post where slug = ? limit ?,?", slug, (pageIndex-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		posts = append(posts, post)
	}

	row := DB.QueryRow("select  count(1) from blog_post where slug = ?", slug)
	if row.Err() != nil {
		return nil, 0, row.Err()
	}
	row.Scan(&count)
	return posts, count, nil
}

func SearchPost(val string) (resp []models.SearchResp, err error) {
	rows, err := DB.Query("select  pid,title from blog_post where title  like ?", "%"+val+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p models.SearchResp
		rows.Scan(
			&p.Pid,
			&p.Title,
		)
		resp = append(resp, p)
	}
	return resp, nil
}

func SavePost(p *models.Post) error {
	sql := "insert into blog_post(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) values (?,?,?,?,?,?,?,?,?,?,)"
	res, err := DB.Exec(sql, p.Title,
		p.Content,
		p.Markdown,
		p.CategoryId,
		p.UserId,
		p.ViewCount,
		p.Type,
		p.Slug,
		p.CreateAt,
		p.UpdateAt,
	)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	p.Pid = int(id)
	return nil
}

func UpdatePost(p *models.Post) error {
	sql := "update  blog_post set title=?,content=?,markdown=?,view_count=?,type=?,slug=?,update_at=? where user_id=? and category_id=? and pid=?"
	_, err := DB.Exec(sql, p.Title,
		p.Content,
		p.Markdown,
		p.ViewCount,
		p.Type,
		p.Slug,
		p.UpdateAt,
		p.UserId,
		p.CategoryId,
		p.Pid,
	)
	if err != nil {
		return err
	}
	return nil

}
