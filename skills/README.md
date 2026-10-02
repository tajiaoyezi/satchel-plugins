# skills

Satchel 的 skills（写给 AI 的操作手册，技术方案第 05 章功能⑥）不在这个仓库：它们放在主控仓库 `satchel` 的 `internal/base/skills/files/`，编进 `satchel` 二进制，由 `satchel mcp init` 装进 AI runtime 的 skills 目录，见 `satchel` 的 README「skills：写给 AI 的操作手册」一节。

放在 `satchel` 里，是为了让 skills 与命令表在同一个仓库：测试直接检查 skills 里写的每条 `satchel` 命令都存在、每条命令都有 skill 讲到，`satchel` 升级时装上的 skills 也总是同一个版本。这个目录只留这份说明。
