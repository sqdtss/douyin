## 第三届字节青训营后端专场项目，前端为字节提供精简版抖音app，后端基于简单的Gin、Gorm及Mysql实现，后面可能会再对方方面面进行优化

## 安装
```shell
git clone https://github.com/sqdtss/douyin.git
```
#### 手机客户端
将apk文件夹下的apk下载到安卓手机/模拟器中下载，快速双击我的即可配置服务器地址，如http://192.168.3.1:8080/

#### 服务端
首先在mysql中创建数据库"douyin"。

然后修改config.yaml中有关server，mysql，redis，upload的相关配置，现版本还未使用redis，不配置的话把redis相关配置及config，global包中相关内容注释即可。

最后运行以下代码即可运行。
```shell
cd douyin
go generate
go build
./douyin
```
此时手机/模拟器打开app即可使用。

### 接口文档
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

### 修改了很多地方，不过发现客户端仍未发送userId或者发送的userId始终为0（客户端巨大bug）