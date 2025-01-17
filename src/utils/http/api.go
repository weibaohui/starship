// Copyright (C) 2023  Tricorder Observability
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package http

import (
	"fmt"
	"strings"
)

// API path component
const (
	API_ROOT      = "api"
	LIST_CODE     = "listCode"
	ADD_CODE      = "addCode"
	UPLOAD        = "uploadFile"
	DEPLOY        = "deploy"
	UN_DEPLOY     = "undeploy"
	DELETE_MODULE = "deleteCode"
)

// GetAPIUrl returns a http URL that corresponds to the requested path.
func GetAPIUrl(addr, root, path string) string {
	return fmt.Sprintf("http://%s", strings.Join([]string{addr, root, path}, "/"))
}
