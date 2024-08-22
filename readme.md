# CHAT APP NEXTJS - GOLANG

## STACKS
- nextjs
- shadcn
- golang echo
- tanstack query
- sqlite

## ROUTES FRONTEND
- /
- /login
- /register
- /logout
- /profile
- /profile/[id]
- /chat/[id]

## ROUTED BACKEND
- login [post]
- registered [post]
- logout [post]

- read message [get]
- send message [post]
- edit message [put]
- delete message [delete]

- edit profile [put]
- send image avatar [post]
- delete image avatar [delete]
- read profile [post]

- search user [get]

## SCHEMA DATABASE
- users
    - id #int id
    - username #string max 60 unique
    - password #string max 255
    - avatar #string max 255
    - createdAt #datetime
    - updatedAt #datetime

- messages
    - id #int id
    - send_id #int f users_id
    - receiven_id #int f users_id
    - message #string max 512
    - file_link #string max 255
    - createdAt #datetime
    - updatedAt #datetime

- refresh_tokens
    - id #int id
    - user_id #int f id
    - token #string max 255
    - createdAt #datetime
    - updatedAt #datetime

## reference
- [nextjs](https://nextjs.org/)
- [shadcn](https://ui.shadcn.com/docs/)
- [echo](https://echo.labstack.com/)
- [graphql](https://gqlgen.com/getting-started/)
- [lucide icons](https://lucide.dev/icons/)
- [golang](https://go.dev/)
- [air-verse](https://github.com/air-verse/air)
- [echo-sqlite-github](https://github.com/LeeGitaek/webgo)
- [echo-graphql-github](https://github.com/yuuu/gqlgen-echo-sample)
- [echo-websocket](https://echo.labstack.com/docs/cookbook/websocket)
- [tanstack-query](https://tanstack.com/query/latest/docs/framework/react/graphql)
- [gorm](https://gorm.io/docs/)