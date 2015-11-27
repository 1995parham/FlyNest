/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     re.go
 * +===============================================
 */
package config

import (
	"encoding/gob"
)

type HTTPRequest struct {
	AppName string
	Verb    string
	Data    []byte
}

type HTTPResponse struct {
	AppName string
	Data    []byte
}

func init() {
	gob.Register(HTTPRequest{})
	gob.Register(HTTPResponse{})
}
