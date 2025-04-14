ads


GET
/api/ads
Get advertisement list

Get list of advertisements with pagination

Parameters
Try it out
Name	Description
page *
integer
(query)
Page number

page
page_size *
integer
(query)
Page size

page_size
title
string
(query)
Filter by title

title
type
string
(query)
Filter by type (internal/external)

type
status
string
(query)
Filter by status (active/inactive)

status
sort_by
string
(query)
Sort by field (e.g. title, created_at, updated_at)

sort_by
sort
string
(query)
Sort order (ASC/DESC)

sort
Responses
Response content type

application/json
Code	Description
200	
Get advertisement list successfully

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active_from": "string",
      "active_to": "string",
      "comic_id": 0,
      "created_at": "string",
      "created_by_name": "string",
      "direct_url": "string",
      "id": 0,
      "image": "string",
      "status": "string",
      "title": "string",
      "type": "string",
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

authors


GET
/api/author
List GetAuthors

List all authors

Parameters
Try it out
Name	Description
Authorization *
string
(header)
Bearer authorization token

Authorization
name
string
(query)
Name

name
sort_by
string
(query)
Sort By

sort_by
sort
string
(query)
Sort

sort
page
integer
(query)
Page number

page
page_size
integer
(query)
Page size

page_size
Responses
Response content type

application/json
Code	Description
200	
List authors

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "biography": "string",
      "birth_day": "string",
      "created_at": "string",
      "created_by_name": "string",
      "id": 0,
      "name": "string",
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

chapter-items


GET
/api/chapter-items
List chapter items

Get list of chapter items with pagination

Parameters
Try it out
Name	Description
chapter_id *
integer
(query)
Chapter ID

chapter_id
page *
integer
(query)
Page number

page
page_size *
integer
(query)
Page size

page_size
sort_by
string
(query)
Sort by field (e.g. page, created_at, updated_at)

sort_by
sort
string
(query)
Sort order (ASC/DESC)

sort
Responses
Response content type

application/json
Code	Description
200	
List chapter items

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active": true,
      "active_from": "string",
      "chapter_id": 0,
      "created_at": "string",
      "created_by_name": "string",
      "id": 0,
      "image": "string",
      "page": 0,
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

chapters


GET
/api/chapters
Get chapters

Get chapters

Parameters
Try it out
Name	Description
comic_id
integer
(query)
Comic ID

comic_id
page
integer
(query)
Page number

page
page_size
integer
(query)
Page size

page_size
sort_by
string
(query)
Sort by

sort_by
sort
string
(query)
Sort

sort
Responses
Response content type

application/json
Code	Description
200	
List chapters

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active": true,
      "active_from": "string",
      "comic_id": 0,
      "cover": true,
      "created_at": "string",
      "created_by_name": "string",
      "id": 0,
      "name": "string",
      "number": 0,
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

comics


GET
/api/comics
List comics

List all comics

Parameters
Try it out
Name	Description
page
integer
(query)
Page number

page
page_size
integer
(query)
Page size

page_size
q
string
(query)
Search query

q
sort_by
string
(query)
Sort by

sort_by
sort
string
(query)
Sort order (asc, desc)

sort
active
boolean
(query)
Active


--
language
string
(query)
Language

language
audience
string
(query)
Audience

audience
author
integer
(query)
Author ID

author
genre
integer
(query)
Genre ID

genre
Responses
Response content type

application/json
Code	Description
200	
List comics

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active": true,
      "audience": "string",
      "authors": [
        {
          "biography": "string",
          "birth_day": "string",
          "created_at": "string",
          "created_by_name": "string",
          "id": 0,
          "name": "string",
          "updated_at": "string",
          "updated_by_name": "string"
        }
      ],
      "code": "string",
      "cover": "string",
      "created_at": "string",
      "created_by": 0,
      "created_by_user": {
        "active": true,
        "avatar": "string",
        "birthday": "string",
        "description": "string",
        "display_name": "string",
        "email": "string",
        "first_name": "string",
        "full_name": "string",
        "id": 0,
        "last_name": "string",
        "phone": "string",
        "role_name": "string",
        "tier_code": "string",
        "tier_id": 0,
        "username": "string"
      },
      "description": "string",
      "genres": [
        {
          "created_at": "string",
          "created_by_name": "string",
          "id": 0,
          "lang": "string",
          "name": "string",
          "position": 0,
          "updated_at": "string",
          "updated_by_name": "string"
        }
      ],
      "id": 0,
      "lang": "string",
      "name": "string",
      "original_language": "string",
      "status": "string",
      "updated_at": "string",
      "updated_by": 0,
      "updated_by_user": {
        "active": true,
        "avatar": "string",
        "birthday": "string",
        "description": "string",
        "display_name": "string",
        "email": "string",
        "first_name": "string",
        "full_name": "string",
        "id": 0,
        "last_name": "string",
        "phone": "string",
        "role_name": "string",
        "tier_code": "string",
        "tier_id": 0,
        "username": "string"
      }
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

dramas


GET
/api/dramas
Get list of dramas

Get list of dramas

Parameters
Try it out
Name	Description
page
integer
(query)
Page

page
page_size
integer
(query)
Page Size

page_size
sort_by
string
(query)
Sort by

sort_by
sort
string
(query)
Sort order

sort
Responses
Response content type

application/json
Code	Description
200	
List of dramas

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active": true,
      "created_at": "string",
      "created_by_name": "string",
      "genres": [
        {
          "id": 0,
          "name": "string",
          "position": 0
        }
      ],
      "id": 0,
      "release_date": "string",
      "thumb": "string",
      "translations": [
        {
          "description": "string",
          "language": "string",
          "title": "string"
        }
      ],
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success"
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

episodes


GET
/api/episodes
Get list of episodes

Get list of episodes

Parameters
Try it out
Name	Description
page *
integer
(query)
Page number

page
page_size *
integer
(query)
Page size (must be between 10 and 50)

page_size
sort_by
string
(query)
Sort by

sort_by
sort
string
(query)
Sort order

sort
Responses
Response content type

application/json
Code	Description
200	
List episodes

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active": true,
      "created_at": "string",
      "created_by_name": "string",
      "drama_id": 0,
      "id": 0,
      "number": 0,
      "subtitles": [
        {
          "language": "string",
          "url": "string"
        }
      ],
      "updated_at": "string",
      "updated_by_name": "string",
      "video": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
genres


GET
/api/genre
List genres

List all genre

Parameters
Try it out
Name	Description
page *
integer
(query)
Page number

page
page_size *
integer
(query)
Page size (must be between 10 and 50)

page_size
name
string
(query)
Name

name
language
string
(query)
Language

language
sort_by
string
(query)
Sort by

sort_by
sort
string
(query)
Sort order (asc, desc)

sort
Responses
Response content type

application/json
Code	Description
200	
List genres

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "created_at": "string",
      "created_by_name": "string",
      "id": 0,
      "lang": "string",
      "name": "string",
      "position": 0,
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
genres-drama


GET
/api/genres-drama
List genres

List all genres with pagination and filtering

Parameters
Try it out
Name	Description
page *
integer
(query)
Page number

page
page_size *
integer
(query)
Page size (must be between 10 and 50)

page_size
name
string
(query)
Filter by name

name
language
string
(query)
Filter by language

language
sort_by
string
(query)
Sort by field

sort_by
sort
string
(query)
Sort direction (asc/desc)

sort
Responses
Response content type

application/json
Code	Description
200	
List of genres

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "created_at": "string",
      "created_by": 0,
      "id": 0,
      "name": "string",
      "position": 0,
      "translations": [
        {
          "created_at": "string",
          "created_by": 0,
          "genre_id": 0,
          "id": 0,
          "language": "string",
          "translated_name": "string",
          "updated_at": "string",
          "updated_by": 0
        }
      ],
      "updated_at": "string",
      "updated_by": 0
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

recommends


GET
/api/recommend
List recommends

List all recommends

Parameters
Try it out
Name	Description
Authorization *
string
(header)
Bearer authorization token

Authorization
sort_by
string
(query)
Sort By

sort_by
sort
string
(query)
Sort

sort
page
integer
(query)
Page number

page
page_size
integer
(query)
Page size

page_size
title
string
(query)
Title

title
Responses
Response content type

application/json
Code	Description
200	
List recommends

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active_from": "string",
      "active_to": "string",
      "comics": [
        {
          "id": 0,
          "name": "string"
        }
      ],
      "cover": "string",
      "created_at": "string",
      "created_by_name": "string",
      "id": 0,
      "position": 0,
      "title": "string",
      "updated_at": "string",
      "updated_by_name": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}

users


GET
/api/user
List users

List all user

Parameters
Try it out
Name	Description
page *
integer
(query)
Page number

page
page_size *
integer
(query)
Page size (must be between 10 and 50)

page_size
username
string
(query)
Username

username
phone
string
(query)
Phone

phone
email
string
(query)
Email

email
name
string
(query)
Name

name
display_name
string
(query)
Display name

display_name
tier_id
integer
(query)
Tier ID

tier_id
sort_by
string
(query)
Sort by

sort_by
sort
string
(query)
Sort order (asc, desc)

sort
Responses
Response content type

application/json
Code	Description
200	
List users

Example Value
Model
{
  "code": "SUCCESS",
  "data": [
    {
      "active": true,
      "avatar": "string",
      "birthday": "string",
      "created_at": "string",
      "created_by_name": "string",
      "description": "string",
      "display_name": "string",
      "email": "string",
      "first_name": "string",
      "full_name": "string",
      "id": 0,
      "last_name": "string",
      "phone": "string",
      "role_name": "string",
      "tier_code": "string",
      "tier_id": 0,
      "updated_at": "string",
      "updated_by_name": "string",
      "username": "string"
    }
  ],
  "msg": "success",
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
400	
Invalid request

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}
500	
Internal server error

Example Value
Model
{
  "code": "ERROR",
  "msg": "string"
}


I have to make Dahsboard page with above APIs.
send me complete project
