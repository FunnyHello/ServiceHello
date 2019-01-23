package session

type Provider interface {
	SessionInit(sid string)(Session,error)
	SessionRead(sid string)(Session,error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var provides = make(map[string]Provider)

func Register(name string,provider Provider)  {
	if provider == nil {
		panic("session: 注册提供者为空")
	}
	if _,dup := provides[name];dup {
		panic("session: 为提供程序调用了两次注册 " + name)
	}
	provides[name] = provider
}