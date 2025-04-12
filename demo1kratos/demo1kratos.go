package demo1kratos

import "github.com/yyle88/runpath"

func SourceRoot() string {
	return runpath.PARENT.Path()
}
