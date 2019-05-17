# Go语言学习笔记

## day01 

入门 goroutine 等等

## day02 

函数 包 常量 值类型 数据类型 操作符

## day03 

strings 和 strconv使用

go中的时间和日期类型

流程控制

函数详解

## day04

内置函数、递归函数、闭包

数组与切片

map数据结构

package介绍

## day05

###  结构体和方法

golang中的struct没有构造函数，一般可以使用工厂模式来解决这个问题

接口

## day06

接口

多态

反射

## day07

终端读写

文件

错误处理

## day08

## day09

redis

## day10

http

mysql

## day11

日志收集系统：问题排查，数据挖掘

Elasticsearch+Logstash+Kibana 就是传说中的ELK了，应该是现在最流行的日志处理平台。

https://www.elastic.co/cn/products/elasticsearch

https://www.elastic.co/cn/products/logstash

https://www.elastic.co/cn/products/kibana

## day12

### etcd介绍和使用

概念：高可用的分布式key-value存储，可以用于配置共享和服务发现

类似项目：zookeeper和consul

开发语言：Go

接口：提供Restful的http接口，使用简单

实现算法：基于raft算法的强一致性，高可用的服务存储目录

### etcd的应用场景

- 服务发现和服务注册

- 配置中心

- 分布式锁

- master选举

https://github.com/etcd-io/etcd/releases/

## day13

elastic search

搜索引擎

kibana:

Kibana 是通向 Elastic 产品集的窗口。 它可以在 Elasticsearch 中对数据进行视觉探索和实时分析。 此视频非常适合 Kibana 的新手用户，如果您正在寻找关于数据探索，可视化和仪表盘的初级读本。 看看吴斌如何在几分钟内从 Kibana 安装到完整仪表板建设的演示。 



elastic search介绍和使用

context

## day14

秒杀系统

- 系统解耦

- 过载保护

- 降级预案

![系统设计](https://gitee.com/gdhu/prvpic/raw/master/2019-05-15-001.jpg)

redis windows版本，已经不更新了

https://github.com/MicrosoftArchive/redis/releases

:bear::whale::dolphin:

