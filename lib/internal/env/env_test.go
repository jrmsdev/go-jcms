package env

import (
	"os"
	"path/filepath"
	"testing"
)

func init() {
	os.Setenv("JCMS_WEBAPP", webapp)
	os.Setenv("JCMS_BASEDIR", filepath.FromSlash(basedir))
	os.Setenv("JCMS_DATADIR", filepath.FromSlash(datadir))
}

func TestWebappName(t *testing.T) {
	n := WebappName()
	if n != "default" {
		t.Error("invalid webapp name:", n)
	}
}

func TestWebappDir(t *testing.T) {
	okd := absPath(filepath.FromSlash("/opt/jcms/default"))
	d := WebappDir()
	if d != okd {
		t.Log("expected:", okd)
		t.Error("invalid webapp dir:", d)
	}
}

func TestDataDir(t *testing.T) {
	okd := absPath(filepath.FromSlash("/var/opt/jcms/default"))
	d := DataDir()
	if d != okd {
		t.Log("expected:", okd)
		t.Error("invalid webapp data dir:", d)
	}
}

func TestGetEnv(t *testing.T) {
	name := "fakeapp"
	os.Setenv("JCMS_WEBAPP", name)
	defer os.Setenv("JCMS_WEBAPP", "")
	n := WebappName()
	if n != name {
		t.Error("invalid get env webapp name:", n)
	}
}

func TestEnvBasedir(t *testing.T) {
	basedir := absPath(filepath.FromSlash("/jcms/testing"))
	os.Setenv("JCMS_BASEDIR", basedir)
	defer os.Setenv("JCMS_BASEDIR", "")
	d := baseDir()
	if d != basedir {
		t.Log("expected:", basedir)
		t.Error("invalid env basedir:", d)
	}
}

func TestEnvDatadir(t *testing.T) {
	datadir := absPath(filepath.FromSlash("/jcms/testing"))
	os.Setenv("JCMS_DATADIR", datadir)
	defer os.Setenv("JCMS_DATADIR", "")
	d := DataDir()
	if d != datadir {
		t.Log("expected:", datadir)
		t.Error("invalid env datadir:", d)
	}
}
