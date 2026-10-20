# Running tower in a container

!!! warning "Community documentation"

    This page is not actively maintained by the tower authors and is
    written by community members. It is _not_ verified by tower developers.

    **It might be outdated and it might miss necessary steps**.

A container runtime such as [Docker](https://www.docker.com) or [Podman](https://podman.io) is required. The container
image can be found on [Docker Hub](https://github.com/issance/tower/releases) and [GitHub Container
Registry](https://github.com/juanfont/headscale/pkgs/container/headscale). The container image URLs are:

- [Docker Hub](https://github.com/issance/tower/releases): `docker.io/issance/tower:<VERSION>`
- [GitHub Container Registry](https://github.com/juanfont/headscale/pkgs/container/headscale):
  `ghcr.io/juanfont/tower:<VERSION>`

## Configure and run tower

1. Create a directory on the container host to store tower's [configuration](../../ref/configuration.md) and the SQLite database:

    ```shell
    mkdir -p ./tower/{config,lib}
    cd ./tower
    ```

1. Download the example configuration for your chosen version and save it as: `$(pwd)/config/config.yaml`. Adjust the
   configuration to suit your local environment. See [Configuration](../../ref/configuration.md) for details.

1. Start tower from within the previously created `./tower` directory:

    ```shell
    docker run \
      --name tower \
      --detach \
      --read-only \
      --tmpfs /tmp \
      --tmpfs /var/run/tower \
      --volume "$(pwd)/config:/etc/tower:ro" \
      --volume "$(pwd)/lib:/var/lib/tower" \
      --publish 127.0.0.1:8080:8080 \
      --publish 127.0.0.1:9090:9090 \
      --health-cmd "CMD tower health" \
      docker.io/issance/tower:<VERSION> \
      serve
    ```

    Note: use `0.0.0.0:8080:8080` instead of `127.0.0.1:8080:8080` if you want to expose the container externally.

    This command mounts the local directories inside the container, forwards port 8080 and 9090 out of the container so
    the tower instance becomes available and then detaches so tower runs in the background.

    A similar configuration for `docker-compose`:

    ```yaml title="docker-compose.yaml"
    services:
      tower:
        image: docker.io/issance/tower:<VERSION>
        restart: unless-stopped
        container_name: tower
        read_only: true
        tmpfs:
          - /var/run/tower
        ports:
          - "127.0.0.1:8080:8080"
          - "127.0.0.1:9090:9090"
        volumes:
          # Please set <HEADSCALE_PATH> to the absolute path
          # of the previously created tower directory.
          - <HEADSCALE_PATH>/config:/etc/tower:ro
          - <HEADSCALE_PATH>/lib:/var/lib/tower
        command: serve
        healthcheck:
            test: ["CMD", "tower", "health"]
    ```

1. Verify tower is running:

    Follow the container logs:

    ```shell
    docker logs --follow tower
    ```

    Verify running containers:

    ```shell
    docker ps
    ```

    Verify tower is available:

    ```shell
    curl http://127.0.0.1:8080/health
    ```

Continue on the [getting started page](../../usage/getting-started.md) to register your first machine.

## Debugging tower running in Docker

The Tower container image is based on a distroless image that does not contain a shell or any other debug tools. If you need to debug tower running in the Docker container, you can use the `-debug` variant, for example `docker.io/issance/tower:x.x.x-debug`.

### Running the debug Docker container

To run the debug Docker container, use the exact same commands as above, but replace `docker.io/issance/tower:x.x.x` with `docker.io/issance/tower:x.x.x-debug` (`x.x.x` is the version of tower). The two containers are compatible with each other, so you can alternate between them.

### Executing commands in the debug container

The default command in the debug container is to run `tower`, which is located at `/ko-app/tower` inside the container.

Additionally, the debug container includes a minimalist Busybox shell.

To launch a shell in the container, use:

```shell
docker run -it docker.io/issance/tower:x.x.x-debug sh
```

You can also execute commands directly, such as `ls /ko-app` in this example:

```shell
docker run docker.io/issance/tower:x.x.x-debug ls /ko-app
```

Using `docker exec -it` allows you to run commands in an existing container.
