#!/bin/sh
echo "初始化容器网络:"$1

sleep 2
cd resource
#进行一次拷贝，将busybox拷贝一份
if [ ! -d $1 ];then
mkdir $1
cp -r busybox $1
fi
cd ..
#启动 lo 回环网卡
ip link set lo up
#启动 veth0-docker 网卡
ip link set $1-c up
#给 veth0-docker 配置一个ip地址
ip addr add $2/24 dev $1-c

#容器指定 网桥 为网关
ip route add default via 10.0.0.1