ИТОГОВЫЙ ПРОЕКТ:

-------Приложение представляет собой простейший планировщик задач и включает следующие возможности:
1.	Аутентификация: система использует пароли и JWT-токены для защиты доступа пользователей.
2.	REST API: взаимодействие с приложением через стандартизированные API-запросы.
3.	Управление задачами: пользователи могут создавать, изменять и удалять задачи, в т.ч. выставлять задачи, которые будут регулярно повторяться.
4.	Хранение данных: используется база данных SQLite для сохранения информации о задачах и пользователях.
5. Добавлена возможность задавать порт сервера через переменную окружения TODO_PORT при запуске.
6. Добавлена возможность указывать путь к файлу базы данных через переменную окружения TODO_DBFILE.
7. Для запуска тестов с переменными окружения используйте TODO_PORT=7540 go test ./tests в bash и zsh, или set TODO_PORT=7540 && go test ./tests в cmd.
8.Реализована обработка параметров поиска в маршруте Get("/api/tasks").
TODO_PASSWORD=1234

Запуск приложения:
с  Docker:
docker build -t todoserver:v1 .
docker run -d --name todo -p 7540:7540 -e TODO_DBFILE=scheduler.db -e TODO_PASSWORD=1234 -v scheduler.db todoserver:v1


без  Docker:
go mod tidy
go build .
go run .

Остановка приложения:
Ctrl+C

Тестирование:
go test ./tests -v

Перед повторным тестированием: 
go clean -testcache
Можно так же запускать:
go test ./tests -v -count=1

Для запуска приложения  необходимо убедиться, что в вашей системе установлена версия Go не ниже 1.23.2.