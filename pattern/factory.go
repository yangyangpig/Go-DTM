package pattern

type Transaction interface {
	Commit()
	Rollback()
	EnlistAction()
	DelistAction()
}

type Recover interface {
}

type Context interface {
}
