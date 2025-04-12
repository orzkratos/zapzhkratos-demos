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

func GetDemo1Path() string {
	return runpath.PARENT.Join("demo1/go-kratos-demo")
}

func GetDemo2Path() string {
	return runpath.PARENT.Join("demo2/go-kratos-demo")
}

func TestShow1Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo1Path())
	path1 := osmustexist.ROOT(GetDemo1Path())
	comparePath(t, path0, path1)
}

func TestShow2Changes(t *testing.T) {
	path1 := osmustexist.ROOT(demokratos.GetDemo2Path())
	path2 := osmustexist.ROOT(GetDemo2Path())
	comparePath(t, path1, path2)
}

func TestCompareDemos(t *testing.T) {
	path1 := osmustexist.ROOT(GetDemo1Path())
	path2 := osmustexist.ROOT(GetDemo2Path())
	comparePath(t, path1, path2)
}

func comparePath(t *testing.T, path1 string, path2 string) {
	t.Log("path1:", path1)
	t.Log("path2:", path2)
	output, err := osexec.NewOsCommand().WithDebug().WithExpectExit(1, "DIFFERENCES FOUND").Exec("diff", "-ru", path1, path2)
	require.NoError(t, err)
	if len(output) == 0 {
		eroticgo.GREEN.ShowMessage("SAME")
	} else {
		t.Log(string(output))
	}
}

func TestCompareModules(t *testing.T) {
	path1 := osmustexist.PATH(filepath.Join(GetDemo1Path(), "go.mod"))
	path2 := osmustexist.PATH(filepath.Join(GetDemo2Path(), "go.mod"))
	comparePath(t, path1, path2)
}
