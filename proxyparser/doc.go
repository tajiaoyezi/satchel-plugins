// Package proxyparser 是订阅解析库：把各家订阅格式解析成统一的节点结构（技术方案附录 B 145）。
// 主控 satchel 以 Go module 依赖引用它，不把源码搬进主控。M0 只有空壳，解析器随 M3 订阅系统交付。
package proxyparser
