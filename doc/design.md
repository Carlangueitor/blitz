## Library choosing

WebSockets: I know github.com/gorilla/websocket but it is now deprecated, next
logic option for me was golang.org/x/net/websocket but documentation recommends
a different package pkg.go.dev/nhooyr.io/websocket which turns out to have a
comparision.

I came after github.com/gobwas/ws and I liked the option of having a low level
API along with a higher level one.
