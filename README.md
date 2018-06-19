# Go-DTM
Distributed manager based on golang

这是一个基于golang语言开发的分布式简单服务事务管理系统

## 简介

这个分布式事务系统是对于服务远程调用的事务处理的一种管理，把发起服务调用的服务称为主服务，被调用的服务称为从服务。而本地服务的一致性遵循ACID原则，远程从服务的
调用默认为一个原子性。

目前支持TCC一致模型、补偿模型、不可靠消息模型、可靠消息模型

##架构图如下

![](https://github.com/developersPHP/Go-DTM/blob/master/source/%E7%B3%BB%E7%BB%9F%E6%9E%B6%E6%9E%84%E8%AE%BE%E8%AE%A1.png)
