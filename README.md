tryGraphQL
==========

基于 GraphQL 协议的 GeoIP 查询 API 的简单实现

[live demo](https://soulogic.com/graphql/)

准备工作
--------

首先获取并编译本代码库

    go get github.com/zhengkai/tryGraphQL

在运行前还需要考虑

**GeoIP 库获取：**

* 免费的数据库在 <http://dev.maxmind.com/geoip/legacy/geolite/> 下载，
可将 `http://geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.gz` 下载并解压至本地的 `/usr/share/GeoIP/GeoLite2-City.mmdb`
* 推荐直接用脚本更新 <https://github.com/extremeshok/geoip-update>
* 付费的数据 `GeoIP2-City.mmdb` 和 `GeoIP2-ISP.mmdb` 为可选项

**Web 环境：**

建议按照 [nginx.txt](https://github.com/zhengkai/tryGraphQL/blob/master/www/nginx.txt) 将代码目录映射到已有域名下做测试

**jq：**

在命令行下调试时很有用的一个 JSON 高亮工具，Ubuntu 下可以直接安装

    apt-get install jq

运行
----

执行编译好的 `tryGraphQL`，会监听 `59999` 端口
这时进入代码目录可以运行 [`./client.sh`](https://github.com/zhengkai/tryGraphQL/blob/master/client.sh) 来测试请求

如果希望通过 web 来访问，注意修改 [`www/script.js`](https://github.com/zhengkai/tryGraphQL/blob/master/www/script.js#L62)，替换掉原有的 url `https://soulogic.com/graphql/api`    
预定义 query 保存在 [`www/query`](https://github.com/zhengkai/tryGraphQL/tree/master/www/query) 目录，其中 [`06.txt`](https://github.com/zhengkai/tryGraphQL/blob/master/www/query/06.txt) 提供了完整的所有可查询字段
