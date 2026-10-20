# Getting started

This page helps you get started with tower and provides a few usage examples for the tower command line tool
`tower`.

!!! note "Prerequisites"

    - Tower is installed and running as system service. Read the [setup section](../setup/requirements.md) for
      installation instructions.
    - The configuration file exists and is adjusted to suit your environment, see
      [Configuration](../ref/configuration.md) for details.
    - Tower is reachable from the Internet. Verify this by visiting the health endpoint:
      https://tower.example.com/health
    - The Tailscale client is installed, see [Client and operating system support](../about/clients.md) for more
      information.

## Getting help

The `tower` command line tool provides built-in help. To show available commands along with their arguments and
options, run:

=== "Native"

    ```shell
    # Show help
    tower help

    # Show help for a specific command
    tower <COMMAND> --help
    ```

=== "Container"

    ```shell
    # Show help
    docker exec -it tower \
      tower help

    # Show help for a specific command
    docker exec -it tower \
      tower <COMMAND> --help
    ```

!!! note "Manage tower from another local user"

    By default only the user `tower` or `root` will have the necessary permissions to access the unix socket
    (`/var/run/tower/tower.sock`) that is used to communicate with the service. In order to be able to
    communicate with the tower service you have to make sure the unix socket is accessible by the user that runs
    the commands. In general you can achieve this by any of the following methods:

    - using `sudo`
    - run the commands as user `tower`
    - add your user to the `tower` group

    To verify you can run the following command using your preferred method:

    ```shell
    tower users list
    ```

## Manage tower users

In tower, a node (also known as machine or device) is [typically assigned to a tower
user](../ref/registration.md#identity-model). Such a tower user[^1] may have many nodes assigned to them and can be
managed with the `tower users` command. Invoke the built-in help for more information: `tower users --help`.

### Create a tower user

=== "Native"

    ```shell
    tower users create <USER>
    ```

=== "Container"

    ```shell
    docker exec -it tower \
      tower users create <USER>
    ```

### List existing tower users

=== "Native"

    ```shell
    tower users list
    ```

=== "Container"

    ```shell
    docker exec -it tower \
      tower users list
    ```

## Register a node

One has to [register a node](../ref/registration.md) first to use tower as coordination server with Tailscale. The
following examples work for the Tailscale client on Linux/BSD operating systems. Alternatively, follow the instructions
to connect [Android](connect/android.md), [Apple](connect/apple.md) or [Windows](connect/windows.md) devices. Read
[registration methods](../ref/registration.md) for an overview of available registration methods.

### [Web authentication](../ref/registration.md#web-authentication)

On a client machine, run the `tailscale up` command and provide the FQDN of your tower instance as argument:

```shell
tailscale up --login-server <YOUR_HEADSCALE_URL>
```

Usually, a browser window with further instructions is opened. This page explains how to complete the registration on
your tower server and it also prints the Auth ID required to approve the node:

=== "Native"

    ```shell
    tower auth register --user <USER> --auth-id <AUTH_ID>
    ```

=== "Container"

    ```shell
    docker exec -it tower \
      tower auth register --user <USER> --auth-id <AUTH_ID>
    ```

### [Pre authenticated key](../ref/registration.md#pre-authenticated-key)

It is also possible to generate a preauthkey and register a node non-interactively. First, generate a preauthkey on the
tower instance. By default, the key is valid for one hour and can only be used once (see `tower preauthkeys --help` for other options):

=== "Native"

    ```shell
    tower preauthkeys create --user <USER>
    ```

=== "Container"

    ```shell
    docker exec -it tower \
      tower preauthkeys create --user <USER>
    ```

The command returns the preauthkey on success which is used to connect a node to the tower instance via the
`tailscale up` command:

```shell
tailscale up --login-server <YOUR_HEADSCALE_URL> --authkey <YOUR_AUTH_KEY>
```

[^1]: [Ensure that the Tower username does not end with `@`.](../ref/oidc.md#reference-a-user-in-the-policy)
