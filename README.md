# satchel-plugins

Satchel（百宝袋）的三合一仓库，对应 mmwx 的 mmwX-plugins：

| 目录 | 是什么 | 形态 | 开始填的里程碑 |
|---|---|---|---|
| `proxyparser/` | 订阅解析库（Surge、Sub-Store 兼容） | Go module `github.com/satchel/satchel-plugins/proxyparser`，主控 `satchel` 以依赖方式引用 | M3 |
| `speedtester/` | 家用测速端：反向连接主控、配对令牌、结果回传 | Go module `github.com/satchel/satchel-plugins/speedtester`，二进制 `satchel-speedtester` | M7 |
| `skills/` | 给 AI 用的 skills 手册，每个 skill 一个子目录 | 纯文件 | M1 起随功能交付 |

**状态：M0 骨架阶段，只有空壳。**

## 构建

```sh
cd proxyparser && go test ./...
cd speedtester && go build ./cmd/satchel-speedtester && ./satchel-speedtester --version
```

## 仓库关系

六个仓库的分工见技术方案第 02 章：`satchel`（主控、CLI、MCP）、`satchel-agent`（节点守护）、`satchel-web`（网页前端）、`satchel-plugins`（本仓库）、`satchel-probe`（外置探针）、`satchel-docs`（文档站）。

## 许可证

GPL-3.0，见 LICENSE。
