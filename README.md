# 🛒 Маркетплейс API



![marketplave](https://github.com/shuklarituparn/Marketplace-Application/assets/66947051/d6e2f000-0277-4331-933d-87dad46f6e8c)

- API Доступен здесь `https://marketplace.rtprnshukla.ru/docs/`
- Скоро приложение для доступа к этому API будет готово (надеюсь) и будет доступно по адресу `vk-marketplace.rtprnshukla.ru`

Вы можете отслеживать прогресс здесь: [Marketplace-Application](https://github.com/shuklarituparn/Marketplace-Application)

---

## Использование и установка

- [Установка](docs/setup.md)

---


## Обзор

Этот API позволяет вам выполнить следующую задачу

### Пользователи

- Пользователи могут зарегистрироваться, отправив логин и пароль
- Пользователи могут зарегистрироваться, отправив логин и пароль


### Объявления

- Зарегистрированные пользователи могут размещать объявления
- отображение объявлений

### Развертывание

- Документация Swagger доступна на сайте: https://marketplace.rtprnshukla.ru/docs/ (на английском)
- Метрики от Prometheus можете здесь посмотреть https://marketplace.rtprnshukla.ru/metrics/
- АПИ доступен по сайту: https://marketplace.rtprnshukla.ru/  
  ` (Проверка работоспособности отображается, если мы открываем сайт без метода)`
- `Сначала прочтите это` - [Установка](docs/setup.md)

---

## Особенности

- **Мониторинг**: Сервис использует prometheus/grafana для сбора метрик из различных эндпойнты
- **OpenAPI/Swagger**: Сервис использует спецификацию OpenAPI/Swagger для лучшего тестирования и документирования

---

## Технологический стек

- **Бэкенд**: Go (net/http)
- **Базы данных**: GORM с PostgreSQL (локальная)
- **Электронная почта**: resend для отправки alerts из grafana
- **Генерация метрики**: Prometheus
- **Мониторинг**: Grafana для визуализации
- **Аутентификация пользователей**: Использовал Jwt
- **Git Хуки**: Использовал husky
- **Linting**: Использовал golang-ci-lint
- **Hot-Reload**- Использовал air
- **Makefile** - Добавил возможность установки сервиса с помощью makefile вместе с docker
- **Deployment**: Docker, Docker-compose
- **CI/CD**: Github Actions, Gitlab

---
## Измерение и отображение метрик


![Screenshot from 2024-03-28 21-55-44](https://github.com/shuklarituparn/Filmoteka/assets/66947051/0f49e775-e0d7-4ba6-b827-d3e31a3093e6)


> При запросе сервера prometheus в grafana добавьте `https://prometheus:9090`


---
