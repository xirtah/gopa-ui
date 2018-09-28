/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package admin

import (
	api "github.com/xirtah/gopa-framework/core/http"
	"github.com/xirtah/gopa-framework/core/model"
	"github.com/xirtah/gopa-spider/modules/ui/admin/ajax"
	"github.com/xirtah/gopa-spider/modules/ui/common"
)

// InitUI register ui handlers
func InitUI() {
	//Nav init
	common.RegisterNav("Console", "Console", "/admin/console/")
	//common.RegisterNav("Dashboard", "Dashboard", "/admin/")
	common.RegisterNav("Tasks", "Tasks", "/admin/tasks/")
	//common.RegisterNav("Explore","Explore","/ui/explore/")

	//common.RegisterNav("Setting", "Setting", "/admin/setting/")

	//UI pages init
	ui := AdminUI{}

	api.HandleUIMethod(api.GET, "/screenshot/:id", ui.GetScreenshotAction)

	api.HandleUIMethod(api.GET, "/admin/", api.NeedPermission(model.PERMISSION_ADMIN_MINIMAL, ui.DashboardAction))
	api.HandleUIMethod(api.POST, "/admin/setting/", api.NeedPermission(model.PERMISSION_ADMIN_MINIMAL, ui.UpdateSettingAction))
	api.HandleUIMethod(api.GET, "/admin/dashboard/", api.NeedPermission(model.PERMISSION_ADMIN_MINIMAL, ui.DashboardAction))
	api.HandleUIMethod(api.GET, "/admin/tasks/", api.NeedPermission(model.PERMISSION_ADMIN_MINIMAL, ui.TasksPageAction))
	api.HandleUIMethod(api.GET, "/admin/task/view/:id", api.NeedPermission(model.PERMISSION_ADMIN_MINIMAL, ui.TaskViewPageAction))
	api.HandleUIMethod(api.GET, "/admin/console/", api.NeedPermission(model.PERMISSION_ADMIN_MINIMAL, ui.ConsolePageAction))

	api.HandleUIFunc("/admin/explore/", ui.ExplorePageAction)
	api.HandleUIFunc("/admin/setting/", ui.SettingPageAction)

	//Ajax
	ajax := ajax.Ajax{}
	api.HandleUIFunc("/setting/logger", ajax.LoggingSettingAction)
	api.HandleUIFunc("/setting/logger/", ajax.LoggingSettingAction)

}
