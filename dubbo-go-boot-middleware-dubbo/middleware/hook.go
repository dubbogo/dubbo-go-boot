package middleware

type DubboSetupHook struct {
	hook func()
}

func NewDubboSetupHook(hook func()) *DubboSetupHook {
	return &DubboSetupHook{
		hook: hook,
	}
}

func (m *DubboSetupHook) Hook() {
	if m.hook != nil {
		m.hook()
	}
}
