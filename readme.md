# mememe

**a blog app with official page host and variables render**

mememe is a blog app using db storage and static official page hosting and variables render.

## Install

### use released binary:
1. create a database named mememe.
2. run ***./mememe init_db*** to create data table.
3. run ***./mememe create_user username password*** to create your account.

### use source code build:
1. build binary
2. goto step: use released binary

## Variables
this app uses those variables stored in table mememe_setting to render the static pages.
- app_name: your app name, effect on title
- app_desc: your app description
- github: your github url at index
- twitter: your twitter url at index
- email: your email at footer
- phone: your phone at footer
- address: your address at footer
- admin: owner username
