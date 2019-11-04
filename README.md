# otus_calendar

## Запуск проекта
### Обычный запуск
```bash
docker-compose build --no-cache
docker-compose up
```
### С запуском "условного" клиента
```bash
docker-compose -f docker-compose.yml -f docker-compose.local.yml build --no-cache
docker-compose -f docker-compose.yml -f docker-compose.local.yml up
```

### С интеграционным тестированием
```bash
docker-compose -f docker-compose.yml -f docker-compose.test.yml build --no-cache
docker-compose -f docker-compose.yml -f docker-compose.test.yml up
```

