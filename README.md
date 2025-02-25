# udp-replicator
A simple UDP Replicator (UDP Proxy)

## Build
```shell
make build
```

## Test

### Test data

```bash
for i in {1..1000}; do $(echo "test-data$i" > /dev/udp/127.0.0.1/6343); done
```

