/*
 * Copyright (C) 2015 ~ 2017 Deepin Technology Co., Ltd.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import "internal/system"
import C "gopkg.in/check.v1"
import "strings"

func (*testWrap) TestParseApt(c *C.C) {
	const d = `Reading package lists... Done
Building dependency tree
Reading state information... Done
Calculating upgrade... Done
The following NEW packages will be installed:
The following packages will be upgraded:
  lastore-daemon abc
The following packages will be upgraded:
1 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
Need to get 7,378 kB of archives.
After this operation, 10.2 kB of additional disk space will be used`

	const upgraded = "The following packages will be upgraded:"

	p := parseAptShowList(strings.NewReader(d), upgraded)
	c.Check(p[0], C.Equals, "lastore-daemon")
	c.Check(len(p), C.Equals, 2)
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
	infos := mapUpgradeInfo(lines, buildUpgradeInfoRegex([]system.Architecture{"amd64", "i386", "alpha"}), buildUpgradeInfo, system.SystemUpdate.JobType())
	c.Assert(len(infos), C.Equals, len(data))
	for i, item := range data {
		info := infos[i]
		c.Check(info.Package, C.Equals, item.Package)
		c.Check(info.CurrentVersion, C.Equals, item.CurrentVersion)
		c.Check(info.LastVersion, C.Equals, item.LastVersion)
	}
}
