---
title: "电信网管微服务平台"
abstract: "容器化管理，业务间隔离；基于微服务架构；深度集成 DevOps；根据业务负载量，快速动态扩缩容；支持租户管理和租户隔离；提供 API 网关，支持对外开放 API"
image: "/img/solution/microservice/architecture.png"
---

## 案例简介
为某电信公司构建下一代网络管理基础平台。

客户希望该平台作为未来网管的 PaaS 平台，不同厂商的产品都能以微服务的形式在上面部署和管理。

客户基本需求如下：

- 平台必须采用微服务架构；
- 平台需要支持租户管理，各厂商能以租户身份部署自己的产品；
- 租户和租户之间完全隔离，但可以通过 API 相互通讯；
- 平台对外能提供开放 API，内部任何业务功能可以以开放 API 的形式对外曝露；
- 平台需要提供 API 网关，过滤非目标客户的请求和目标客户的不正常访问，安全策略能随时调整；
- 平台根据不同当前业务负载量，能快速动态扩缩容，不中断业务；
- 平台能充分使用当前的服务器和网络资源，效率最大化；
- 平台支持 DevOps，能提高公司自有研发部门的产品研发、运维的迭代效率；

## 硬件环境
10台 KVM 虚拟机，每台配置：

- CPU：8核心；
- Ram：32G；
- Disk：200G；
- NIC：1000M * 1；

## 软件环境
- CentOS 7.2 64Bit；
- Oracle JDK 1.8 64Bit；

## 技术方案
### 1. PaaS 架构
PaaS 平台整体架构如下图所示：

![微服务网络模型](/img/case/microservice/paas.png "Microservice")

由于客户有自己的 KVM 虚拟机集群，只需要在虚拟层之上构建微服务和 DevOps 平台。

### 2. 微服务架构
微服务平台的基本方案如下：

- 资源管理基于 Docker + Kubernetes；
- Spring Cloud 作为微服务框架；
- Zabbix + InfluxDB + Grafana 实现平台的性能监控；

Kubernetes 网络模型基于 Flannel 叠加网模型（overlay network），映射方式如下图所示：

![微服务网络模型](/img/case/microservice/network-model.png "Microservice")

### 3. DevOps 架构
DevOps 平台基于 GitLab + Jenkins + Maven Server + Swagger 搭建。

各组件作用如下：

- GitLab： 管理源代码和各微服务应用的配置文件；
- Jenkins：
  + 从 GitLab 拉取源码；
  + 编译和构建微服务应用；
  + 将构建好的微服务应用打包成 Docker 镜像；
  + 推送 Docker 镜像 至 Docker Registry；
- Maven Server：为编译和构建过程提供依赖包；
- Swagger: 生成 Restful API 文档；

### 4. 系统集成
系统最终的整体架构如下图所示：

![微服务架构](/img/solution/microservice/architecture.png "Microservice")

## 交付周期
- 咨询：1个月。
- 实施：3个月。
