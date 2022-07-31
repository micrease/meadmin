package strings

import uuid "github.com/satori/go.uuid"

//封装一下,方便把uuid替换别的方案
func GenUuid() string {
	return uuid.NewV4().String()
}
