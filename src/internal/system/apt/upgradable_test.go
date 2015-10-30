package apt

import "testing"
import "internal/system"
import C "gopkg.in/check.v1"

type testWrap struct{}

func Test(t *testing.T) { C.TestingT(t) }
func init() {
	C.Suite(&testWrap{})
}

func (*testWrap) TestBuildUpgradeInfo(c *C.C) {
	data := []struct {
		Raw            string
		Package        string
		LastVersion    string
		CurrentVersion string
	}{
		{
			"python3-apt/unknown 1.0.1 amd64 [upgradable from: 1.0.0+b1]",
			"python3-apt", "1.0.1", "1.0.0+b1",
		},
		{
			"python3-cairo/unknown 1.10.0+dfsg-5 amd64 [upgradable from: 1.10.0+dfsg-4+b1]",
			"python3-cairo", "1.10.0+dfsg-5", "1.10.0+dfsg-4+b1",
		},
		{
			"python3-dbus/unknown 1.2.0-2+b4 amd64 [upgradable from: 1.2.0-2+b3]",
			"python3-dbus", "1.2.0-2+b4", "1.2.0-2+b3",
		},
		{
			"python3-gi/unknown 3.18.0-1 amd64 [upgradable from: 3.16.2-1]",
			"python3-gi", "3.18.0-1", "3.16.2-1",
		},
		{
			"python3-gi-cairo/unknown 3.18.0-1 amd64 [upgradable from: 3.16.2-1]",
			"python3-gi-cairo", "3.18.0-1", "3.16.2-1",
		},
		{
			"python3-pyqt5/unknown 5.4.2+dfsg-1+b2 amd64 [upgradable from: 5.4.2+dfsg-1+b1]",
			"python3-pyqt5", "5.4.2+dfsg-1+b2", "5.4.2+dfsg-1+b1",
		},
		{
			"python3-pyqt5.qtquick/unknown 5.4.2+dfsg-1+b2 amd64 [upgradable from: 5.4.2+dfsg-1+b1]",
			"python3-pyqt5.qtquick", "5.4.2+dfsg-1+b2", "5.4.2+dfsg-1+b1",
		},
		{
			"python3-pyqt5.qtwebkit/unknown 5.4.2+dfsg-1+b2 amd64 [upgradable from: 5.4.2+dfsg-1+b1]",
			"python3-pyqt5.qtwebkit", "5.4.2+dfsg-1+b2", "5.4.2+dfsg-1+b1",
		},
		{
			"python3-sip/unknown 4.16.9+dfsg-2+b1 amd64 [upgradable from: 4.16.9+dfsg-2]",
			"python3-sip", "4.16.9+dfsg-2+b1", "4.16.9+dfsg-2",
		},
		{
			"qtcreator/unknown 3.5.0+dfsg-2 amd64 [upgradable from: 3.4.1+dfsg-2+b1]",
			"qtcreator", "3.5.0+dfsg-2", "3.4.1+dfsg-2+b1",
		},
		{
			"qtcreator-data/unknown 3.5.0+dfsg-2 all [upgradable from: 3.4.1+dfsg-2]",
			"qtcreator-data", "3.5.0+dfsg-2", "3.4.1+dfsg-2",
		},
		{
			"qtcreator-doc/unknown 3.5.0+dfsg-2 all [upgradable from: 3.4.1+dfsg-2]",
			"qtcreator-doc", "3.5.0+dfsg-2", "3.4.1+dfsg-2",
		},
		{
			"synergy/unknown 1.4.16-1+b1 amd64 [upgradable from: 1.4.16-1]",
			"synergy", "1.4.16-1+b1", "1.4.16-1",
		},
		{
			"systemd/unknown 226-4 amd64 [upgradable from: 226-3]",
			"systemd", "226-4", "226-3",
		},
		{
			"virtualbox/unknown 5.0.6-dfsg-1 amd64 [upgradable from: 5.0.4-dfsg-2]",
			"virtualbox", "5.0.6-dfsg-1", "5.0.4-dfsg-2",
		},
	}

	var lines []string
	for _, item := range data {
		lines = append(lines, item.Raw)
	}
	infos := mapUpgradeInfo(lines, buildUpgradeInfoRegex([]system.Architecture{"amd64", "i386", "alpha"}), buildUpgradeInfo)
	c.Assert(len(infos), C.Equals, len(data))
	for i, item := range data {
		info := infos[i]
		c.Check(info.Package, C.Equals, item.Package)
		c.Check(info.CurrentVersion, C.Equals, item.CurrentVersion)
		c.Check(info.LastVersion, C.Equals, item.LastVersion)
	}
}