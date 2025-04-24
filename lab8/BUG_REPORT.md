- Пограничное максимальное значение category_id > 15 
Шаги воспроизведения
Создать или редактировать продукт с category_id = 16
Ожидаемый результат: 200 статус со страницей ошибки
Фактический результат: продукт сохраняется с category_id = 16

- Нет строго проверки для статуса и хита на значения 0, 1
Шаги воспроизведения создать/редактировать продукт с такими кейсами
```json
{
  "test_cases": {
    "invalid_hit_above_range": {
      "category_id": "14",
      "title": "Invalid Hit Above Range",
      "content": "This is a test product.",
      "price": "100",
      "old_price": "120",
      "status": "1",
      "keywords": "test, product",
      "description": "123",
      "hit": "2"
    },
    "invalid_hit_below_range": {
      "category_id": "14",
      "title": "Invalid Hit Below Range",
      "content": "This is a test product.",
      "price": "100",
      "old_price": "120",
      "status": "1",
      "keywords": "test, product",
      "description": "123",
      "hit": "-1"
    },
    "invalid_hit_type": {
      "category_id": "14",
      "title": "Invalid Hit Type",
      "content": "This is a test product.",
      "price": "100",
      "old_price": "120",
      "status": "1",
      "keywords": "test, product",
      "description": "123",
      "hit": "eee"
    },
    "invalid_status_above_range": {
      "category_id": "14",
      "title": "Invalid Status Above Range",
      "content": "This is a test product.",
      "price": "100",
      "old_price": "120",
      "status": "2",
      "keywords": "test, product",
      "description": "123",
      "hit": "1"
    },
    "invalid_status_below_range": {
      "category_id": "14",
      "title": "Invalid Status Below Range",
      "content": "This is a test product.",
      "price": "100",
      "old_price": "120",
      "status": "-1",
      "keywords": "test, product",
      "description": "123",
      "hit": "1"
    },
    "invalid_status_type": {
      "category_id": "14",
      "title": "Invalid Status Type",
      "content": "This is a test product.",
      "price": "100",
      "old_price": "120",
      "status": "ewweew",
      "keywords": "test, product",
      "description": "123",
      "hit": "1"
    }
  }
}
```
Ожидаемый результат: 200 статус со страницей ошибки
Фактический результат: продукт сохраняется с неверным хитом и статусом
