# Team Tasks Checklist

## Authentication & Authorization
- [x] 🟢 **POST /api/login** - User login
- [x] 🟢 **POST /api/register** - User registration

## Users API
- [ ] 🔴 **GET /users** - List users
- [ ] 🔴 **GET /users/{id}** - Get user details
- [ ] 🔴 **POST /users** - Create a new user
- [ ] 🔴 **PUT /users/{id}** - Update user details
- [ ] 🔴 **DELETE /users/{id}** - Soft delete user

## Roles & Permissions
- [ ] **GET /roles** - List roles
- [ ] **POST /roles** - Create a new role
- [ ] **PUT /roles/{id}** - Update role
- [ ] **DELETE /roles/{id}** - Delete role
- [ ] **GET /user-roles** - List user roles
- [ ] **POST /user-roles** - Assign role to user

## Comics
- [x] 🟢 **GET /comics** - List comics
- [x] 🟢 **GET /comic/{id}** - Get comic details
- [x] 🟢 **POST /comic** - Create a new comic
- [x] 🟢 **PUT /comic/** - Update comic
- [x] 🟢 **DELETE /comic/{id}** - Delete comic 
- [x] 🟢 **POST /comic/upload-cover** - Upload comic cover

## Genres
- [x] 🔴 **GET /genre** - List genres
- [x] 🔴 **GET /genre/{id}** - Get genre details
- [x] 🔴 **POST /genre** - Create a new genre
- [x] 🔴 **PUT /genres/{id}** - Update genre
- [x] 🔴 **DELETE /genres/{id}** - Delete genre

## Authors
- [ ] 🟡 **GET /authors** - List authors
- [ ] 🟡 **POST /authors** - Create a new author
- [ ] 🟡 **PUT /authors/{id}** - Update author
- [ ] 🟡 **DELETE /authors/{id}** - Delete author

## Chapters
- [ ] 🟢 **GET /comics/{comic_id}/chapters** - List chapters of a comic
- [ ] 🟢 **GET /chapters/{id}** - Get chapter details
- [ ] 🟢 **POST /comics/{comic_id}/chapters** - Add new chapter to comic
- [ ] 🟢 **PUT /chapters/{id}** - Update chapter
- [ ] 🟢 **DELETE /chapters/{id}** - Delete chapter
- [ ] 🟢 **POST /chapters/upload-cover** - Upload chapter cover

## Chapter Items
- [ ] **GET /chapters/{chapter_id}/items** - List pages of a chapter
- [ ] **GET /chapter-items/{id}** - Get page details
- [ ] **POST /chapters/{chapter_id}/items** - Add page to chapter
- [ ] **PUT /chapter-items/{id}** - Update page details
- [ ] **DELETE /chapter-items/{id}** - Delete page
- [ ] **POST /chapter-items/upload-image** - Upload chapter page image

## Recommendations
- [ ] 🟡 **GET /recommendations** - List recommendations
- [ ] 🟡 **GET /recommendations/{id}** - Get recommendation details
- [ ] 🟡 **POST /recommendations** - Create a recommendation
- [ ] 🟡 **PUT /recommendations/{id}** - Update recommendation
- [ ] 🟡 **DELETE /recommendations/{id}** - Delete recommendation

## Ads
- [ ] 🔵 **GET /ads** - List ads
- [ ] 🔵 **GET /ads/{id}** - Get ad details
- [ ] 🔵 **POST /ads** - Create an ad
- [ ] 🔵 **PUT /ads/{id}** - Update ad
- [ ] 🔵 **DELETE /ads/{id}** - Delete ad

---

### Team Members and Their Assigned Colors

indicate member's respective tasks 

| member   | color | 
|:---------|:----:| 
| Dang     | 🟢 | 
| NinhKieu | 🔴 | 
| Zachdemp | 🔵 | 
| Anhmien  | 🟡 |

