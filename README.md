# Go 中的 OPC DA

DA从 OPC 服务器在 Go 中获取和写入流程和自动化数据，用于监控和数据分析目的（OPC DA 协议）。



# 编译运行前准备

**1**、下载opc核心包 ``OPC Core Components Redistributable (x86) 3.00.108.msi``并安装

**2**、在``C:\Windows\SysWOW64``安装``OPCDAAuto.dll``，以管理员身份执行``Regsvr32 OPCDAAuto.dll``

**3**、运行时环境搭建，GOARCH=386



# 安装

``go get github.com/mrzwzw/opc``



## 代码示例 

```go
package main

import (
	"fmt"
	"github.com/konimarti/opc"
)

func main() {
	client, _ := opc.NewConnection(
		"opcserversim.Instance.1", // ProgId
		[]string{"localhost"}, //  OPC servers nodes
		[]string{"numeric.sin.int64", "numeric.saw.float"}, // slice of OPC tags
	)
	defer client.Close()

	// read single tag: value, quality, timestamp
	fmt.Println(client.ReadItem("numeric.sin.int64"))

	// read all added tags
	fmt.Println(client.Read())
}
```





# OPC DA数据类型

 

| 值(十进制) | 类型        | 描述                       |
| :--------- | ----------- | -------------------------- |
| 0          | VT_EMPTY    | 默认/空（无）              |
| 2          | VT_I2       | 2字节有符号整数            |
| 3          | VT_I4       | 4字节有符号整数            |
| 4          | VT_R4       | 4字节实数                  |
| 5          | VT_R8       | 8字节实数                  |
| 6          | VT_C        | currency                   |
| 7          | VT_DATE     | 日期                       |
| 8          | VT_BSTR文本 | 文本                       |
| 10         | VT_ERROR    | 错误代码                   |
| 11         | VT_BOOL     | 布尔值（TRUE=-1，FALSE=0） |
| 17         | VT_I1       | 1个字节有符号字符          |
| 18         | VT_UI1      | 1个字节无符号字符          |
| 19         | VT_UI2      | 2字节无符号整数            |
| 20         | VT_UI4      | 4字节无符号整数            |
| +8192      | VT_ARRAY    | 值数组(即8200=文本值数组)  |


