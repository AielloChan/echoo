package libs

import "path"

// TplFile template
type TplFile struct {
	Base     string
	Path     string
	FileName string
}

func (t TplFile) FilePath() string {
	return path.Join(t.Base, t.Path, t.FileName)
}
