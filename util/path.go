package util

import (
	"path"

	"github.com/cacastelmeli/eir/constants"
)

func PathCacheTemplate(parts ...string) string {
	parts = append([]string{constants.CacheTemplateDirName}, parts...)
	return path.Join(parts...)
}
