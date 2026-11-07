package login_test

import (
	"testing"

	"github.com/qubicfire/sl-service/internal/repository"
)

func TestAddUserMemory(t *testing.T) {
	r := repository.NewMemoryRepository()

	_, err := r.CreateUser("Alexandro", "just_developer@yandex.ru", "fuck_arch_linux")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	user, err := r.GetUserByEmail("just_developer@yandex.ru")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	t.Log(user.Name, user.Email)
}
