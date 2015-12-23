#!/usr/bin/env python
# In The Name Of God
# ========================================
# [] File Name : mn.py
#
# [] Creation Date : 12/24/15
#
# [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
# =======================================
from mininet.net import Mininet
from mininet.net import CLI
from mininet.log import setLogLevel
from mininet.node import RemoteController
from mininet.topo import Topo
import argparse

__author__ = 'Elahe Jalalpour'


#
# We want to build following topology:
# h1 --- s1 --- s2 --- h2
#
class PathTopo(Topo):
    def build(self, *args, **params):
        h1 = self.addHost(name='h1')
        s1 = self.addSwitch(name='s1')
        s2 = self.addSwitch(name='s2')
        h2 = self.addHost(name='h2')
        self.addLink(h1, s1)
        self.addLink(s2, s1)
        self.addLink(h2, s2)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--ip', dest='ip', help='Beehive Network Controller IP Address', default='127.0.0.1', type=str)
    args = parser.parse_args()

    setLogLevel('info')

    net = Mininet(topo=PathTopo())
    net.addController(cls=RemoteController, name='beehive-netctrl', ip=args.ip, port=6633)
    net.start()
    CLI(net)
    net.stop()
