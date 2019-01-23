package utils

import "github.com/satori/go.uuid"

//随机uuid
func RandomUUID() (uuid.UUID, error) {
	return uuid.NewV4()
}
