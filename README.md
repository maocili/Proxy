# proxy_pool

一个简单的动态代理池

A simple proxy pool written in go

## 功能

 - 定时抓取公开免费的代理
 - 定时验证可用代理
 - 支持动态代理(仅支持http)
 - 支持web api
 - ip 等级机制

## 使用
### 编译
```bash
go build cmd
```
### 端口介绍
- :8080 web接口
- :10088 http动态代理端口

### web接口介绍 
- /list 展示代理池所有的ip
- /rand 随机抽取一个ip（默认等级>=60）

## 目录介绍
- /cmd 主程序入口

- /internal 程序内部包

- /pkg 第三方包封装

- /spiderProject 公开代理爬虫

## ip等级机制
```初始等级50，定时校验，失败-30，成功+10。上不封顶，等级小于等于0时删除IP。``` 

## TODO 

- [x] 支持动态代理http、https

- [ ] 支持动态代理socket

- [ ] 更丰富的IP属性

- [ ] 在randIP 时可以分类获取IP 
