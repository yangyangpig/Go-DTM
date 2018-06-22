package link

//TODO 由于这里的rabbit和redis和mysql几个驱client差异比较大，所以先设置一个空接口,以后抽象一层共性可以放这里
type dbBase interface {
	//在这里要定义数据库所可以提供的操作方法
}
