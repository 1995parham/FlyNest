#!/usr/bin/env python
# In The Name Of God
# ========================================
# [] File Name : mn.py
#
# [] Creation Date : 12/24/15
#
# [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
# =======================================
__author__ = 'Elahe Jalalpour'

from mininet.net import Mininet
from mininet.node import RemoteController

# tree4 = TreeTopo(depth=2, fanout=2)

net = Mininet()
net.addController(name='c0', controller=RemoteController(name='beehive-netctrl', ip='192.168.200.1', port=6633))
net.start()
