# REST-сервис для агрегации данных об онлайн-подписках

Этот проект представляет собой REST-сервис, который позволяет управлять записями о подписках пользователей. Сервис поддерживает CRUDL-операции и может подсчитывать суммарную стоимость подписок за выбранный период.

## Требования

- Golang
- PostgreSQL
- Docker и Docker Compose

## Установка

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/ваш_логин/ваш_репозиторий.git
   cd ваш_репозиторий
   ```

2. **Создайте файл окружения:**
   Создайте файл `.env` в корне проекта и добавьте следующие параметры:
   ```dotenv
   DB_PASSWORD=ваш_пароль
   ```

   DB_USER=ваш_пользователь
   DB_NAME=ваше_имя_базы_данных
   DB_HOST=localhost
   DB_PORT=5432
   Эти параметры менять в config.yaml

4. **Запуск с помощью Docker Compose:**
   Выполните команду для сборки и запуска контейнеров:
   ```bash
   docker-compose up --build
   ```

## API

Сервис предоставляет следующие HTTP-ручки:

### CRUDL-операции

- **Создать подписку**
  - **POST** `/subscriptions`
  - **Тело запроса:**
    ```json
    {
      "service_name": "Yandex Plus",
      "price": 400,
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "start_date": "07-2025"
    }
    ```

- **Получить все подписки**
  - **GET** `/subscriptions`

- **Получить подписку по ID**
  - **GET** `/subscriptions/{id}`

- **Обновить подписку**
  - **POST** `/subscriptions/{id}`
  - **Тело запроса:**
    ```json
    {
      "service_name": "Yandex Plus",
      "price": 400,
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "start_date": "07-2025",

    }
    ```

- **Удалить подписку**
  - **DELETE** `/subscriptions/{id}`

### Подсчет стоимости подписок

- **POST** `/subscriptions/total`
  - **Тело запроса:**
    ```json
    {
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "service_name": "Yandex Plus",
      "start_date": "01-2025",
    }
    ```

## Логирование

Код покрыт логами для отслеживания операций и ошибок.

## Документация

Swagger-документация доступна по адресу: `http://localhost:8080/swagger/index.html` после запуска сервиса.

## Docker

Для загрузки образа с Docker Hub используйте команду:
```bash
docker push ramin1234/forjun-app:tagname
```
