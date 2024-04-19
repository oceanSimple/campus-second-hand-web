package model

type Server struct {
	Connections map[uint64]*Connection // 保存所有连接
	DelConn     chan uint64            // 根据发送的userId删除连接
}

func (server *Server) Run() {
	// 循环监听channel
	for {
		select {
		case userId := <-server.DelConn:
			// 删除连接(首先判断连接是否存在，若关闭不存在的channel会引发panic)
			if v, ok := server.Connections[userId]; ok {
				delete(server.Connections, userId)
				close(v.Sc)
			}
		}
	}
}
