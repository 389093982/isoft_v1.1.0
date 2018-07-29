package common

import "github.com/satori/go.uuid"

func RandomUUID() string {
	return uuid.NewV4().String()
}
