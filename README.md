# udp-replicator

A simple UDP Replicator (UDP Proxy)

## Build

```shell
make build
```

## Build deb-package

1. Clone the Repository:
    ```shell
    git clone https://github.com/exsver/udp-replicator.git
    ```
2. Install Dependencies:
    ```shell
    apt install make devscripts debhelper build-essential
    ```
3. Build the DEB Package:
    ```shell
    cd ./udp-replicator
    debuild -us -uc -b
    ```
4. Clean up after build:
    ```shell
    debuild -T clean
    ```

## Test

### Test data

```bash
for i in {1..1000}; do $(echo "test-data$i" > /dev/udp/127.0.0.1/6343); done
```

