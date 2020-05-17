```bash
docker run -it -v $(pwd):/app/ -v /var/run/docker.sock:/var/run/docker.sock -v /root/.docker/config.json:/root/.docker/config.json 164073796161.dkr.ecr.ap-northeast-1.amazonaws.com/ci/docker:stretch bash
```

