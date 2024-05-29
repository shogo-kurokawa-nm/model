package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type View struct {
	Template *template.Template
}

func NewView(templatePath string) (*View, error) {
	// 現在の作業ディレクトリを取得
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// テンプレートファイルのパスを生成
	indexPath := filepath.Join(cwd, templatePath)

	// テンプレートをパース
	tmpl, err := template.ParseFiles(indexPath)
	if err != nil {
		return nil, err
	}

	return &View{Template: tmpl}, nil
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.Execute(w, data)
}
