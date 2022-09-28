# Amigo-bot
## 简介
这是基于 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) ，使用 Go 编写的一个非常简陋的 qq 机器人样例。

> 基本上只是自用的，还请不要期待我会维护，X)

## 食用说明
1. 下载对应平台的 [ go-cqhttp ](https://github.com/Mrs4s/go-cqhttp/releases) 并配置，启动好
2. 把正向 websocket 的 address 填写到`config/application.yaml`中的`bot.ws.addr`下
3. 编译成目标平台的可执行文件，运行

## 插件
说是插件，其实就是简单的 modules <del>, 毕竟 plugin 包 实在太坑了 pwp</del>

这部分的想法是参考 [MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template) 写的

### 如何编写插件
姑且写了一个 [ example ](https://github.com/mosqu1t0/Amigo-bot/blob/master/plugins/ezres/ezres.go), 请参考一下

把插件包放在`plugins/`下后，别忘了在`main.go`中`import`

```go
import (
	"github.com/mosqu1t0/Amigo-bot/bot"
    //...
    // 你的插件包
    //...
)
```
## 版本
模板目前最新为 `v1.0`

编写时，使用 go-cqhttp 的版本：`1.0.0-rc3`
