package main

import "fmt"

func main() {

	// initConfig
	err := initConfig("ini", "D:\\project\\src\\go_dev\\day13\\log_transfer\\logtrans.conf")
	if err != nil {
		panic(err)
	}

	fmt.Println(logConfig)

	// // initLog
	// err = initLogger()
	// if err != nil {
	// 	panic(err)
	// }
	// // initKafkaConsuemr
	// err = initKafka()
	// if err != nil {
	// 	logs.Error("init kafka err: %v", err)
	// 	return
	// }
	// // initES
	// err = initES()
	// if err != nil {
	// 	logs.Error("init ES err: %v", err)
	// 	return
	// }
	// // runServer: kafka consuemr send message to ES, then Kibana
	// err = run()
	// if err != nil {
	// 	logs.Error("run failed, err: %v", err)
	// 	return
	// }
	// logs.Warn("warning, log transfer exited.")
}
