# Admin Comics API Documentation
Swagger Admin Comics API Documentation.

## Version: 1.0


### Security
**BearerAuth**

|apiKey|*API Key*|
|---|---|
|In|header|
|Name|Authorization|

### /api/ads

#### GET
##### Summary:

Get advertisement list

##### Description:

Get list of advertisements with pagination

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page number | Yes | integer |
| page_size | query | Page size | Yes | integer |
| title | query | Filter by title | No | string |
| type | query | Filter by type (internal/external) | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.AdsResponse](#dto.AdsResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### POST
##### Summary:

Create a new advertisement

##### Description:

Create a new advertisement

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| ad | body | Advertisement Request | Yes | [dto.AdsCreateRequest](#dto.AdsCreateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.AdsResponse](#dto.AdsResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update advertisement

##### Description:

Update an existing advertisement

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| ad | body | Advertisement Update Request | Yes | [dto.AdsUpdateRequest](#dto.AdsUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.AdsResponse](#dto.AdsResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/ads/{id}

#### DELETE
##### Summary:

Delete advertisement

##### Description:

Delete an advertisement by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Advertisement ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Advertisement deleted successfully | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/author

#### POST
##### Summary:

Create author

##### Description:

Create a new author

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| author | body | Author Request | Yes | [dto.AuthorRequest](#dto.AuthorRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.AuthorResponse](#dto.AuthorResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/author/{id}

#### GET
##### Summary:

Get GetAuthorById

##### Description:

Get a author by GetAuthorById

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Author ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.AuthorResponse](#dto.AuthorResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update author by Id

##### Description:

Update an existing author

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Author ID | Yes | integer |
| author | body | Author Update Request | Yes | [dto.AuthorUpdateRequest](#dto.AuthorUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.AuthorResponse](#dto.AuthorResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/author/list

#### GET
##### Summary:

List GetAuthors

##### Description:

List all authors

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page number | No | integer |
| page_size | query | Page size | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.AuthorResponse](#dto.AuthorResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/chapter-items

#### GET
##### Summary:

List chapter items

##### Description:

Get list of chapter items with pagination

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| chapter_id | query | Chapter ID | Yes | integer |
| page | query | Page number | Yes | integer |
| page_size | query | Page size | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.ChapterItemResponse](#dto.ChapterItemResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### POST
##### Summary:

Create chapter item

##### Description:

Create a new chapter item (page)

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| item | body | Chapter Item Request | Yes | [dto.ChapterItemCreateRequest](#dto.ChapterItemCreateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ChapterItemResponse](#dto.ChapterItemResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update chapter item

##### Description:

Update an existing chapter item

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| item | body | Chapter Item Update Request | Yes | [dto.ChapterItemUpdateRequest](#dto.ChapterItemUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ChapterItemResponse](#dto.ChapterItemResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/chapter-items/{id}

#### DELETE
##### Summary:

Delete chapter item

##### Description:

Delete a chapter item by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Chapter Item ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Chapter item deleted successfully | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### GET
##### Summary:

Get chapter item

##### Description:

Get a chapter item by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Chapter Item ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ChapterItemResponse](#dto.ChapterItemResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/chapter-items/upload-image

#### POST
##### Summary:

Upload chapter image

##### Description:

Upload an image for a chapter item

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| file | formData | Image file | Yes | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/chapters

#### GET
##### Summary:

Get chapters

##### Description:

Get chapters

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| comic_id | query | Comic ID | No | integer |
| order | query | ASC or DESC | No | string |
| page | query | Page number | No | integer |
| page_size | query | Page size | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.ChapterResponse](#dto.ChapterResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### POST
##### Summary:

Create a new chapter

##### Description:

Create a new chapter

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| chapter | body | Chapter Request | Yes | [dto.ChapterRequest](#dto.ChapterRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ChapterResponse](#dto.ChapterResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update a chapter

##### Description:

Update a chapter

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| chapter | body | Chapter Update Request | Yes | [dto.ChapterUpdateRequest](#dto.ChapterUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ChapterResponse](#dto.ChapterResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/chapters/{id}

#### DELETE
##### Summary:

Delete a chapter

##### Description:

Delete a chapter by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Chapter ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Record deleted successfully | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### GET
##### Summary:

Get a chapter

##### Description:

Get a chapter

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Chapter ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ChapterResponse](#dto.ChapterResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comics

#### GET
##### Summary:

List comics

##### Description:

List all comics

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page number | No | integer |
| page_size | query | Page size | No | integer |
| q | query | Search query | No | string |
| sort_by | query | Sort by | No | string |
| sort | query | Sort order | No | string |
| active | query | Active | No | boolean |
| language | query | Language | No | string |
| audience | query | Audience | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | List of comics | [dto.ListResponse](#dto.ListResponse) & object |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### POST
##### Summary:

Create comic

##### Description:

Create a new comic

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | formData | Comic name | Yes | string |
| description | formData | Comic description | No | string |
| lang | formData | Language | No | string |
| audience | formData | Audience | No | string |
| genres | formData | Genre ID | No | [ integer ] |
| authors | formData | Author ID | No | [ integer ] |
| cover | formData | Comic cover image | No | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Comic created successfully | [dto.ComicResponse](#dto.ComicResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 401 | Unauthorized | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update comic

##### Description:

Update an existing comic

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | formData | Comic ID | Yes | integer |
| name | formData | Comic name | Yes | string |
| description | formData | Comic description | No | string |
| active | formData | Active | No | boolean |
| lang | formData | Language | No | string |
| audience | formData | Audience | No | string |
| genres | formData | Genre ID | No | [ integer ] |
| authors | formData | Author ID | No | [ integer ] |
| cover | formData | Comic cover image | No | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Comic updated successfully | [dto.ComicResponse](#dto.ComicResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 401 | Unauthorized | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comics/{id}

#### DELETE
##### Summary:

Delete a comic

##### Description:

Delete a comic by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Comic ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Comic successfully deleted | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### GET
##### Summary:

Get a comic

##### Description:

Get a comic by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Comic ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ComicReturn](#dto.ComicReturn) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comics/upload-cover

#### POST
##### Summary:

Upload a comic cover

##### Description:

Upload a comic cover image

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| file | formData | Comic cover image | Yes | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Comic cover uploaded successfully | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/general/tiers

#### GET
##### Summary:

Get tiers

##### Description:

Get tiers

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.TierModel](#dto.TierModel) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/genre

#### GET
##### Summary:

List genres

##### Description:

List all genre

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page number | Yes | integer |
| page_size | query | Page size (must be between 10 and 50) | Yes | integer |
| name | query | Name | No | string |
| language | query | Language | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.GenreResponse](#dto.GenreResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### POST
##### Summary:

Create a new genre

##### Description:

Create a new genre

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| genre | body | Genre Request | Yes | [dto.GenreCreateRequest](#dto.GenreCreateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.GenreResponse](#dto.GenreResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 401 | User not authenticated | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update a new genre

##### Description:

Update a new genre

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| genre | body | Genre Request | Yes | [dto.GenreUpdateRequest](#dto.GenreUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.GenreResponse](#dto.GenreResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/genre/{id}

#### DELETE
##### Summary:

Delete a genre

##### Description:

Delete a genre by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Genre ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Genre successfully deleted | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### GET
##### Summary:

Get a genre

##### Description:

Get a genre by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Genre ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.GenreResponse](#dto.GenreResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/recommend

#### POST
##### Summary:

Create recommend

##### Description:

Create a new recommend

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| recommend | body | Recommend Create Request | Yes | [dto.RecommendCreateRequest](#dto.RecommendCreateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.RecommendResponse](#dto.RecommendResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/recommend/{id}

#### GET
##### Summary:

Get recommend by Id

##### Description:

Get a recommend by Id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Recommend ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.RecommendResponse](#dto.RecommendResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update recommend by Id

##### Description:

Update an existing recommend

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Recommend ID | Yes | integer |
| recommend | body | Recommend Update Request | Yes | [dto.RecommendUpdateRequest](#dto.RecommendUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.RecommendResponse](#dto.RecommendResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/recommend/list

#### GET
##### Summary:

List recommends

##### Description:

List all recommends

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page number | No | integer |
| page_size | query | Page size | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.RecommendResponse](#dto.RecommendResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/user

#### GET
##### Summary:

List users

##### Description:

List all user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page number | Yes | integer |
| page_size | query | Page size (must be between 10 and 50) | Yes | integer |
| username | query | Username | No | string |
| phone | query | Phone | No | string |
| email | query | Email | No | string |
| name | query | Name | No | string |
| display_name | query | Display name | No | string |
| tier_id | query | Tier ID | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.UserResponse](#dto.UserResponse) ] |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### POST
##### Summary:

Create a new user

##### Description:

Create a new user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| username | formData | Username (Alphanumeric) | Yes | string |
| phone | formData | Phone Number | No | string |
| email | formData | Email Address | No | string |
| birthday | formData | Birthday (YYYY-MM-DD) | No | string |
| password | formData | Password (Min: 7 chars) | Yes | string |
| first_name | formData | First Name | No | string |
| last_name | formData | Last Name | No | string |
| display_name | formData | Display Name | Yes | string |
| description | formData | Description | No | string |
| avatar | formData | Avatar File | No | file |
| tier_id | formData | Tier ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.UserDetailDto](#dto.UserDetailDto) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 401 | User not authenticated | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update a new user

##### Description:

Update a new user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | formData | ID | Yes | integer |
| username | formData | Username (Alphanumeric) | Yes | string |
| phone | formData | Phone Number | No | string |
| email | formData | Email Address | No | string |
| birthday | formData | Birthday (YYYY-MM-DD) | No | string |
| first_name | formData | First Name | No | string |
| last_name | formData | Last Name | No | string |
| display_name | formData | Display Name | Yes | string |
| description | formData | Description | No | string |
| avatar | formData | Avatar File | No | file |
| tier_id | formData | Tier ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.UserDetailDto](#dto.UserDetailDto) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 401 | User not authenticated | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/user/{id}

#### DELETE
##### Summary:

Delete a user

##### Description:

Delete a user by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | User successfully deleted | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### GET
##### Summary:

Get a user

##### Description:

Get a user by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.UserDetailDto](#dto.UserDetailDto) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/user/{id}/active

#### PUT
##### Summary:

Activate/Deactivate a user

##### Description:

Activate/Deactivate a user by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | User successfully activated/deactivated | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/user/change-password

#### PUT
##### Summary:

Change password

##### Description:

Change password

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| ChangePassword | body | Change Password Request | Yes | [dto.UserChangePasswordRequest](#dto.UserChangePasswordRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Password changed successfully | [dto.ResponseMessage](#dto.ResponseMessage) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/user/login

#### POST
##### Summary:

Login

##### Description:

Authenticates the user and returns a JWT token

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| login | body | User Login Data | Yes | [dto.LoginRequest](#dto.LoginRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Login successful | [dto.LoginResponse](#dto.LoginResponse) |
| 400 | Invalid request |  |

### /api/user/profile

#### GET
##### Summary:

Get profile

##### Description:

Get profile

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.UserDetailDto](#dto.UserDetailDto) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

#### PUT
##### Summary:

Update profile

##### Description:

Update profile

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| username | formData | Username (Alphanumeric) | Yes | string |
| phone | formData | Phone Number | No | string |
| email | formData | Email Address | No | string |
| birthday | formData | Birthday (YYYY-MM-DD) | No | string |
| first_name | formData | First Name | No | string |
| last_name | formData | Last Name | No | string |
| display_name | formData | Display Name | Yes | string |
| description | formData | Description | No | string |
| avatar | formData | Avatar File | No | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.UserResponse](#dto.UserResponse) |
| 400 | Invalid request | [dto.ResponseMessage](#dto.ResponseMessage) |
| 401 | User not authenticated | [dto.ResponseMessage](#dto.ResponseMessage) |
| 500 | Internal server error | [dto.ResponseMessage](#dto.ResponseMessage) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### Models


#### dto.AdsCreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active_from | string |  | No |
| active_to | string |  | No |
| comic_id | integer |  | No |
| direct_url | string |  | No |
| image | string |  | No |
| title | string |  | Yes |
| type | string |  | Yes |

#### dto.AdsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active_from | string |  | No |
| active_to | string |  | No |
| comic_id | integer |  | No |
| created_at | string |  | No |
| created_by_name | string |  | No |
| direct_url | string |  | No |
| id | integer |  | No |
| image | string |  | No |
| title | string |  | No |
| type | string |  | No |
| updated_at | string |  | No |
| updated_by_name | string |  | No |

#### dto.AdsUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active_from | string |  | No |
| active_to | string |  | No |
| comic_id | integer |  | No |
| direct_url | string |  | No |
| id | integer |  | Yes |
| image | string |  | No |
| title | string |  | No |
| type | string |  | No |

#### dto.AuthorRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| biography | string |  | No |
| birth_day | string |  | No |
| name | string |  | Yes |

#### dto.AuthorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| biography | string |  | No |
| birth_day | string |  | No |
| created_at | string |  | No |
| created_by | integer |  | No |
| id | integer |  | No |
| name | string |  | No |
| updated_at | string |  | No |
| updated_by | integer |  | No |

#### dto.AuthorUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| biography | string |  | No |
| birth_day | string |  | No |
| name | string |  | No |

#### dto.ChapterItemCreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| active_from | string |  | No |
| chapter_id | integer |  | Yes |
| image | string |  | Yes |
| page | integer |  | Yes |

#### dto.ChapterItemResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| active_from | string |  | No |
| chapter_id | integer |  | No |
| created_at | string |  | No |
| created_by_name | string |  | No |
| id | integer |  | No |
| image | string |  | No |
| page | integer |  | No |
| updated_at | string |  | No |
| updated_by_name | string |  | No |

#### dto.ChapterItemUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| active_from | string |  | No |
| id | integer |  | Yes |
| image | string |  | No |
| page | integer |  | No |

#### dto.ChapterRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| comic_id | integer |  | Yes |
| cover | boolean |  | No |
| name | string |  | No |
| number | integer |  | No |

#### dto.ChapterResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| comic_id | integer |  | No |
| cover | boolean |  | No |
| created_at | string |  | No |
| created_by | string |  | No |
| created_by_user | string |  | No |
| id | integer |  | No |
| name | string |  | No |
| number | integer |  | No |
| updated_at | string |  | No |
| updated_by | string |  | No |
| updated_by_user | string |  | No |

#### dto.ChapterUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| comic_id | integer |  | Yes |
| cover | boolean |  | No |
| id | integer |  | Yes |
| name | string |  | No |
| number | integer |  | No |

#### dto.ComicResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| audience | string |  | No |
| code | string |  | No |
| cover | string |  | No |
| created_at | string |  | No |
| created_by | integer |  | No |
| description | string |  | No |
| id | integer |  | No |
| lang | string |  | No |
| name | string |  | No |
| updated_at | string |  | No |
| updated_by | integer |  | No |

#### dto.ComicReturn

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| audience | string |  | No |
| code | string |  | No |
| cover | string |  | No |
| created_at | string |  | No |
| created_by | integer |  | No |
| created_by_user | [dto.UserDetailDto](#dto.UserDetailDto) |  | No |
| description | string |  | No |
| id | integer |  | No |
| lang | string |  | No |
| name | string |  | No |
| updated_at | string |  | No |
| updated_by | integer |  | No |
| updated_by_user | [dto.UserDetailDto](#dto.UserDetailDto) |  | No |

#### dto.GenreCreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| language | string |  | Yes |
| name | string |  | Yes |
| position | integer |  | Yes |

#### dto.GenreResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| created_at | string |  | No |
| created_by_name | string |  | No |
| id | integer |  | No |
| language | string |  | No |
| name | string |  | No |
| position | integer |  | No |
| updated_at | string |  | No |
| updated_by_name | string |  | No |

#### dto.GenreUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | Yes |
| language | string |  | Yes |
| name | string |  | Yes |
| position | integer |  | Yes |

#### dto.ListResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data |  |  | No |
| pagination | [dto.Pagination](#dto.Pagination) |  | No |

#### dto.LoginRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| password | string |  | Yes |
| username | string |  | Yes |

#### dto.LoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| access_token | string |  | No |
| access_token_expires_at | string |  | No |
| user | [dto.UserDetailDto](#dto.UserDetailDto) |  | No |

#### dto.Pagination

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| page | integer |  | No |
| page_size | integer |  | No |
| total | integer |  | No |

#### dto.RecommendCreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| activeFrom | integer |  | No |
| activeTo | integer |  | No |
| cover | string |  | No |
| id | integer |  | No |
| position | integer |  | No |
| title | string |  | No |

#### dto.RecommendResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| activeFrom | string |  | No |
| activeTo | string |  | No |
| cover | string |  | No |
| createdAt | string |  | No |
| createdBy | integer |  | No |
| id | integer |  | No |
| position | integer |  | No |
| title | string |  | No |
| updatedAt | string |  | No |
| updatedBy | integer |  | No |

#### dto.RecommendUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| activeFrom | string |  | No |
| activeTo | string |  | No |
| cover | string |  | No |
| position | integer |  | No |
| title | string |  | No |

#### dto.ResponseMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | string |  | No |
| message | string |  | No |
| status | string |  | No |

#### dto.TierModel

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | string |  | No |
| id | integer |  | No |

#### dto.UserChangePasswordRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| current_password | string |  | Yes |
| new_password | string |  | Yes |

#### dto.UserDetailDto

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| avatar | string |  | No |
| birthday | string |  | No |
| description | string |  | No |
| display_name | string |  | No |
| email | string |  | No |
| first_name | string |  | No |
| full_name | string |  | No |
| id | integer |  | No |
| last_name | string |  | No |
| phone | string |  | No |
| role_name | string |  | No |
| tier_code | string |  | No |
| tier_id | integer |  | No |
| username | string |  | No |

#### dto.UserResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| avatar | string |  | No |
| birthday | string |  | No |
| created_at | string |  | No |
| created_by_name | string |  | No |
| description | string |  | No |
| display_name | string |  | No |
| email | string |  | No |
| first_name | string |  | No |
| full_name | string |  | No |
| id | integer |  | No |
| last_name | string |  | No |
| phone | string |  | No |
| role_name | string |  | No |
| tier_code | string |  | No |
| tier_id | integer |  | No |
| updated_at | string |  | No |
| updated_by_name | string |  | No |
| username | string |  | No |