package commons

//客户端服务器端数据交互模版
type EgoResult struct{
    Status int//状态 200 表示成功
    Date   interface{}//返回数据
    Msg    string//返回的消息
}