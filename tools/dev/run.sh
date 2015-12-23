#!/bin/bash
# In The Name Of God
# ========================================
# [] File Name : run.sh
#
# [] Creation Date : 19-11-2015
#
# [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
# =======================================
echo "github.com/elahejalalpour/beehive-netctrl/..."
go get -u github.com/elahejalalpour/beehive-netctrl/...

if [ -d /tmp/beehive ]; then
	echo "Remove beehive states.. :)"
	rm -Rf /tmp/beehive
fi

echo "Run Beehive-Netctrl :)"
go run main.go
