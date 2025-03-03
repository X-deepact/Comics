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
- [x] 🟢 **GET /comic/{id}** - Get comic details
- [x] 🟢 **POST /comic** - Create a new comic
- [x] 🟢 **PUT /comic/** - Update comic
- [x] 🟢 **DELETE /comic/{id}** - Delete comic 
- [x] 🟢 **POST /comic/upload-cover** - Upload comic cover

## Genres
- [x] 🔴 **GET /genres** - List genres
- [x] 🔴 **GET /genres/{id}** - Get genre details
- [x] 🔴 **POST /genres** - Create a new genre
- [x] 🔴 **PUT /genres/{id}** - Update genre
- [x] 🔴 **DELETE /genres/{id}** - Delete genre

## Authors
- [x] 🟡 **GET /authors** - List authors
- [x] 🟡 **GET /authors/{id}** - Get author details
- [ ] 🟡 **GET /authors/{id}/comics** - List comics of an author
- [x] 🟡 **POST /authors** - Create a new author
- [ ] 🟡 **PUT /authors/{id}** - Update author
- [x] 🟡 **DELETE /authors/{id}** - Delete author

## Chapters
- [x] 🟢 **GET /comics/{comic_id}/chapters** - List chapters of a comic
- [x] 🟢 **GET /chapters/{id}** - Get chapter details
- [x] 🟢 **POST /comics/{comic_id}/chapters** - Add new chapter to comic
- [x] 🟢 **PUT /chapters/{id}** - Update chapter
- [x] 🟢 **DELETE /chapters/{id}** - Delete chapter
- [x] 🟢 **POST /chapters/upload-cover** - Upload chapter cover

## Chapter Items
- [x] 🔵 **GET /chapters/{chapter_id}/items** - List pages of a chapter
- [x] 🔵 **GET /chapter-items/{id}** - Get page details
- [x] 🔵 **POST /chapters/{chapter_id}/items** - Add page to chapter
- [x] 🔵 **PUT /chapter-items/{id}** - Update page details
- [x] 🔵 **DELETE /chapter-items/{id}** - Delete page
- [x] 🔵 **POST /chapter-items/upload-image** - Upload chapter page image

## Recommendations
- [x] 🟡 **GET /recommendations** - List recommendations
- [x] 🟡 **GET /recommendations/{id}** - Get recommendation details
- [x] 🟡 **POST /recommendations** - Create a recommendation
- [ ] 🟡 **PUT /recommendations/{id}** - Update recommendation
- [x] 🟡 **DELETE /recommendations/{id}** - Delete recommendation

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

---

### Team Members and Their Assigned Colors

indicate member's respective tasks 

| member   | color | 
|:---------|:----:| 
| Dang     | 🟢 | 
| NinhKieu | 🔴 | 
| Zachdemp | 🔵 | 
| Anhmien  | 🟡 |



