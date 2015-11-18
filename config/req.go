/*
 * In The Name Of God
 * ========================================
 * [] File Name : req.go
 *
 * [] Creation Date : 18-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
 */
package config

import (
	"encoding/gob"
)

type ConfigRequest struct {
	AppName         string
	RequiredFeature string
	RequiredData    []byte
}

type ConfigResponse struct {
	AppName      string
	ResponseData []byte
}

func init() {
	gob.Register(ConfigRequest{})
	gob.Register(ConfigResponse{})
}
