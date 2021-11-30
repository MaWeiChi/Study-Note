# develop environment

under edge-device-mil folder which cloned from repo:

```[bash]
git clone git@gitlab.com:moxa/ibg/software/platform/thingspro/edge-device-mil.git
```

make build environment with debug container(Dockerfile.debug.amd64/armhf):

```[bash]
# armhf
$ ARCH=armhf make test/dev

# amd64
$ ARCH=amd64 make test/dev
```

success:

```[bash]
....
docker start device
device
docker attach device
root@6a73a0b9c7c4:/data# 
```

---

## troubleshooting

### 1. cloudn't download the depend debs with the makefile command

if you aren't under the MOXA internal network, it will occur this issue:

```[bash]
------
 > [4/6] RUN wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mx-sio-agent/develop/4/build-amd64/libmx-sio-agent1_1.0.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mx-sio-agent/develop/4/build-amd64/libmx-sio-agent-dev_1.0.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mosquitto/feat/support-unixsocket/16/libmosquitto-dev_1.6.8-1%2Bun1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mosquitto/feat/support-unixsocket/16/libmosquitto1_1.6.8-1%2Bun1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/debian/all/libparson1_1.1.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/debian/all/libparson-dev_1.1.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/edge-dx-engine/develop/145/build-amd64/libmx-dx1_1.0.1-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/edge-dx-engine/develop/145/build-amd64/libmx-dx-dev_1.0.1-1_amd64.deb &&     apt-get update &&     apt-get install -y libprotobuf-c1 qemu-user-static &&     apt-get install sshpass &&     apt-get install -y -f ./*.deb:
#7 1.878 --2021-11-29 02:39:57--  https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mx-sio-agent/develop/4/build-amd64/libmx-sio-agent1_1.0.0-1_amd64.deb
#7 1.896 Resolving mxswdc2.s3-ap-northeast-1.amazonaws.com (mxswdc2.s3-ap-northeast-1.amazonaws.com)... 52.219.9.62
#7 2.026 Connecting to mxswdc2.s3-ap-northeast-1.amazonaws.com (mxswdc2.s3-ap-northeast-1.amazonaws.com)|52.219.9.62|:443... connected.
#7 2.312 HTTP request sent, awaiting response... 403 Forbidden
#7 2.426 2021-11-29 02:39:58 ERROR 403: Forbidden.
#7 2.426 
------
executor failed running [/bin/sh -c wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mx-sio-agent/develop/4/build-amd64/libmx-sio-agent1_1.0.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mx-sio-agent/develop/4/build-amd64/libmx-sio-agent-dev_1.0.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mosquitto/feat/support-unixsocket/16/libmosquitto-dev_1.6.8-1%2Bun1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/mosquitto/feat/support-unixsocket/16/libmosquitto1_1.6.8-1%2Bun1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/debian/all/libparson1_1.1.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/debian/all/libparson-dev_1.1.0-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/edge-dx-engine/develop/145/build-amd64/libmx-dx1_1.0.1-1_amd64.deb &&     wget https://mxswdc2.s3-ap-northeast-1.amazonaws.com/v3/edge/builds/edge-dx-engine/develop/145/build-amd64/libmx-dx-dev_1.0.1-1_amd64.deb &&     apt-get update &&     apt-get install -y libprotobuf-c1 qemu-user-static &&     apt-get install sshpass &&     apt-get install -y -f ./*.deb]: exit code: 8
make: *** [test/dev] Error 1
```

mark the `wget command` and `apt-get install -y -f ./*.deb` in Dockerfile.debug.amd64/armhf.
download the dependencies through aws or ibg-edge-s3.moxa.com with vpn mannually when you enter the contianer then install tht depend debs.

open a new terminal, create a new image with the modified device container:

```[bash]
$ docker commit device                  
$ docker commit device                  
sha256:8a3a8ce0a2b5df2a075a380692700a36eb49447cc1f2869c7e4e7284bdc74923
```

exit the container:

```[bash]
root@6a73a0b9c7c4:/data#  exit
```

remove the unfinished container:

```[bash]
# armhf
$ docker remove moxaisd/device-dev-armhf

# amd64
$ docker remove moxaisd/device-dev-amd64
```

tag the completed image:  

```[bash]
# armhf
$ moxaisd/device-dev-armhf

# amd64
$ moxaisd/device-dev-amd64
```

and you can just run debug mode without docker build images:

```[bash]
docker create -it --rm \
    --name device \
    -w /data \
    -v ${PWD}:/data \
    -p 59001:59001 \
    moxaisd/device-dev-$(ARCH) \
    bash
docker start device
docker attach device
```

### 2. the container name "/device" is already in use by container

remove the running container or attach it

```[bash]
# remove
docker remove moxaisd/device-dev-{ARCH}

# attach
docker attach moxaisd/device-dev-{ARCH}
```

---

## feature

### build binary to alter the running device

```[bash]
# armhf
ARCH=armhf make deviced

# amd64
ARCH=amd64 make deviced
```

scp the binary to deviced

```[bash]
scp build/armhf/deviced moxa@192.168.3.127:/tmp
```

ssh to your device and change the running binary

```[bash]
docker cp /tmp/deviced device_app_1:/var/tpdevice/
```

restart app

```[bash]
appman app restart device
```

log

```[bash]
journalctl APPNAME=device -f -a
```

enter device app

```[bash ]
docker exec -ti device_app_1 bash
```
