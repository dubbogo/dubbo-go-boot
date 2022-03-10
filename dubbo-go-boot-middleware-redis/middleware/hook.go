package middleware

type RedisSetupHook struct {
	hook func()
}

func NewRedisSetupHook(hook func()) *RedisSetupHook {
	return &RedisSetupHook{
		hook: hook,
	}
}

func (m *RedisSetupHook) Hook() {
	if m.hook != nil {
		m.hook()
	}
}
