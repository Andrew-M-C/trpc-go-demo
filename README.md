# trpc-go-demo

- [trpc-go-demo](#trpc-go-demo)
  - [app/http-auth-server](#apphttp-auth-server)
  - [app/achievement](#appachievement)


tRPC-Go 的 demo。这个项目是服务于作者的 [博客](https://cloud.tencent.com/developer/user/1307425) 中关于 [tRPC 的介绍文章](https://cloud.tencent.com/developer/search/article-%E8%85%BE%E8%AE%AF%20tRPC-Go%20%E6%95%99%E5%AD%A6%20%E5%90%8E%E5%8F%B0%E5%85%A8%E6%A0%88%E4%B9%8B%E8%B7%AF) 的, 请读者移步后搭配文章阅读。

## [app/http-auth-server](./app/http-auth-server)

原来只是作为一个登录接口的示例服务, 现在多实现了几个接口, 反正都是示例, 就随意了

调试命令举例:

- `curl http://127.0.0.1:8001/demo/auth/UserSpace?username=amc`

## [app/achievement](./app/achievement)

用户成就服务
