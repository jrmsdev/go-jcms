package logger

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func newTestLogger(t *testing.T) *Logger {
	Level(defaultLevel)
	l := New("testing")
	l.tstamp = false
	l.exitOnPanic = false
	tmpfile, err := ioutil.TempFile(os.TempDir(), "jcmstest.logger.")
	if err != nil {
		t.Fatal(err)
	}
	err = File(tmpfile)
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func testCleanup() {
	fn := ""
	if outfh != nil {
		fn = outfh.Name()
	}
	Close()
	if fn != "" {
		os.Remove(fn)
	}
	Level(defaultLevel)
}

func testReadlog(t *testing.T) string {
	blob, err := ioutil.ReadFile(outfh.Name())
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(blob))
}

func TestLogger(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	msg := "message"
	l.log(l.logEntry(ERROR, ""), "test log %s", msg)
	content := testReadlog(t)
	if content != "testing: test log message" {
		t.Error("invalid log content:", content)
	}
}

func TestD(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	Level(DEBUG)
	l.D("mes%s", "sage")
	content := testReadlog(t)
	if content != "[D] testing: message" {
		t.Error("invalid log content:", content)
	}
}

func TestE(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	Level(ERROR)
	l.E("mes%s", "sage")
	content := testReadlog(t)
	if content != "[E] testing: message" {
		t.Error("invalid log content:", content)
	}
}

func TestV(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	Level(VERBOSE)
	l.V("mes%s", "sage")
	content := testReadlog(t)
	if content != "testing: message" {
		t.Error("invalid log content:", content)
	}
}

func TestW(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	Level(WARNING)
	l.W("message")
	content := testReadlog(t)
	if content != "[W] testing: message" {
		t.Error("invalid log content:", content)
	}
}

func TestPanic(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	Level(QUIET) // should log even in quiet mode
	l.Panic("message")
	content := testReadlog(t)
	if content != "[PANIC] testing: message" {
		t.Error("invalid log content:", content)
	}
}

func TestCloseNil(t *testing.T) {
	newTestLogger(t)
	testCleanup()
	err := Close()
	if err != nil {
		t.Error("logger Close failed:", err)
	}
}

func TestLevels(t *testing.T) {
	l := newTestLogger(t)
	defer testCleanup()
	t.Log("DEBUG")
	Level(DEBUG)
	{
		l.E("message")
		checkRead(t, "DE")
		l.V("message")
		checkRead(t, "DV")
		l.W("message")
		checkRead(t, "DW")
		l.D("message")
		checkRead(t, "DD")
	}
	t.Log("VERBOSE")
	Level(VERBOSE)
	{
		l.D("message")
		checkEmpty(t, "VD")
		l.W("message")
		checkRead(t, "VW")
		l.E("message")
		checkRead(t, "VE")
		l.V("message")
		checkRead(t, "VV")
	}
	t.Log("WARNING")
	Level(WARNING)
	{
		l.D("message")
		checkEmpty(t, "WD")
		l.V("message")
		checkEmpty(t, "WV")
		l.E("message")
		checkRead(t, "WE")
		l.W("message")
		checkRead(t, "WW")
	}
	t.Log("ERROR")
	Level(ERROR)
	{
		l.D("message")
		checkEmpty(t, "WD")
		l.V("message")
		checkEmpty(t, "WV")
		l.W("message")
		checkEmpty(t, "WW")
		l.E("message")
		checkRead(t, "WE")
	}
	t.Log("QUIET")
	Level(QUIET)
	{
		l.D("message")
		checkEmpty(t, "WD")
		l.V("message")
		checkEmpty(t, "WV")
		l.W("message")
		checkEmpty(t, "WW")
		l.E("message")
		checkEmpty(t, "WE")
	}
}

func checkRead(t *testing.T, id string) {
	content := testReadlog(t)
	if !strings.HasSuffix(content, ": message") {
		t.Fatal(id, "invalid content:", content)
	}
	err := outfh.Sync()
	if err != nil {
		t.Fatal(id, err)
	}
	err = outfh.Truncate(0)
	if err != nil {
		t.Fatal(id, err)
	}
	err = outfh.Sync()
	if err != nil {
		t.Fatal(id, err)
	}
	_, err = outfh.Seek(0, 0)
	if err != nil {
		t.Fatal(id, err)
	}
}

func checkEmpty(t *testing.T, id string) {
	content := testReadlog(t)
	if content != "" {
		t.Fatal(id, "content not empty:", content)
	}
	err := outfh.Sync()
	if err != nil {
		t.Fatal(id, err)
	}
	err = outfh.Truncate(0)
	if err != nil {
		t.Fatal(id, err)
	}
	err = outfh.Sync()
	if err != nil {
		t.Fatal(id, err)
	}
	_, err = outfh.Seek(0, 0)
	if err != nil {
		t.Fatal(id, err)
	}
}
