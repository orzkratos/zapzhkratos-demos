package demokratos

import "github.com/yyle88/runpath"

func SourceRoot() string {
	return runpath.PARENT.Path()
}
