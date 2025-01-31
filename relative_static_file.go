package jin

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

func RelativeStaticFs(relative string, fsys fs.FS, pathfs string) Handler {
	return func(c Context) (any, error) {
		url := c.Request.URL.String()
		if !strings.Contains(url, relative) {
			return nil, nil
		}
		subFsys := subFs(fsys, pathfs)
		fileName := filepath.Base(url)
		c.FileFromFS(fileName, http.FS(subFsys))
		c.Abort()
		return nil, nil
	}
}
