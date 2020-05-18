

```bash
docker run -it -v $(pwd):/app/ -v /var/run/docker.sock:/var/run/docker.sock -v /root/.docker/config.json:/root/.docker/config.json 164073796161.dkr.ecr.ap-northeast-1.amazonaws.com/ci/docker:stretch bash
```

```bash
cd /app
export ARCH=armhf
make mpkg MODEL=$(modle_name)
```

```bash
scp 
appman app install 
```

Author: ErikWCMa <ericma19920728@gmail.com>
Date:   Mon May 11 17:25:45 2020 +0800

    fix: mxsparkplug.service exec message remind