package demokratos

import "github.com/yyle88/runpath"

func GetDemo1Path() string {
	return runpath.PARENT.Join("demo1/go-kratos-demo")
}

func GetDemo2Path() string {
	return runpath.PARENT.Join("demo2/go-kratos-demo")
}
