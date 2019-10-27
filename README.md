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

## План работ

- Закончить перенос функциональной части сервера, с предыдущих ДЗ
- Закончить перенос функциональной части клиента, с предыдущих ДЗ
- Переделать структуру сервисов get_notify и send_notify, что бы можно было использовать общий код(например модель Event )
- Доработать сервис get_notify
    - Сделать выбор событий в определенном временно промежутке, что бы отправлялись актуальные уведомления\события
    - Сделать обновление событий в БД(поле notify), если уведомление было отправлено
    - Настроить обработку ошибок от send_notify
- Доработать сервис send_notify
    - Сделать случайную ошибку при отправке сообщения пользователю
- ТЕСТЫ
