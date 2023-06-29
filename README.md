# SberDS

Тестовое задание в sberDS

#### Локальный запуск

1. Создать файл ".env" с необходимыми переменными окружения рядом с файлом "Makefile"

   *.env example*

   ```
    DATABASE_HOST=localhost  
    DATABASE_PORT=5432
    DATABASE_NAME=sandbox
    DATABASE_USERNAME=postgres
    DATABASE_PASSWORD=postgres123
   
    SERVER_HOST=localhost
    SERVER_PORT=8080
    SERVER_WRITE_TIMEOUT=10s
    SERVER_READ_TIMEOUT=10s
   ```

2.  Выполнить make run в консоли

#### Запуск с помощью Docker-compose

1. В docker-compose.yml укажите все необходимые переменные окружения
2. Выполните команду docker-compose up

