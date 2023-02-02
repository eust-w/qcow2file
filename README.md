[🇨🇳](doc/README_CN.md)[中文](doc/README_CN.md)

# :tada:dockerfileparser

qcow2file is a tool to generate qcow2 image from dockerfile


# Usage

1.  Download qcow2file
2.  Install the necessary dependencies (demo stage, you need to install dependencies and operating environment by yourself), libvirt, etc.
3.  Run qcow2file, `qcow2file run -q base qcow2 image -f your docker file -o`


| The flag of the run command | Explanation | Description |
| ------------- | ------------------------------------- | --------------------- |
| -q, –qcow | base qcow2 image | required parameter |
| -f, –file | qcow2file, consistent with dockerfile format | Mandatory parameters |
| -o, –out | Output qcow2 image | Mandatory parameter |
| -p, --pause | Pause after completion, blocking subsequent delete vm operations | optional parameter, default is false |

# Sample

`sudo ./qcow2file run -q c74-minimal.qcow2 -f ./dockerfile -o tem.qcow2 --pause`

```bash
[root@localhost ~]# sudo ./qcow2file run -q c74-minimal.qcow2 -f ./dockerfile -o tem.qcow2 --pause
cp c74-minimal.qcow2 tem.qcow2
RUN:cd ./;adduser -u 10001 -D app-runner out: err:<nil>
COPY : CP ./echo.sh echo.sh <nil>
RUN:cd ./;chmod +x echo.sh out: err:<nil>
RUN:cd ./;./echo.sh out:Loaded plugins: fastestmirror
Could not retrieve mirrorlist http://mirrorlist.centos.org/?release=7&arch=x86_64&repo=os&infra=stock error was
14: curl#7 - "Failed to connect to 2604:1580:fe02:2::10: Network is unreachable"
 err:<nil>
COPY : CP ./pkg . <nil>
RUN:cd ./;echo lfdjlsfjd > 2m2d9999999.tem out: err:<nil>
RUN:cd ./;echo 34024093240 >> 2m2d9999999.tem out: err:<nil>
COPY : CP ./src . <nil>
RUN:cd ./;mkdir 2m2d out: err:<nil>
RUN:cd ./;chmod +x ./src/echo.sh out: err:<nil>
RUN:cd ./;./src/echo.sh out:Loaded plugins: fastestmirror
Could not retrieve mirrorlist http://mirrorlist.centos.org/?release=7&arch=x86_64&repo=os&infra=stock error was
14: curl#7 - "Failed to connect to 2600:1f16:c1:5e01:4180:6610:5482:c1c0: Network is unreachable"
 err:<nil>
RUN:cd ./;yum remove git -y out:Loaded plugins: fastestmirror
No Packages marked for removal
 err:<nil>
RUN:cd ./;sync out: err:<nil>
pause vm, enter any key to destroy

out qcow2 is: /root/tem.qcow2
```

A new qcow2 image `tem.qcow2` will be generated through `./dockerfile`

## Dev note

1. The priority of the command line is greater than the priority in qcow2file, which is greater than the priority in the configuration table

## todo

1. Add more dockerfile syntax support
2. The required parameter `-q, --qcow` is changed to an optional parameter, and the image defined in from in qcow2file can be used as the base image