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

func init() {
	NotUseDBus = true
}

/*
import (
	proxy "./dbusproxy"
	. "github.com/smartystreets/goconvey/convey"
	//	"github.com/linuxdeepin/go-lib/dbus"
	"fmt"
	"testing"
)

func TestReleaseFDs(t *testing.T) {
	m, _ := proxy.NewManager("com.deepin.lastore", "/com/deepin/lastore")
	jobp, _ := m.InstallPackages([]string{"deepin-movie"})
	job, _ := proxy.NewJob("com.deepin.lastore", jobp)
	m.StartJob(job.Id.Get())
}

func TestSetup(t *testing.T) {
	return
	var job *proxy.Job
	var m *proxy.Manager
	var err error
	var done = make(chan bool)
	ps := []string{"deepin-movie"}
	Convey("Test dbus service features, please setup lastore-dameon before test this", t, func() {
		m, err = proxy.NewManager("com.deepin.lastore", "/com/deepin/lastore")
		So(err, ShouldBeNil)

		Convey(fmt.Sprintf("Try removing the package of %v", ps), func() {
			Convey("Call Manager.RemovePackages ", func() {
				jobp, err := m.RemovePackages(ps)
				So(err, ShouldBeNil)
				job, err = proxy.NewJob("com.deepin.lastore", jobp)
				So(err, ShouldBeNil)

				Convey("Get the Job object from "+string(job.Path)+" and start it", func() {
					r, err := m.StartJob(job.Id.Get())
					So(err, ShouldBeNil)
					So(r, ShouldEqual, true)

					Convey("Wait the package removed", func(c C) {

						job.Progress.ConnectChanged(func() {
							c.Printf("\nAction:%q Name:%q Progress:%f Status:%q\n",
								job.Type.Get(), job.PackageId.Get(), job.Progress.Get(), job.Status.Get())
							if job.Progress.Get() == 1 {
								done <- true
							}
						})
						So(job.Status.Get(), ShouldEqual, "ready")
						<-done
						So(job.Status.Get(), ShouldEqual, "success")

						Convey("Whether this job is still live", func(c C) {
							So(job.Status.Get(), ShouldEqual, "success")
							//So(m.JobList.Get(), ShouldContain, jobp)
						})

						Convey("Clean this Job", func() {
							//So(m.JobList.Get(), ShouldContain, jobp)
						})
						//So(job.Status.Get(), ShouldEqual, "success")
						So(err, ShouldBeNil)
					})
				})

			})

			Convey("Call Manager.InstallPackages ", func() {
				jobp, err := m.InstallPackages(ps)
				So(err, ShouldBeNil)
				job, err = proxy.NewJob("com.deepin.lastore", jobp)
				So(err, ShouldBeNil)

				Convey("Get the Job object from "+string(job.Path)+" and start it", func() {
					So(job.Status.Get(), ShouldEqual, "ready")

					r, err := m.StartJob(job.Id.Get())
					So(err, ShouldBeNil)
					So(r, ShouldEqual, true)

					Convey("Wait the package removed", func(c C) {
						job.Progress.ConnectChanged(func() {
							c.Printf("\nAction:%q Name:%q Progress:%f Status:%q\n",
								job.Type.Get(), job.PackageId.Get(), job.Progress.Get(), job.Status.Get())
							if job.Progress.Get() == 1 {
								done <- true
							}

						})
					})
					<-done

					So(job.Status.Get(), ShouldEqual, "success")
				})

			})
		})
	})
}
*/
