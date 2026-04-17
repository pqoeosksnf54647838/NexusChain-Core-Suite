# NexusChain-Core-Suite
NexusChain 是一款高性能、模块化、安全可扩展的区块链底层核心项目，基于 Go 语言构建，支持 PoS 共识、跨链桥、默克尔证明、零知识验证、UTXO 模型、分片存储、轻客户端等全栈区块链功能，适用于公链、联盟链、Layer2 等多场景应用。

## 项目特性
- 原生区块链核心引擎与区块生成机制
- 高性能 P2P 点对点网络
- 安全的交易签名与多签钱包
- 质押奖励 + 惩罚机制的 PoS 共识
- 智能合约虚拟机与状态存储
- 跨链桥、分片、IPFS 集成
- 零知识证明、轻客户端、链重组
- 完整 CLI、RPC、监控、索引系统

## 文件清单与功能
1. **blockchain_core.go** - 区块链核心引擎，实现区块结构、哈希计算、工作量证明
2. **p2p_network.go** - P2P 节点网络模块，支持节点监听、连接、消息传输
3. **consensus_pos.go** - 权益证明共识，按质押权重随机选择验证节点
4. **transaction_sign.go** - 交易 ECDSA 签名与验签模块
5. **merkle_tree.go** - 默克尔树实现，支持交易批量哈希验证
6. **wallet_manager.go** - 区块链钱包生成，公钥/私钥/地址管理
7. **utxo_model.go** - UTXO 模型，余额计算与未花费输出管理
8. **chain_sync.go** - 区块链自动同步，高度检测与区块同步
9. **smart_contract_vm.go** - 极简合约虚拟机，栈运算与存储操作
10. **gas_calculator.go** - Gas 费计算器，支持转账/存储/执行计费
11. **block_validator.go** - 区块全量验证，哈希、难度、前序区块校验
12. **peer_discovery.go** - P2P 节点自动发现与可用节点筛选
13. **token_standard.go** - 标准化通证实现，转账、余额查询
14. **cross_chain_bridge.go** - 跨链桥模块，支持跨链转账创建与确认
15. **block_indexer.go** - 区块/交易索引器，快速查询区块哈希与交易高度
16. **crypto_aes.go** - AES 加密解密工具，用于链上敏感数据存储
17. **stake_reward.go** - 质押收益计算，按日/月分发奖励
18. **tx_pool.go** - 交易池模块，支持交易添加、移除、批量获取
19. **chain_monitor.go** - 链状态监控，实时出块间隔与 TPS 统计
20. **contract_storage.go** - 合约键值存储，支持读写删与存在校验
21. **ecdsa_utils.go** - ECDSA 密钥工具，十六进制格式转换
22. **block_serialize.go** - 区块序列化/反序列化，用于存储与传输
23. **node_cli.go** - 节点命令行交互工具，支持区块/节点/余额查询
24. **multi_signature.go** - 多签钱包，支持多账户签名确认
25. **rpc_server.go** - RPC 服务，提供远程链数据查询接口
26. **data_sharding.go** - 数据分片管理，提升链存储与处理性能
27. **zero_knowledge.go** - 零知识证明生成与验证模块
28. **block_reorg.go** - 链重组管理器，自动切换最长合法链
29. **state_db.go** - 账户状态数据库，管理余额、合约代码、合约存储
30. **dns_seed.go** - DNS 种子节点解析，自动发现网络节点
31. **tx_compress.go** - 交易数据压缩，减少网络传输体积
32. **epoch_manager.go** - 共识周期管理器，控制验证节点轮换
33. **light_client.go** - 轻客户端，仅同步区块头实现快速验证
34. **fee_market.go** - 手续费市场，按手续费排序交易
35. **block_builder.go** - 区块构建器，批量打包交易生成区块
36. **ipfs_integration.go** - IPFS 集成，数据上传下载
37. **slashing_guard.go** - 恶意节点惩罚，黑名单与质押扣除
38. **genesis_config.go** - 创世区块配置，链初始参数管理

## 技术栈
- 主语言：Go
- 加密：SHA256、ECDSA、AES、零知识证明
- 网络：P2P、RPC、DNS 种子
- 存储：键值存储、UTXO、状态库、IPFS
- 共识：PoS、周期验证、惩罚机制
- 扩展：跨链、分片、轻客户端

## 使用说明
直接编译运行任意模块即可，核心入口为 blockchain_core.go。
