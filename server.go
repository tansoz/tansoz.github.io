package main

import "net"

func main(){


	// 监听某端口建立服务
	server,err := net.Listen("tcp","0.0.0.0:1515")
	if err != nil{
		panic(err)
	}
	println("正在运行...")

	// 循环接受客户端
	for{

		// 接受客户端
		if client,err := server.Accept();err == nil{

			// 逻辑部分 协程
			go func(client net.Conn) {

				// 接收客户端的数据
				var b = make([]byte,1024)
				for{

					num,err := client.Read(b)
					print(string(b[0:num]))
					if err != nil || num < 1024 {

						break
					}
				}

				// 发送数据给客户端
				client.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/html;charset=utf-8\r\nServer: Hello/0.1\r\n\r\nHello World!<br>你好,世界!"))

				// 结束 关闭client的socket
				client.Close()

			}(client)
		}

	}
}
