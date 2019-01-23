package session

type Session interface {
	//会话接口
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID()string
}