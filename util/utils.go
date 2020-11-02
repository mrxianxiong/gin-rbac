/**
 * @Author: xianxiong
 * @Date: 2020/11/1 14:39
 */

package util

import (
	"fmt"
	guuid "github.com/google/uuid"
)

var uuid string

func genUUID() string {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())
	uuid = id.String()
	return uuid
}
