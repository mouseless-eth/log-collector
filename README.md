# Log collector
A quick hack to quickly spin up a service that aggregates logs from multiple clients

## Running
Clients can send logs using
```bash
echo "Hello from client!" | nc <place-holder>.railway.app 9999
```
