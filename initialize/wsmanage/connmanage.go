package wsmanage

import (
	"errors"
	"fmt"
	"sync"
)

/*
连接管理模块
*/
type ConnManager struct {
	connections map[uint64]Connection //管理的连接信息
	connLock    sync.RWMutex          //读写连接的读写锁
	connList    map[int64]Connection  //用户连接管理
}

/*
创建一个链接管理
*/
func NewConnManager() ConnManager {
	return ConnManager{

		connList:    make(map[int64]Connection),
		connections: make(map[uint64]Connection),
	}
}

// 添加用户ID和链接信息
func (connMgr *ConnManager) AddConn(id int64, conn Connection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn连接添加到ConnMananger中
	connMgr.connList[id] = conn

	fmt.Println("connection add to ConnManager successfully: conn num = ", connMgr.Len())
}

// 删除连接
func (connMgr *ConnManager) RemoveConn(id int64, conn Connection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除连接信息
	delete(connMgr.connList, id)

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", connMgr.Len())
}

// 利用ConnID获取链接
func (connMgr *ConnManager) GetConn(id int64) (*Connection, error) {
	//保护共享资源Map 加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connList[id]; ok {
		return &conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// 获取当前连接
func (connMgr *ConnManager) LenConn() int {
	return len(connMgr.connList)
}

// 添加链接
func (connMgr *ConnManager) Add(conn Connection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn连接添加到ConnMananger中
	connMgr.connections[conn.GetConnID()] = conn

	fmt.Println("connection add to ConnManager successfully: conn num = ", connMgr.Len())
}

// 删除连接
func (connMgr *ConnManager) Remove(conn Connection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除连接信息
	delete(connMgr.connections, conn.GetConnID())

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", connMgr.Len())
}

// 利用ConnID获取链接
func (connMgr *ConnManager) Get(connID uint64) (*Connection, error) {
	//保护共享资源Map 加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return &conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// 获取当前连接
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// 清除并停止所有连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range connMgr.connections {
		//停止
		conn.Close()
		//删除
		delete(connMgr.connections, connID)
	}
	for userID, conn := range connMgr.connList {
		//停止
		conn.Close()
		//删除
		delete(connMgr.connList, userID)
	}

	fmt.Println("Clear All connections successfully: conn num = ", connMgr.Len())
}

//func (connMgr *ConnManager) PushAll(msg []byte) {
//	for _, conn := range connMgr.connections {
//		conn.SendMessage(1,msg)
//	}
//}
