package dblink

type mysql struct {
}

func NewdbBaseMysql() *mysql {
	return new(mysql)
}
