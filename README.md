## Development

We provide make targest for ergonomic development:

- `watch-n-serve`: Will watch for all go source code files and start a development server, watching for changes and restarting server.
- `watch-n-test`: Same as target bellow but running tests instead of start the server
- `serve`: run development server
- `test`: run tests

# API

Service has a message based JSON API, the message structure is the following:
```json
{
    "op": "<operation>",
    "topic": "<topic>",
    "payload": "<payload>"
}
```

`op` is the only required field, `topic` and `payload` will be required by `op` basis.

## `op`
Operations are the commands you can send or recieve from server. Valid Operations are:

- `subscriber-id`: Sent by server at client connect to get assigned subscriber ID, `payload` will contain the subscriber id.
- `subscribe-topic`: Subscribe to topic, `topic` is required.
- `broadcast`: Broadcast message, `topic` required and `payload` must contain the message.
- `broadcasted`: Broadcasted message in `topic` (require client to be subscribed to `topic`.
- `unsubscribe-topic`: Unsuscribe from `topic`, requires `topic`.
