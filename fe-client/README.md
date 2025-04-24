# fe-client

## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin http://206.238.118.2:10180/comics/fe-client.git
git branch -M main
git push -uf origin main
```

# comic-frontend

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

## Git rule

- Branch out from `develop` branch
- Do not push into `develop` branch directly, make a merge request
- Update your local `develop` branch and do an interactive rebase before pushing your feature and making a merge request.
- `(Optional)` Using _rebase_ instead of _merge_. [How to do it](https://www.atlassian.com/git/tutorials/merging-vs-rebasing)
- Naming branch:
  - `feature/your-name_your-branch-feature`
  - `bugfix/your-name_your-branch-fix`
  - example: `feature/bebong_add-i18n-config`

## Folder structure

- locales: This folder contains all files necessary to handle multi-language support.

- models: all client models and interfaces

- components: components for part of page

- views: component for each page (view)

- stores: states management using pinia

- router: handle routing

- layouts: component for common layout (guest, user,...)

- assets: static resources folder (styles, icons, images,...)
