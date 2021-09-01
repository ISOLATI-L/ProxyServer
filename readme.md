## 数据库配置
### 数据库配置文件
根目录下的SQL.config.ini设置数据库信息：
```
[SQL_Config]
server   = <地址>
port     = <端口>
user     = <用户名>
password = <密码>
database = <数据库名称>
```
### 数据库设置
#### blacklist表
```
CREATE TABLE `blacklist` (
  `Bid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `host` varchar(64) NOT NULL,
  PRIMARY KEY (`Bid`)
)
```
