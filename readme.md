# mememe

**a blog app with official page host and variables render**

mememe is a blog app using db storage and static official page hosting and variables render.

## Install

### use released binary:
1. put your ssl .crt&.key files at the same directory of mememe
2. create a database named mememe.
3. run ***./mememe init_db*** to create data table.
4. run ***./mememe create_user username password*** to create your account.

### use source code build:
1. build binary
2. fllow steps ***use released binary***

## Features
1. http redirect
2. ip blocker(incorrect password input for three times)
3. markdown edit & preview(blog article)

## Render Variables
this app uses those variables stored in table mememe_setting to render the static pages.
- app_name: your app name, effect on title
- app_desc: your app description
- github: your github url at index page header
- twitter: your twitter url at index page header
- email: your email at page footer
- phone: your phone at page footer
- address: your address at page footer
- admin: owner username

## Dependences
1. [kross-hugo](https://github.com/themefisher/kross-hugo)
2. [gorm](https://github.com/jinzhu/gorm)
3. [pongo2](https://github.com/flosch/pongo2)
4. [echo](https://github.com/labstack/echo)
5. [jwt-go](github.com/dgrijalva/jwt-go)
6. [fsnotify](github.com/fsnotify/fsnotify)
7. [pandoc](https://github.com/jgm/pandoc)