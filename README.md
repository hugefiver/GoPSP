# GoPSP of Pixiv Spider Platform

这是将一个用于`Pixiv`的爬虫系统。

通过事件驱动爬虫线程运行。

事件可以是定时器信号、脚本触发以及外部的请求等。

## Feature

- Server
  - [ ] 登录/爬取组件
  - [ ] 网页解析
  - [ ] 文件储存管理
  - [ ] 配置管理
  - [ ] 爬取队列
  - [ ] 网络检测/失败重试
- Mid
  - [ ] 添加URL/ID爬取
  - [ ] 订阅作者
  - [ ] 定时器
  - [ ] 基于HTTP的外部接口

## TODO

- [ ] ~~可设置代理/保存cookies的HTTP Client~~(还是用既有的爬虫框架好了)
- [ ] 可限制并发的队列下载器
- [ ] 配置文件解析

## Requirements

- `gocolly/colly`: github.com/gocolly/colly