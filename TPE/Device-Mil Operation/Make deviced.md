# Make deviced

```bash
$ make test/dev ARCH=armhf
make test/dev ARCH=armhf 
docker build \
        -t moxaisd/device-dev-armhf \
        -f Dockerfile.debug.armhf \
        .
....
95b5b127abe21d7dcb6c5b2e0c27764a1a6b1db095580631f91d1f092a38abf2
docker start device
device
docker attach device
root@95b5b127abe2:/data#
```

root@95b5b127abe2

```bash
/data# make deviced
```

