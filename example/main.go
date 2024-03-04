package main

import (
	"log"
	"time"

	opc "github.com/mrzwzw/OPC-DA"
	"github.com/spf13/viper"
)

var v *viper.Viper

func InitConfig() *viper.Viper {
	v := viper.New()

	v.SetConfigType("toml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件异常, Error:", err.Error())
	}

	return v
}

func main() {
	v = InitConfig()
	log.Println("配置文件初始化完成")

	tt := v.GetStringSlice("TAGS.TAG")
	server := v.GetString("TAGS.SERVER")

	client, err := opc.NewConnection(server, []string{"localhost"}, tt)
	if err != nil {
		log.Println("err:", err)
	}

	defer client.Close()

	for {
		mm := client.Read()

		for key, v := range mm {
			if v.Good() {
				if v.Type >= 8192 {
					log.Printf("Item ID: %s, value: %v\n", key, v.SliceValue)
				} else {
					log.Printf("Item ID: %s, value: %v\n", key, v.Value)
				}
			} else {
				log.Printf("Item ID: %s, value is bad \n", key)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
