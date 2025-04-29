
# 🎬 Dvd Rental API

API-сервис для работы с открытой базой данных **dvd_rental**.  
Позволяет выполнять CRUD-операции над сущностями: фильмы, актёры, аренды и другие.

---

## 📌 Основные возможности

- Получение списка фильмов
- Поиск актёров по имени
- Регистрация аренды
- Фильтрация, пагинация, поиск
- Swagger-документация (генерация через `swag`)
- Поддержка ElasticSearch для поиска
- Очередь RabbitMQ
- Redis для кеширования

---

## 🛠️ Технологии

- **Язык:** Go (Gin)
- **База данных:** PostgreSQL (dvd_rental)
- **ORM:** GORM
- **Поиск:** ElasticSearch
- **Кеш:** Redis
- **Очередь:** RabbitMQ
- **Документация:** Swagger

---

## 🚀 Запуск проекта

> Убедитесь, что установлены `docker` и `docker-compose`.

```bash
# Клонируем репозиторий
git clone https://github.com/your-username/dvd_rental_api.git
cd dvd_rental_api

# Генерация Swagger-документации
swag init

# Сборка и запуск контейнеров
docker-compose up -d --build

# Остановка контейнеров и удаление томов
docker-compose down -v
```

### 🔧 Подготовка окружения (Ubuntu 24.04.2)

```bash
sudo apt update
sudo apt install docker.io docker-compose python3-setuptools

# Добавить пользователя в группу docker (перезапустить сессию после)
sudo usermod -aG docker $USER

# (Если установлен Apache — удалить)
sudo service apache2 stop
sudo apt-get purge apache2 apache2-utils apache2.2-bin
sudo rm -rf /etc/apache2
```

---

## 📄 Swagger-документация

Доступна по адресу:  
`http://localhost:8080/swagger/index.html`

---

## 📇 Метаданные (OpenAPI)

```go
// @title           Dvd Rental API
// @version         1.0
// @description     This is an API for working with the public database **dvd_rental**.
// @description     Supports CRUD operations on movies, actors, rentals, etc.
// @termsOfService  http://example.com/terms/

// @contact.name   Roman S.
// @contact.url    https://www.linkedin.com/in/roman-s-bba6021a5/
// @contact.email  unpredictableanonymous639@gmail.com

// @license.name   GPL-3.0 license
// @license.url    https://www.gnu.org/licenses/gpl-3.0.html

// @host           localhost:8080
// @BasePath       /
```

---
