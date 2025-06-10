# API
## Роуты

- `/worker/autocomplete`:
```typescript
const result = await pb.send('/worker/autocomplete', {
  method: 'GET',
  query: {
    query: 'прога'
  }
});
```
Ответ: 
```json5
{
  "query": "прога", // исходный запрос
  "response": {
    "count": 5, // количество ответов (ограничено в 5)
    "matches": [
      {
        "id": "1ee1lj1f647xww8", // id записи интереса в Pocketbase
        "metadata": {
          "tag": "Программирование" // словесное название интереса
        },
        "score": 0.9374236 // Евклидова метрика похожести
      },
      {
        "id": "i9qql84d0b9oyuu",
        "metadata": {
          "tag": "Рисование"
        },
        "score": 1.0360423
      },
      ...
    ]
  }
}
```

- `/worker/feed`:
```typescript
const result = await pb.send('/worker/feed', {
  method: 'GET',
});
```
не нужны параметры, возвращает несколько анкет со всеми данными
