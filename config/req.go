/*
 * In The Name Of God
 * ========================================
 * [] File Name : req.go
 *
 * [] Creation Date : 18-11-2015
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Parham Alvani.
 */
package config

type ConfigRequest struct {
	AppName         string
	RequiredFeature string
	RequiredData    []byte
}

type ConfigResponse struct {
	AppName      string
	ResponseData []byte
}
