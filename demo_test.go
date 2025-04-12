package demokratos_test

import (
	"path/filepath"
	"testing"

	"github.com/orzkratos/demokratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

func TestShow1Changes(t *testing.T) {
	compareDemo(t, "demo1/go-kratos-demo")
}

func TestShow2Changes(t *testing.T) {
	compareDemo(t, "demo2/go-kratos-demo")
}

func compareDemo(t *testing.T, projectName string) {
	path1 := osmustexist.ROOT(filepath.Join(demokratos.SourceRoot(), projectName))
	t.Log(path1)

	path2 := osmustexist.ROOT(runpath.PARENT.Join(projectName))
	t.Log(path2)

	comparePath(t, path1, path2)
}

func TestCompareDemos(t *testing.T) {
	path1 := osmustexist.ROOT(runpath.PARENT.Join("demo1/go-kratos-demo"))
	path2 := osmustexist.ROOT(runpath.PARENT.Join("demo2/go-kratos-demo"))
	comparePath(t, path1, path2)
}

func comparePath(t *testing.T, path1 string, path2 string) {
	output, err := osexec.NewOsCommand().WithDebug().WithExpectExit(1, "DIFFERENCES FOUND").Exec("diff", "-ru", path1, path2)
	require.NoError(t, err)
	if len(output) == 0 {
		eroticgo.GREEN.ShowMessage("SAME")
	} else {
		t.Log(string(output))
	}
}

func TestCompareModules(t *testing.T) {
	path1 := osmustexist.PATH(runpath.PARENT.Join("demo1/go-kratos-demo/go.mod"))
	path2 := osmustexist.PATH(runpath.PARENT.Join("demo2/go-kratos-demo/go.mod"))
	comparePath(t, path1, path2)
}
