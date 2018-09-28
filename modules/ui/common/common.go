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

package common

import (
	"bytes"
)

// NavCurrent return the current nav html code snippet
func NavCurrent(cur, nav string) string {
	if cur == nav {
		return " class=\"uk-active\" "
	}
	return ""
}

type navObj struct {
	name        string
	displayName string
	url         string
}

var navs []navObj

// RegisterNav register a custom nav link
func RegisterNav(name, displayName string, url string) {
	obj := navObj{name: name, displayName: displayName, url: url}
	navs = append(navs, obj)
}

// GetJSBlock return a JS wrapped code block
func GetJSBlock(buffer *bytes.Buffer, js string) {

	buffer.WriteString("<script type=\"text/javascript\">")
	buffer.WriteString("    $(function() {")
	buffer.WriteString(js)
	buffer.WriteString("   });")
	buffer.WriteString("</script>")

}
