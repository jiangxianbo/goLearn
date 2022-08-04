package main

import (
	"gopkg.in/ini.v1"
	"learn/log_transfer/conf"
	"learn/log_transfer/es"
	"learn/log_transfer/kafka"
)

// log transfer
// 将日志数据取出发往es
var (
	cfg = new(conf.LogTransfer)
)

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		panic(err)
	}
	// 1.初始化es
	// 1.1 初始化ES链接的一个client
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		panic(err)
	}

	// 2.初始化kafka
	// 2.1 链接kafka，创建分区消费者
	// 2.2 每个分区的消费者分别取出数据，通过SendES()将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		panic(err)
	}

	select {}
}
