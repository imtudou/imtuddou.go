package models

import (
	"html/template"
	"io"
	"strings"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Customer   TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
	About      TemplateBlog
}

func IsODD(i int) bool {
	return i%2 == 0
}

func GetNextName(str []string, index int) string {
	return str[index+1]
}
func Date(fomart string) string {
	return time.Now().Format(fomart)

}

func DateDay(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {
	var htmlTemplate HtmlTemplate
	tb, err := readerHtmlTemplate([]string{"Index", "Category", "Custom", "Detail", "Login", "Pigeonhole", "Writing", "About"},
		templateDir,
	)
	if err != nil {
		return htmlTemplate, err
	}

	htmlTemplate.Index = tb[0]
	htmlTemplate.Category = tb[1]
	htmlTemplate.Customer = tb[2]
	htmlTemplate.Detail = tb[3]
	htmlTemplate.Login = tb[4]
	htmlTemplate.Pigeonhole = tb[5]
	htmlTemplate.Writing = tb[6]
	htmlTemplate.About = tb[7]
	return htmlTemplate, nil
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func readerHtmlTemplate(templateNames []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, i := range templateNames {
		viewName := strings.ToLower(i) + ".html"

		t := template.New(viewName)
		// 因为访问首页的时候有多个模版嵌套，所以解析的时候需要解析全部
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		about := templateDir + "about.html"

		//映射模版定义的方法
		t.Funcs(template.FuncMap{
			"isODD":       IsODD,
			"getNextName": GetNextName,
			"date":        Date,
			"dateDay":     DateDay,
		})

		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination, about) // 解析
		if err != nil {
			return nil, err
		}

		tbs = append(tbs, TemplateBlog{t})
	}
	return tbs, nil
}
