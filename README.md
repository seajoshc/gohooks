# gohooks

Work in progress...

Handle incoming webhooks with the Go stdlib.

Kinda focused on catching payloads from Atlassian Compass templates at the moment.

## Use

Clone and `go build`. Install [ngrok](https://ngrok.com) or similar.

`./gohooks`

`ngrok http 8888`

Use the URL ngrok gives you + `/compass` i.e. `https://1e99-69-420-11-214.ngrok-free.app/compass`
