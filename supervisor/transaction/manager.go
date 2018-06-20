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

func (self *Manager) EnlistAction() {

}

func (self *Manager) DelistAction() {
}

func (self *Manager) CreateTransactionID() (int, error) {
	return 0, nil
}
