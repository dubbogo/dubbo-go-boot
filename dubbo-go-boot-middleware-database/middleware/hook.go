package middleware

type DatabaseSetupHook struct {
	hook func()
}

func NewDatabaseSetupHook(hook func()) *DatabaseSetupHook {
	return &DatabaseSetupHook{
		hook: hook,
	}
}

func (m *DatabaseSetupHook) Hook() {
	if m.hook != nil {
		m.hook()
	}
}
