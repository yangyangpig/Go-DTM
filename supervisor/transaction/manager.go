package transaction

type Manager struct {
}

func Start() *Manager {
	return new(Manager)
}

func (self *Manager) Commit() error {
	return nil
}

func (self *Manager) Rollback() error {
	return nil
}

//用于列出当前原子活动列表
func (self *Manager) EnlistAction() {

}

//用于取消当前原子活动列表
func (self *Manager) DelistAction() {
}
