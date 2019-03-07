/*
Copyright 2018 Sameer Saini (sameer.saini@outlook.com)

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

package modules

import (
	"github.com/xirtah/gopa-framework/core/module"
	"github.com/xirtah/gopa-framework/modules/persist"
	"github.com/xirtah/gopa-framework/modules/storage"
	"github.com/xirtah/gopa-ui/modules/index"
	"github.com/xirtah/gopa-ui/modules/ui"
)

// Register is where modules are registered
func Register() {
	module.Register(module.Index, index.IndexModule{})
	module.Register(module.System, ui.UIModule{})
	module.Register(module.Database, persist.DatabaseModule{})
	module.Register(module.Storage, storage.StorageModule{})
}
