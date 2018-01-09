package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWebappName(t *testing.T) {
	n := WebappName()
	if n != "default" {
		t.Error("invalid webapp name:", n)
	}
}

func TestWebappDir(t *testing.T) {
	okd := filepath.FromSlash("/opt/jcms/default")
	d := WebappDir()
	if d != okd {
		t.Error("invalid webapp dir:", d)
	}
}

func TestSettingsFile(t *testing.T) {
	okfn := filepath.FromSlash("/opt/jcms/default/settings.xml")
	fn := SettingsFile()
	if fn != okfn {
		t.Error("invalid webapp settings file:", fn)
	}
}

func TestDataDir(t *testing.T) {
	okd := filepath.FromSlash("/var/opt/jcms/default")
	d := DataDir()
	if d != okd {
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
	basedir := filepath.FromSlash("/jcms/testing")
	os.Setenv("JCMS_BASEDIR", basedir)
	defer os.Setenv("JCMS_BASEDIR", "")
	d := baseDir()
	if d != basedir {
		t.Error("invalid env basedir:", d)
	}
}

func TestEnvDatadir(t *testing.T) {
	datadir := filepath.FromSlash("/jcms/testing")
	os.Setenv("JCMS_DATADIR", datadir)
	defer os.Setenv("JCMS_DATADIR", "")
	d := DataDir()
	if d != datadir {
		t.Error("invalid env datadir:", d)
	}
}
