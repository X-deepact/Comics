# be-admin



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin http://206.238.118.2:10180/comics/be-admin.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](http://206.238.118.2:10180/comics/be-admin/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Set auto-merge](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing (SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thanks to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README

Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.


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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 200 | OK |  |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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

| Code | Description |
| ---- | ----------- |
| 200 | OK |

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
| 200 | OK | [dto.UploadImageResponse](#dto.UploadImageResponse) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comic/comic/{id}

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
| 200 | OK | [dto.ComicResponse](#dto.ComicResponse) |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comic/comics

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

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.ComicResponse](#dto.ComicResponse) ] |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comic/create

#### POST
##### Summary:

Create a new comic

##### Description:

Create a new comic with the provided details

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| comic | body | Comic Request | Yes | [dto.ComicRequest](#dto.ComicRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ComicResponse](#dto.ComicResponse) |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comic/delete/{id}

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
| 200 | OK |  |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comic/update

#### PUT
##### Summary:

Update a comic

##### Description:

Update a comic with the provided details

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| comic | body | Comic Update Request | Yes | [dto.ComicUpdateRequest](#dto.ComicUpdateRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [dto.ComicResponse](#dto.ComicResponse) |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| BearerAuth | |

### /api/comic/upload-cover

#### POST
##### Summary:

Upload a comic cover

##### Description:

Upload a comic cover image

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cover | formData | Comic Cover Image | Yes | file |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Image URL | [dto.UploadImageResponse](#dto.UploadImageResponse) |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| page | query | Page number | No | integer |
| page_size | query | Page size | No | integer |
| name | query | Name | No | string |
| language | query | Language | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [ [dto.GenreResponse](#dto.GenreResponse) ] |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 200 | OK |  |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| page | query | Page number | No | integer |
| page_size | query | Page size | No | integer |
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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 200 | OK |  |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 200 | OK |  |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 200 | OK |  |
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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
| 400 | Invalid request | [dto.ErrorResponse](#dto.ErrorResponse) |
| 401 | User not authenticated | [dto.ErrorResponse](#dto.ErrorResponse) |
| 500 | Internal server error | [dto.ErrorResponse](#dto.ErrorResponse) |

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

#### dto.ComicRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| audience | string |  | No |
| code | string |  | No |
| cover | string |  | No |
| description | string |  | No |
| language | string |  | No |
| title | string |  | Yes |

#### dto.ComicResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| audience | string |  | No |
| code | string |  | No |
| cover | string |  | No |
| created_at | string |  | No |
| created_by | integer |  | No |
| created_by_name | string |  | No |
| description | string |  | No |
| id | integer |  | No |
| language | string |  | No |
| title | string |  | No |
| updated_at | string |  | No |
| updated_by | integer |  | No |
| updated_by_name | string |  | No |

#### dto.ComicUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | No |
| audience | string |  | No |
| code | string |  | No |
| cover | string |  | No |
| created_by | integer |  | No |
| description | string |  | No |
| id | integer |  | Yes |
| language | string |  | No |
| title | string |  | No |

#### dto.ErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

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

#### dto.TierModel

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | string |  | No |
| id | integer |  | No |

#### dto.UploadImageResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| image | string |  | No |

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