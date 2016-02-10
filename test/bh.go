/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 10-02-2016
 * |
 * | File Name:     bh.go
 * +===============================================
 */
package test

import (
	"fmt"

	nom "github.com/1995parham/flynest/nom"
	bh "github.com/kandoo/beehive"
)

func hostJoinedRcvf(msg bh.Msg, ctx bh.RcvContext) error {
	ctx.Printf("Rcv of HostJoinedHandler Called")
	ctx.Printf("%v", msg.Data().(nom.HostJoined))
	return nil
}

func StartTest(hive bh.Hive) error {
	app := hive.NewApp("TestApp")
	fmt.Println("Test app is comming ... :)))")
	app.HandleFunc(nom.HostJoined{}, bh.RuntimeMap(hostJoinedRcvf), hostJoinedRcvf)

	return nil
}
