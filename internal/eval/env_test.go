package eval

import (
	"testing"
)

func TestNewEnv(t *testing.T) {
	parent := NewEnv(nil)
	env := NewEnv(parent)

	if env.parent != parent {
		t.Errorf("expected parent to be %v, got %v", parent, env.parent)
	}

	if len(env.vars) != 0 {
		t.Errorf("expected vars to be empty, got %v", env.vars)
	}
}

func TestEnv_SetAndGet(t *testing.T) {
	env := NewEnv(nil)
	env.Set("foo", "bar")

	val := env.Get("foo")
	if val != "bar" {
		t.Errorf("expected value to be 'bar', got %v", val)
	}
}

func TestEnv_GetFromParent(t *testing.T) {
	parent := NewEnv(nil)
	parent.Set("foo", "bar")

	env := NewEnv(parent)
	val := env.Get("foo")

	if val != "bar" {
		t.Errorf("expected value to be 'bar', got %v", val)
	}
}

func TestEnv_GetNonExistentKey(t *testing.T) {
	env := NewEnv(nil)
	val := env.Get("nonexistent")

	if val != nil {
		t.Errorf("expected value to be nil, got %v", val)
	}
}
