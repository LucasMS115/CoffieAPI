# Coffee App — High-Level Scratch

---

## Core Entities

### User

| Attribute    | Description                                | Example Value             |
|-------------|--------------------------------------------|---------------------------|
| id          | Unique identifier (UUID)                   | `a1b2c3d4-e5f6-...`       |
| name        | Display name                               | `Lucas`                   |
| email       | Login email, unique                        | `lucas@email.com`         |
| created_at  | Account creation timestamp                 | `2026-04-05T10:00:00Z`    |

### Coffee

| Attribute      | Description                                        | Example Value              |
|---------------|----------------------------------------------------|----------------------------|
| id            | Unique identifier (UUID)                           | `b2c3d4e5-f6a7-...`        |
| name          | Display name of the coffee (usually origin + lot)  | `Ethiopia Yirgacheffe`     |
| brand         | Roaster or producer                                | `Boteco Coffee Roasters`   |
| type          | Primary flavor profile category                    | `fruity`                   |
| flavor_notes  | Free-text tasting notes                            | `blueberry, jasmine, honey`|
| created_at    | Registration timestamp                             | `2026-04-05T10:00:00Z`     |

### Recipe

| Attribute    | Description                                         | Example Value         |
|-------------|-----------------------------------------------------|-----------------------|
| id          | Unique identifier (UUID)                            | `c3d4e5f6-a7b8-...`   |
| user_id     | Owner of the recipe (foreign key → User)            | `a1b2c3d4-e5f6-...`   |
| coffee_id   | Coffee used (foreign key → Coffee)                  | `b2c3d4e5-f6a7-...`   |
| method      | Brewing method name                                 | `V60`                 |
| water_temp  | Water temperature in Celsius                        | `93`                  |
| dose        | Amount of dry coffee grounds in grams               | `18`                  |
| yield       | Final beverage weight in grams (output)             | `300`                 |
| brew_time   | Total brew time in seconds                          | `180`                 |
| description | Free-text notes about the recipe/result             | `Clean and bright cup`|
| created_at  | Creation timestamp                                  | `2026-04-05T10:00:00Z`|
| updated_at  | Last modification timestamp                         | `2026-04-06T08:30:00Z`|

### Rating

| Attribute  | Description                                      | Example Value               |
|-----------|--------------------------------------------------|-----------------------------|
| id        | Unique identifier (UUID)                         | `d4e5f6a7-b8c9-...`         |
| recipe_id | Rated recipe (foreign key → Recipe)              | `c3d4e5f6-a7b8-...`         |
| user_id   | Rater (foreign key → User)                       | `e5f6a7b8-c9d0-...`         |
| score     | Score from 1 to 5                                | `5`                         |
| comment   | Optional review text                             | `Best V60 I've tried!`      |
| created_at| Rating timestamp                                 | `2026-04-05T11:00:00Z`      |

---

## Entity Relationships

```
User 1 ──── N Recipe
Coffee 1 ──── N Recipe
Recipe 1 ──── N Rating
User 1 ──── N Rating
```

- A **User** owns many **Recipes**
- A **Coffee** can be used in many **Recipes**
- A **Recipe** can receive many **Ratings** from different users
- A **User** can rate many **Recipes** (but presumably once per recipe)

---

## API Endpoints

### Auth / Users

| Method | Path                  | Description       |
|--------|-----------------------|-------------------|
| POST   | `/api/users`          | Register          |
| GET    | `/api/users/{id}`     | Get profile       |
| GET    | `/api/users/{id}/stats`| Get user stats   |

### Recipes

| Method | Path                            | Description              |
|--------|---------------------------------|--------------------------|
| POST   | `/api/recipes`                  | Create recipe            |
| GET    | `/api/recipes/{id}`             | Get recipe detail        |
| GET    | `/api/recipes?user_id=X&search=Y`| List/search recipes     |
| PUT    | `/api/recipes/{id}`             | Update recipe (owner)    |
| DELETE | `/api/recipes/{id}`             | Delete recipe (owner)    |

### Coffee

| Method | Path                            | Description           |
|--------|---------------------------------|-----------------------|
| POST   | `/api/coffees`                  | Register a coffee     |
| GET    | `/api/coffees/{id}`             | Get coffee detail     |
| GET    | `/api/coffees?search=X&type=Y`  | Search coffees        |

### Ratings

| Method | Path                            | Description           |
|--------|---------------------------------|-----------------------|
| POST   | `/api/recipes/{id}/ratings`     | Rate a recipe         |
| GET    | `/api/recipes/{id}/ratings`     | List ratings         |

---

## Example Requests & Responses

### Create a Recipe

**Request:**
```
POST /api/recipes
```
```json
{
  "coffee_id": "b2c3d4e5-f6a7-...",
  "method": "V60",
  "water_temp": 93,
  "dose": 18,
  "yield": 300,
  "brew_time": 180,
  "description": "Clean and bright cup"
}
```
*(user_id comes from auth token)*

**Response 201:**
```json
{
  "id": "c3d4e5f6-a7b8-...",
  "user": { "id": "a1b2c3d4-e5f6-...", "name": "Lucas" },
  "coffee": { "id": "b2c3d4e5-f6a7-...", "name": "Ethiopia Yirgacheffe", "brand": "Boteco Coffee" },
  "method": "V60",
  "water_temp": 93,
  "dose": 18,
  "yield": 300,
  "brew_time": 180,
  "description": "Clean and bright cup",
  "avg_rating": null,
  "rating_count": 0,
  "created_at": "2026-04-05T10:00:00Z"
}
```

### Search Recipes

**Request:**
```
GET /api/recipes?method=V60&search=bright
```

**Response 200:**
```json
{
  "items": [
    {
      "id": "c3d4e5f6-a7b8-...",
      "user": { "id": "a1b2c3d4-e5f6-...", "name": "Lucas" },
      "coffee": { "id": "b2c3d4e5-f6a7-...", "name": "Ethiopia Yirgacheffe" },
      "method": "V60",
      "avg_rating": 4.5,
      "rating_count": 12,
      "created_at": "2026-04-01T08:00:00Z"
    }
  ],
  "total": 1,
  "page": 1
}
```

### Rate a Recipe

**Request:**
```
POST /api/recipes/{id}/ratings
```
```json
{
  "score": 5,
  "comment": "Best V60 recipe I've tried!"
}
```

**Response 201:**
```json
{
  "id": "d4e5f6a7-b8c9-...",
  "user": { "id": "e5f6a7b8-c9d0-...", "name": "Ana" },
  "score": 5,
  "comment": "Best V60 recipe I've tried!",
  "created_at": "2026-04-05T11:00:00Z"
}
```

### User Stats

**Request:**
```
GET /api/users/{id}/stats
```

**Response 200:**
```json
{
  "recipes_count": 14,
  "avg_rating_given": 4.2,
  "fav_method": "V60",
  "fav_coffee_type": "fruity"
}
```

---

## Error Response Pattern

```json
{
  "error_code": "INVALID_INPUT",
  "message": "Water temperature must be between 70 and 100 degrees celsius",
  "fields": [
    { "field": "water_temp", "message": "out of range (70-100)" }
  ]
}
```
