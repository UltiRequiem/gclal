# gclal

An utility to clone easily all your repositories from GitHub.
The process is done concurrently so it is quite fast.

## Usage

Basic usage:

```sh
gclal --username UltiRequiem
```

To use SSH:

```sh
gclal --username UltiRequiem --ssh true
```

If you have more than 100 Repositories, you will need an [API KEY](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token):

```
gclal --username UltiRequiem --ssh true --apiKey YOUR_API_KEY
```

## Install

```sh
go install github.com/UltiRequiem/gclal@latest
```

Or use a binary from [releases](https://github.com/UltiRequiem/gclal/releases/latest).

## FAQ

- Why `gclal`?

It was originally called `gclone_all`, but it seems too long to me so I took
some letters out of it.

## License

This project is licensed under the [MIT License](./license).
