# satchel-plugins

Satchel（百宝袋）的二合一仓库，对应 mmwx 的 mmwX-plugins（mmwX-plugins 里的 skills 在 Satchel 里放进了主控仓库 `satchel`，见 `skills/README.md`）：

| 目录 | 是什么 | 形态 | 开始填的里程碑 |
|---|---|---|---|
| `proxyparser/` | 订阅解析库（Surge、Sub-Store 兼容） | Go module `github.com/satchel/satchel-plugins/proxyparser`，主控 `satchel` 以依赖方式引用 | M3 |
| `speedtester/` | 家用测速端：反向连接主控、配对令牌、结果回传 | Go module `github.com/satchel/satchel-plugins/speedtester`，二进制 `satchel-speedtester` | M7 |

**状态：M0 骨架阶段，只有空壳。**

## 发布

打 tag `v*` 触发 `.github/workflows/release.yml`：构建后签名 job 停在受保护环境 `release-signing` 等仓库拥有者批准，签名程序检出 `satchel` 仓库的 `tools/sign` 来跑（一把发布密钥签三个二进制），产物与 `.sig`、`checksums.txt` 挂到 GitHub Release。

## 构建

```sh
cd proxyparser && go test ./...
cd speedtester && go build ./cmd/satchel-speedtester && ./satchel-speedtester --version
```

## 仓库关系

六个仓库的分工见技术方案第 02 章：`satchel`（主控、CLI、MCP、skills）、`satchel-agent`（节点守护）、`satchel-web`（网页前端）、`satchel-plugins`（本仓库）、`satchel-probe`（外置探针）、`satchel-docs`（文档站）。

## 许可证

GPL-3.0，见 LICENSE。
