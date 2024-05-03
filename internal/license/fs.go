package license

import (
	"bytes"
	"embed"
	"text/template"
)

type LicenseFs struct {
	fs embed.FS
}

func NewLicenseFs(fs embed.FS) *LicenseFs {
	return &LicenseFs{fs: fs}
}

func (l *LicenseFs) readFile(path string) ([]byte, error) {
	return l.fs.ReadFile(path)
}

// ReadFile reads a file from the filesystem and returns the content with the license template data applied.
func (l *LicenseFs) ReadFile(path string) ([]byte, error) {
	data, err := l.readFile(path)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	tpl, err := template.New("lic").Parse(string(data))
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(&buf, getLicTplData())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
