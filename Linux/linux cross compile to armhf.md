# linux cross compile to armhf

reference: https://www.96boards.org/documentation/guides/crosscompile/commandline.html

---

#### Step 3: Install cross compilers on host machine

**For ARM 32bit toolchain**

```
$ sudo apt-get install gcc-arm-linux-gnueabihf g++-arm-linux-gnueabihf
```

#### Step 4: Install package dependencies

```
$ sudo apt-get install build-essential autoconf libtool cmake pkg-config git python-dev swig3.0 libpcre3-dev nodejs-dev
```