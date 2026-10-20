# Tower

The private network of issance. Tower is the control server that lets your phone reach the computers where
your agents run, from anywhere, without opening ports: every device gets a private address and talks to the
others directly, end to end encrypted, with Tower's own relay when a direct path is not possible.

Devices connect with the standard Tailscale apps (iPhone, Android, Mac, Windows, Linux) pointed at your Tower.
The issance installer does that for you.

```bash
tower serve -c /etc/tower/config.yaml        # run the server
tower users create acme                      # a customer
tower preauthkeys create --user 1 --reusable # the key their installer uses
```

Tower is a rebranded distribution of a BSD-3-Clause project; see NOTICE-ISSANCE.md and LICENSE.
