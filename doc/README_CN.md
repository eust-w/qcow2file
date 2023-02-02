[🇺🇸](../README.md)[English](../README.md)

# :tada:dockerfileparser
qcow2file 是一个从dockerfile 生成qcow2镜像的工具

# 如何使用
1. 下载qcow2file
2. 安装必要的依赖(demo阶段，需要自行安装依赖和运行环境)，libvirt等
3. 运行qcow2file， `qcow2file run -q 基础qcow2镜像 -f 你的docker文件 -o`

| run命令的flag | 解释                               | 说明                  |
| ------------- | ---------------------------------- | --------------------- |
| -q, –qcow     | 基础qcow2镜像                      | 必选参数              |
| -f, –file     | qcow2file，与dockerfile格式一致    | 必选参数              |
| -o, –out      | 输出的qcow2镜像                    | 必选参数              |
| -p, –pause    | 在完成后暂停，阻塞后续的删除vm操作 | 可选参数，默认为false |

# 例子
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

会通过`./dockerfile`生成一个新的qcow2镜像`tem.qcow2`

## 开发注意

1. 命令行的优先级大于qcow2file中的优先级，大于配置表中的优先级

## 后续计划

1. 增加更多dockerfile语法的支持
2. 必选参数`-q,--qcow`更改为可选参数,可以使用qcow2file中from中定义的镜像作为基础镜像
3. 像dockerhub一样的qcow2hub(可能很大)
4. 移除qemu-img，压缩镜像
5. 从libvirt go sdk 依赖转换为对rpc的依赖，从而关闭cgo，移除所有动态链接，增加平台可移植性