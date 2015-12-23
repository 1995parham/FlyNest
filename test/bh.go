/*
 * In The Name Of God
 * ========================================
 * [] File Name : bh.go
 *
 * [] Creation Date : 15-11-2015
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Parham Alvani.
 */

package test

import (
	"fmt"

	nom "github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"
)

func hostJoinedRcvf(msg bh.Msg, ctx bh.RcvContext) error {
	fmt.Println("Rcv of HostJoinedHandler Called")
	fmt.Println(msg.Data().(nom.HostJoined))
	return nil
}

func StartTest(hive bh.Hive) error {
	app := hive.NewApp("TestApp")
	fmt.Println("Test app is comming ... :)))")
	app.HandleFunc(nom.HostJoined{}, bh.RuntimeMap(hostJoinedRcvf), hostJoinedRcvf)

	return nil
}
