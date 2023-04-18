package builder

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/koki-develop/qiita-ranking/internal/config"
	"github.com/koki-develop/qiita-ranking/internal/qiita"
)

type Template string

const (
	TemplateLikesDaily Template = "likes_daily"
)

type Builder struct{}

func New() *Builder {
	return &Builder{}
}

type BuildParameters struct {
	Template   Template
	Tags       config.ConfigItems
	Conditions map[string]string
	Items      qiita.Items
}

func (b *Builder) Build(params *BuildParameters) ([]byte, error) {
	filename := fmt.Sprintf("%s.md", params.Template)
	filenames := []string{
		fmt.Sprintf("templates/%s", filename),
		"templates/_tags.md",
		"templates/_conditions.md",
		"templates/_github.md",
		"templates/_ranking.md",
	}
	tpl, err := template.New(filename).Funcs(template.FuncMap{"inc": b.inc}).ParseFiles(filenames...)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	params.Items.Sort()
	if err := tpl.Execute(buf, params); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (*Builder) inc(i int) int {
	return i + 1
}
