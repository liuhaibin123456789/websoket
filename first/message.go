package main

type Message struct {
	Sender   string `json:"sender"`   //发送者
	Receiver string `json:"receiver"` //接收者
	Content  string `json:"content"`  //内容
}
