# Log collector
A quick hack that spins up a helper service to aggregate logs from multiple clients

## Running
Clients can send logs using
```bash
echo "Hello from client!" | nc <place-holder>.railway.app 9999
```
