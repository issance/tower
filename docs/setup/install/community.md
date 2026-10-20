# Community packages

Several Linux distributions and community members provide packages for tower. Those packages may be used instead of
the [official releases](official.md) provided by the tower maintainers. Such packages offer improved integration
for their targeted operating system and usually:

- setup a dedicated local user account to run tower
- provide a default configuration
- install tower as system service

!!! warning "Community packages might be outdated"

    The packages mentioned on this page might be outdated or unmaintained. Use the [official releases](official.md) to
    get the current stable version or to [test pre-releases](main.md).

    [![Packaging status](https://repology.org/badge/vertical-allrepos/tower.svg)](https://repology.org/project/tower/versions)

## Arch Linux

Arch Linux offers a package for tower, install via:

```shell
pacman -S tower
```

## Fedora, RHEL, CentOS

A third-party repository for various RPM based distributions is available at:
<https://copr.fedorainfracloud.org/coprs/jonathanspw/tower/>. The site provides detailed setup and installation
instructions.

## Nix, NixOS

A Nix package is available as: `tower`. See the [NixOS package site for installation
details](https://search.nixos.org/packages?show=tower).

## Gentoo

```shell
emerge --ask net-vpn/tower
```

Gentoo specific documentation is available [here](https://wiki.gentoo.org/wiki/User:Maffblaster/Drafts/Tower).

## OpenBSD

Tower is available in ports. The port installs tower as system service with `rc.d` and provides usage
instructions upon installation.

```shell
pkg_add tower
```
