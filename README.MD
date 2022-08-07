# configDeliverClient 配置分发客户端

## 下载地址

- [linux_amd64](http://hjmcloud.gitee.io/config-deliver-client-release/linux_amd64/config-deliver-client)
- [linux_386](http://hjmcloud.gitee.io/config-deliver-client-release/linux_386/config-deliver-client)
- [darwin_amd64](http://hjmcloud.gitee.io/config-deliver-client-release/darwin_amd64/config-deliver-client)

## 配置文件

```yaml
# config.yaml 位于应用同级目录
logger:
  level: "info"
  stdout: true

apiServer: http://config.lidong.xin/getconfig # api server address
tasks:
  - name: "5472E3F6-37DE-B9A0-899C-39838E8C1336" # 任务名称
    file: "*" # 文件名
    corn: "@every 1m" # corn表达式
    localDir: "./localDir/5472E3F6-37DE-B9A0-899C-39838E8C1336" # 本地目录
    isRunOnStart: true # 是否启动时执行一次
```

## systemd 部署

本部署方式仅支持 linux 系统,且需要安装 systemd.

### 在线安装

进入要安装到的目录，执行以下命令

```shell
curl -sSL https://gitee.com/hjmcloud/scripts/raw/master/config-deliver-client/systemd/config-deliver-client-install.sh | sh
```

### 在线卸载

```shell
curl -sSL https://gitee.com/hjmcloud/scripts/raw/master/config-deliver-client/systemd/config-deliver-client-uninstall.sh | sh
```

### 管理服务

```shell
# 启动服务
systemctl start config-deliver-client
# 停止服务
systemctl stop config-deliver-client
# 重启服务
systemctl restart config-deliver-client
# 查看服务状态
systemctl status config-deliver-client
# 查看服务日志
journalctl -u config-deliver-client -f
```

## Docker compose 部署

下载示例文件

```bash
wget https://gitee.com/hjmcloud/config-deliver-client/raw/master/docker-compose.yml -O docker-compose.yml

wget https://gitee.com/hjmcloud/config-deliver-client/raw/master/config.yaml -O config.yaml
```

目录结构如下

```bash
.
├── config.yaml
└── docker-compose.yml
```

```yaml
# docker-compose.yml
version: "3.7"
services:
  config-deliver-client:
    image: hjmcloud/config-deliver-client:latest
    container_name: config-deliver-client
    restart: always
    volumes:
      - ./config.yaml:/app/config.yaml
      - ./localDir:/app/localDir
    environment:
      - TZ=Asia/Shanghai
```

## GoFrame Template For SingleRepo

Project Makefile Commands:

- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.
- `make bin.publish`: build and publish binary files to gitee pages.
