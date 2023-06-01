# My_RPC
A simple attempt at RPC



效果演示：

Server：go run .\Server.go

~~~shell
2023/06/01 09:50:17 Serving RPC server on port 1234
2023/06/01 09:50:24 New connection established
2023/06/01 09:50:24 Multiplying 7 with 8
2023/06/01 09:50:25 Connection closed
~~~

Client:go run .\Client.go

~~~shell
2023/06/01 09:50:24 Arith: 7*8=56
~~~

