# Nginx in stretch with moxa patch

Debian9 uses the Nginx in main version 1.10.3-1 with the security patchs, deb9u7, and there is a time issue which was fixed in 1.13.10. So we have to patch the commit into 1.10.3-1 to avoid the [issue](https://moxacorp.atlassian.net/browse/TPE-1845).

[Nginx offical documentation](http://nginx.org)

[Debian tracker](https://tracker.debian.org/pkg/nginx)

## Our patch

- ### monotonic time
  
  - [jira issue description](https://moxacorp.atlassian.net/browse/TPE-1845)
  
  - [nginx official commit](https://github.com/nginx/nginx/commit/c7e8a6f2123c653b63ed8013a805eddff502b9ee)
  
## Build Deb

1. clone repo

    ```bash
    git clone git@gitlab.com:moxa/ibg/software/platform/thingspro/nginx.git
    ```

2. select arch

   - build amd64 with **docker**:

     ```bash
     ### pull debian:stretch image
     docker pull debian:stretch

     ### create a container and run it
     docker create -it --rm --name debian9 -w /data -v ${PWD}:/data debian:stretch bash
     docker start debian9
     docker attach debian9
     ```

   - build armhf with **device**:

     ```bash
     scp -q -r nginx moxa@DeviceIP:~

     ssh moxa@DeviceIP
     ```

3. append sources in the `/etc/apt/sources.list`

    ```bash
    echo -e "deb-src http://deb.debian.org/debian stretch main contrib non-free\ndeb-src http://deb.debian.org/debian stretch-updates main contrib non-free\ndeb-src http://security.debian.org/ stretch/updates main contrib non-free" >> /etc/apt/sources.list
    ```

4. install depends

    ```bash
    apt-get update && apt-get install -y dpkg-dev autotools-dev debhelper po-debconf dh-systemd libexpat-dev \
    libgd-dev libgeoip-dev libhiredis-dev libluajit-5.1-dev libmhash-dev libpam0g-dev libpcre3-dev \
    libperl-dev libssl-dev libxslt1-dev quilt zlib1g-dev
    ```

5. get nginx source code

    ```bash
    apt-get source nginx
    ```

6. append patch and changelog

    ```bash
    cp nginx/debian/patches/moxa-monotonic-time.patch nginx/debian/patches/series \
    nginx-1.10.3/debian/patches/ && cp nginx/debian/changelog nginx-1.10.3/debian/
    ```

    [option] build light config

    ```bash
    cp nginx/light/* nginx-1.10.3/debian
    ```

7. build package

    ```bash
    cd nginx-1.10.3 && dpkg-buildpackage -b && ls -l ..
    ```

    done.
