// Generated by ego.
// DO NOT EDIT

package auth

import (
	"fmt"
	"github.com/xirtah/gopa-framework/core/util"
	"github.com/xirtah/gopa-spider/modules/ui/common"
	"html"
	"io"
	"net/http"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Login(w http.ResponseWriter, url string) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Login", "")
	_, _ = io.WriteString(w, "\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n<div class=\"tm-middle\">\n    <div class=\"uk-container uk-container-center\">\n        <br/>\n        <div class=\"uk-grid\">\n            <div class=\"uk-align-center\"><br>\n                <a href=\"/auth/github/")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
	_, _ = io.WriteString(w, "\" class=\"uk-icon-hover uk-icon-medium uk-icon-github\"> Github login</a>\n            </div>\n        </div>\n    </div>\n</div>\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
func LoginFail(w http.ResponseWriter) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Login Failed", "")
	_, _ = io.WriteString(w, "\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n<div class=\"tm-middle\">\n    <div class=\"uk-container uk-container-center\">\n        <br/>\n        <div class=\"uk-grid\">\n            <div class=\"uk-align-center\"><br>\n                <div class=\"uk-alert uk-alert-danger\">Login failed.</div>\n            </div>\n        </div>\n    </div>\n</div>\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
func LoginSuccess(w http.ResponseWriter, url string) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Login Success", "")
	_, _ = io.WriteString(w, "\n\n<META HTTP-EQUIV=\"refresh\" CONTENT=\"0;URL=")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
	_, _ = io.WriteString(w, "\">\n\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n<div class=\"tm-middle\">\n    <div class=\"uk-container uk-container-center\">\n        <br/>\n        <div class=\"uk-grid\">\n            <div class=\"uk-align-center\"><br>\n                <div>Redirecting to: ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
	_, _ = io.WriteString(w, "</div>\n\n                <div class=\"uk-alert uk-alert-success\">\n                    <a href=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(util.UrlDecode(url))))
	_, _ = io.WriteString(w, "\" class=\"uk-icon-hover uk-icon-medium\"> Login success, click to continue.</a></div>\n            </div>\n        </div>\n    </div>\n</div>\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
func Logout(w http.ResponseWriter, url string) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Logout Success", "")
	_, _ = io.WriteString(w, "\n\n<META HTTP-EQUIV=\"refresh\" CONTENT=\"5;URL=")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
	_, _ = io.WriteString(w, "\">\n\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n<div class=\"tm-middle\">\n    <div class=\"uk-container uk-container-center\">\n        <br/>\n        <div class=\"uk-grid\">\n            <div class=\"uk-align-center\"><br>\n                <div>Redirecting to: ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
	_, _ = io.WriteString(w, "</div>\n\n                <div class=\"uk-alert uk-alert-success\">\n                    <a href=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(util.UrlDecode(url))))
	_, _ = io.WriteString(w, "\" class=\"uk-icon-hover uk-icon-medium\"> Logout success, click to continue.</a></div>\n            </div>\n        </div>\n    </div>\n</div>\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
