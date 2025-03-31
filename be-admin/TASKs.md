# Team Tasks Checklist

## Authentication & Authorization
- [x] 🟢🔴 **POST /login** - User login
- [x] 🟢 **POST /api/register** - User registration

## Users API
- [x] 🔴 **GET /users** - List users
- [x] 🔴 **GET /users/{id}** - Get user details
- [x] 🔴 **POST /users** - Create a new user
- [x] 🔴 **PUT /users/{id}** - Update user details
- [x] 🔴 **DELETE /users/{id}** - Soft delete user
- [x] 🔴 **GET /user/profile** - Get user profile
- [x] 🔴 **PUT /user/profile** - Update user profile
- [x] 🔴 **PUT /user/{id}/active** - Activate/deactivate user


## Comics
- [x] 🟢 **GET /comics** - List comics
- [x] 🟢 **GET /comics/{id}** - Get comic details
- [x] 🟢 **POST /comics** - Create a new comic
- [x] 🟢 **PUT /comics/** - Update comic
- [x] 🟢 **DELETE /comics/{id}** - Delete comic 
- [x] 🟢 **POST /comics/upload-cover** - Upload comic cover

## Genres
- [x] 🔴 **GET /genre** - List genres
- [x] 🔴 **GET /genre/{id}** - Get genre details
- [x] 🔴 **POST /genre** - Create a new genre
- [x] 🔴 **PUT /genre/{id}** - Update genre
- [x] 🔴 **DELETE /genre/{id}** - Delete genre

## Authors
- [x] 🟡 **GET /authors** - List authors
- [x] 🟡 **GET /authors/{id}** - Get author details
- [x] 🟡 **POST /authors** - Create a new author
- [x] 🟡 **PUT /authors/{id}** - Update author
- [x] 🟡 **DELETE /authors/{id}** - Delete author

## Chapters
- [x] 🟢 **GET /chapters** - List chapters
- [x] 🟢 **GET /chapters/{id}** - Get chapter details
- [x] 🟢 **POST /comics/{comic_id}/chapters** - Add new chapter to comic
- [x] 🟢 **PUT /chapters/{id}** - Update chapter
- [x] 🟢 **DELETE /chapters/{id}** - Delete chapter

## Chapter Items
- [x] 🔵 **GET /chapter-items** - List item of a chapter
- [x] 🔵 **GET /chapter-items/{id}** - Get item details
- [x] 🔵 **POST /chapter-items** - Add item to chapter
- [x] 🔵 **PUT /chapter-items/{id}** - Update item details
- [x] 🔵 **DELETE /chapter-items/{id}** - Delete item
- [x] 🔵 **POST /chapter-items/upload-image** - Upload chapter page image

## Recommendations
- [x] 🟡 **GET /recommend** - List recommendations
- [x] 🟡 **GET /recommend/{id}** - Get recommendation details
- [x] 🟡 **POST /recommend** - Create a recommendation
- [x] 🟡 **PUT /recommend/** - Update recommendation
- [x] 🟡 **DELETE /recommend/{id}** - Delete recommendation

## Ads
- [x] 🔵 **GET /ads** - List ads
- [x] 🔵 **GET /ads/{id}** - Get ad details
- [x] 🔵 **POST /ads** - Create an ad
- [x] 🔵 **PUT /ads/{id}** - Update ad
- [x] 🔵 **DELETE /ads/{id}** - Delete ad

## General
- [x] 🔴 **GET /general/tiers** - Get tiers
- [x] 🔴 **GET /general/genres** - Get genres
- [x] 🔴 **GET /general/authors** - Get authors
- [x] 🟢 **GET /general/comics** - Get comics
- [ ] 🔴 **GET /general/genre-for-short-drama** - Get genres for short drama

---

# Short Drama API

## Drama API
- [ ] 🟢 **GET /dramas** - List dramas
- [ ] 🟢 **GET /dramas/{id}** - Get drama details
- [ ] 🟢 **POST /dramas** - Create a new drama
- [ ] 🟢 **PUT /dramas/{id}** - Update drama
- [ ] 🟢 **DELETE /dramas/{id}** - Delete drama
- [ ] 🟢 **PUT /dramas/{id}/active** - Activate/deactivate drama

### Request
```json
{
  "release_date": "2023-11-15",
  "thumb": "http://server.com/thumb.jpg",
  "translations": [
    {
      "language_code": "en",
      "title": "Love in Tokyo",
      "description": "A romantic story set in Tokyo"
    },
    {
      "language_code": "es",
      "title": "Amor en Tokio",
      "description": "Una historia romántica ambientada en Tokio"
    }
  ],
  "genres": [1, 3, 5]
}
```

## Episodes API
- [ ] 🟡 **GET /episodes** - List episodes
- [ ] 🟡 **GET /episodes/{id}** - Get episode details
- [ ] 🟡 **POST /episodes** - Add new episode to drama
- [ ] 🟡 **PUT /episodes/** - Update episode
- [ ] 🟡 **DELETE /episodes/{id}** - Delete episode
- [ ] 🟡 **PUT /episodes/{id}/active** - Activate/deactivate episode

### Request
```json
{
  "number": 1,
  "drama_id": 123,
  "video": "https://storage.com/dramas/123/ep1.mp4",
  "subtitles": [
    {
      "language_code": "en",
      "url": "https://storage.com/dramas/123/ep1-en.srt"
    },
    {
      "language_code": "es",
      "url": "https://storage.com/dramas/123/ep1-es.srt"
    }
  ]
}
```

## Genre API
- [ ] 🔵 **GET /genres** - List genres
- [ ] 🔵 **GET /genres/{id}** - Get genre details
- [ ] 🔵 **POST /genres** - Create a new genre
- [ ] 🔵 **PUT /genres/** - Update genre
- [ ] 🔵 **DELETE /genres/{id}** - Delete genre

### Request
```json
{
  "name": "Romance",
  "position": 1,
  "translations": [
    {
      "language_code": "en",
      "name": "Romance"
    },
    {
      "language_code": "es",
      "name": "Romance"
    }
  ]
}
```

## Uploads API
- [ ] 🔴 **POST /uploads/episode-video** - Upload episode video
- [ ] 🔴 **POST /uploads/episode-subtitle** - Upload episode subtitle
- [ ] 🔴 **POST /uploads/drama-thumb** - Upload drama thumb


### Team Members and Their Assigned Colors

indicate member's respective tasks 

| member   | color | 
|:---------|:----:| 
| Dang     | 🟢 | 
| NinhKieu | 🔴 | 
| Zachdemp | 🔵 | 
| Anhmien  | 🟡 |



