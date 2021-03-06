[![Build Status](https://travis-ci.com/axetroy/go-server.svg?token=QMG6TLRNwECnaTsy6ssj&branch=master)](https://travis-ci.com/axetroy/go-server)
[![Coverage Status](https://coveralls.io/repos/github/axetroy/go-server/badge.svg?branch=master)](https://coveralls.io/github/axetroy/go-server?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/go-server)](https://goreportcard.com/report/github.com/axetroy/go-server)
![License](https://img.shields.io/github/license/axetroy/go-server.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/go-server.svg)

### Golang 实现的基础服务

这是我在闲暇时间写的一些基础服务

写一些工作中常用的服务和实现，以备在以后中用到

想到哪里写哪里, 我会不断的完善它

### 包含哪些服务

- [x] 验证类

  - [x] 注册
  - [x] 登陆
  - [x] 账号激活
  - [x] 忘记密码
  - [x] 双重身份验证
  - [ ] 接入短信验证码服务商
  - [ ] 图片验证码

- [ ] 用户类

  - [x] 登出
  - [x] 获取用户资料
  - [x] 更改用户资料
  - [x] 修改登陆密码
  - [x] 忘记登陆密码
  - [x] 设置交易密码
  - [x] 修改交易密码
  - [x] 忘记交易密码
  - [x] 获取用户已邀请的用户列表
  - 用户头像
    - [ ] 上传用户头像
    - [ ] 第三方头像
  - oAuth2 第三方登陆
    - [ ] 微信
    - [ ] QQ
    - [x] Google
    - [ ] Github

- [x] 钱包类

  - [x] 用户用户钱包
  - [x] 钱包转账

- [ ] 财务流水

  - [ ] 财务日志

- [x] 新闻公告
- [x] 系统通知
- [x] 个人通知

- [x] 上传类
  - [x] 文件上传
    - [x] 获取上传的文件
    - [x] 下载上传的文件
    - [x] 限制文件大小/类型
  - [x] 图片上传

    - [x] 生成缩略图
    - [x] 下载图片
    - [x] 限制图片大小/类型
- [x] 邮件服务

- [x] 静态文件服务

## TODO

- [ ] i18n 的错误信息
- [ ] 分离出管理员接口
- [ ] 启用消息队列
- [ ] 提供 RPC 接口
- [ ] 数据库动态分表

## License

The [MIT License](https://github.com/axetroy/go-server/blob/master/LICENSE)
