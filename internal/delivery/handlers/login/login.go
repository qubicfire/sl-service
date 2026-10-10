package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qubicfire/sl-service/internal/repository"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := repo.GetUserByEmail(loginRequest.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, user.ID, user.Name, user.Email, user.Password, user.CreatedAt)

		//login_system.validate_user(login, passowrd)
		// И вот уже внути login_system будет использоваться репозиторий,
		// который является интерфейсом для данных о пользоветеле
		// Данный репозиторий может быть любым (InMemory, SQL, Mongo, etc)
		// Главное чтобы они все следовали интерфейсу
		// Первым делом определяешь что тебе нужно от интерфейся (общие функции, что хочешь получить)
		// После создаешь сам интерфейс
		// Дальше пишешь реализацию
		// И почти всегда это касается систем, например.
		// На это построен, например TDD.
		// Также можешь на хабре чекнуть что такое DDD

		// А откуда в данном случае мне получать условный login_system?
		// Разве я могу в wi

		// В таком случае возникает вопрос, где будет храниться и доставаться
		// условное хранилище? параметр? хде.Создаем в main, передаем в хенделры через аргументы функции
		// Нротт ттут надо подумать. я ЕЮАЛ ГОШКУ. ТРУ
	}
}
