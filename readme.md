# gclal

An utility to clone easily all your repositories from GitHub. The process is
done concurrently so it is quite fast.

## Usage

Basic usage:

```sh
gclal --username UltiRequiem
```

To use SSH:

```sh
gclal --username UltiRequiem --ssh true
```

If you have more than 100 Repositories, you will need an
[API KEY](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token):

```
gclal --username UltiRequiem --ssh true --apiKey YOUR_API_KEY
```

## Install

```sh
go install github.com/UltiRequiem/gclal@latest
```

Or use a binary from
[releases](https://github.com/UltiRequiem/gclal/releases/latest).

## FAQ

- Why `gclal`?

It was originally called `gclone_all`, but it seems too long to me so I took
some letters out of it.

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://twitter.com/UltiRequiem) 😆

Consider [supporting me on Patreon](https://patreon.com/UltiRequiem) if you like
my work 🙏

Don't forget to start the repo ⭐

## Authors

[Eliaz Bobadilla](https://ultirequiem.com) - Creator and Maintainer 💪

See also the full list of
[contributors](https://github.com/UltiRequiem/gclal/contributors) who
participated in this project ✨

## Licence

Licensed under the MIT License 📄
