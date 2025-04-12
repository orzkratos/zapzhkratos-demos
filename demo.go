package demokratos

import (
	"github.com/orzkratos/demokratos/demo1kratos"
	"github.com/orzkratos/demokratos/demo2kratos"
)

func GetDemo1Path() string {
	return demo1kratos.SourceRoot()
}

func GetDemo2Path() string {
	return demo2kratos.SourceRoot()
}
