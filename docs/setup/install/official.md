# Official releases

Official releases for tower are available as binaries for various platforms and DEB packages for Debian and Ubuntu.
Both are available on the [GitHub releases page](https://github.com/juanfont/headscale/releases).

## Using packages for Debian/Ubuntu (recommended)

It is recommended to use our DEB packages to install tower on a Debian based system as those packages configure a
local user to run tower, provide a default configuration and ship with a systemd service file. Supported
distributions are Ubuntu 22.04 or newer, Debian 12 or newer.

1. Download the [latest tower package](https://github.com/juanfont/headscale/releases/latest) for your platform (`.deb` for Ubuntu and Debian).

    ```shell
    HEADSCALE_VERSION="" # See above URL for latest version, e.g. "X.Y.Z" (NOTE: do not add the "v" prefix!)
    HEADSCALE_ARCH="" # Your system architecture, e.g. "amd64"
    wget --output-document=tower.deb \
     "https://github.com/juanfont/headscale/releases/download/v${HEADSCALE_VERSION}/headscale_${HEADSCALE_VERSION}_linux_${HEADSCALE_ARCH}.deb"
    ```

1. Install tower:

    ```shell
    sudo apt install ./tower.deb
    ```

1. [Configure tower by editing the configuration file](../../ref/configuration.md). An up-to date example
   configuration file is also available in `/usr/share/doc/tower/examples/config-example.yaml`:

    ```shell
    sudo nano /etc/tower/config.yaml
    ```

1. Restart tower to pick up configuration changes:

    ```shell
    sudo systemctl restart tower
    ```

1. Verify that tower is running as intended:

    ```shell
    sudo systemctl status tower
    ```

Continue on the [getting started page](../../usage/getting-started.md) to register your first machine.

## Using standalone binaries (advanced)

!!! warning "Advanced"

    This installation method is considered advanced as one needs to take care of the local user and the systemd
    service themselves. If possible, use the [DEB packages](#using-packages-for-debianubuntu-recommended) or a
    [community package](community.md) instead.

This section describes the installation of tower according to the [Requirements and
assumptions](../requirements.md#assumptions). Tower is run by a dedicated local user and the service itself is
managed by systemd.

1. Download the latest [`tower` binary from GitHub's release page](https://github.com/juanfont/headscale/releases):

    ```shell
    sudo wget --output-document=/usr/bin/tower \
    https://github.com/juanfont/headscale/releases/download/v<TOWER VERSION>/headscale_<TOWER VERSION>_linux_<ARCH>
    ```

1. Make `tower` executable:

    ```shell
    sudo chmod +x /usr/bin/tower
    ```

1. Add a dedicated local user to run tower:

    ```shell
    sudo useradd \
     --create-home \
     --home-dir /var/lib/tower/ \
     --system \
     --user-group \
     --shell /usr/sbin/nologin \
     tower
    ```

1. Download the example configuration for your chosen version and save it as: `/etc/tower/config.yaml`. Adjust the
   configuration to suit your local environment. See [Configuration](../../ref/configuration.md) for details.

    ```shell
    sudo mkdir -p /etc/tower
    sudo nano /etc/tower/config.yaml
    ```

1. Copy [tower's systemd service file](https://github.com/juanfont/headscale/blob/main/packaging/systemd/headscale.service)
   to `/etc/systemd/system/tower.service` and adjust it to suit your local setup. The following parameters likely need
   to be modified: `ExecStart`, `WorkingDirectory`, `ReadWritePaths`.

1. In `/etc/tower/config.yaml`, override the default `tower` unix socket with a path that is writable by the
   `tower` user or group:

    ```yaml title="config.yaml"
    unix_socket: /var/run/tower/tower.sock
    ```

1. Reload systemd to load the new configuration file:

    ```shell
    systemctl daemon-reload
    ```

1. Enable and start the new tower service:

    ```shell
    systemctl enable --now tower
    ```

1. Verify that tower is running as intended:

    ```shell
    systemctl status tower
    ```

Continue on the [getting started page](../../usage/getting-started.md) to register your first machine.
