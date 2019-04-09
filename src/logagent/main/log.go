package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLevel(level string) int{
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogger()(err error){
	config := make(map[string]interface{})
	config["filename"] = appconfig.log_path
	config["level"] = convertLevel(appconfig.log_level)
	configStr ,err := json.Marshal(config)
	if err != nil{
		fmt.Printf("json.Marshal failed err %v",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))
	//logs.Debug("this is a test,my name is %s", "stu111")
	//logs.Trace("this is a trace,my name is %s", "stu02")
	//logs.Warn("this is a warn,my name is %s", "stu03")
	return
}