# Example Deepgram CLI Plugin Template

A template repo for creating a [Deepgram CLI](https://github.com/deepgram-devs/deepgram-cli) Plugin.

## Create a New Repo Using This Template (Well, Not Really)

Ideally, the easier way of creating a [Deepgram CLI](https://github.com/deepgram-devs/deepgram-cli) Plugin would be to use this repo template as a starting point. The reality is that this wwon't be the normal workflow for creating a template. The typical use case will probably be:

1. Some creates a "thing" OR a "project" already exists.
2. It would be nice for this "thing" to have a plugin.

Because of this, starting a project using this template probably won't be the way people get started.

Having said that, I think you could `git clone` this repo and the copy over the scaffolding that exists into your repo and open a PR to include it. Since this will probably be the path for inclusion, the documentation will focus on this method instead.

## Using the Example CLI

### Build the Plugin

When you create a [Deepgram CLI](https://github.com/deepgram-devs/deepgram-cli) plugin, you are effectively creating a Go plugin or in more basic terms, a shared library/object; hence, the `so` extension on the resulting file.

To read more about Go plugin, take a look at this simple [blog](https://dev.to/jacktt/plugin-in-golang-4m67).

To build your plugin:

```bash
cd ./pkg/plugin
go build -buildmode=plugin -ldflags="-s -w" -v
```

> **IMPORTANT:** You need to have build mechanisms (ie a Makefile for example) to orchestrate making the binaries for all your target platforms (macOS x86, macOS arm64, Linux amd64, etc, etc).

### Using the CLI Locally

Once you create the `.so`, you can also use this repo as example to build a standalone plugin without requiring the [Deepgram CLI](https://github.com/deepgram-devs/deepgram-cli).

To build a CLI for just your project, run:

```bash
cd ./cmd/plugin
cp ../../pkg/plugin/plugin.so ./
go run main.go
```

This will generate an executable that requires no external dependencies (ie everything is statically linked) and packaged with the `.so`, you have a standalone CLI for your project.

## Development and Contributing

Interested in contributing? We ❤️ pull requests!

To make sure our community is safe for all, be sure to review and agree to our
[Code of Conduct](./CODE_OF_CONDUCT.md). Then see the
[Contribution](./CONTRIBUTING.md) guidelines for more information.

## Getting Help

We love to hear from you so if you have questions, comments or find a bug in the
project, let us know! You can either:

- [Open an issue](https://github.com/deepgram-starters/example-starter-plugin/issues) on this repository
- [Join the Deepgram Github Discussions Community](https://github.com/orgs/deepgram/discussions)
- [Join the Deepgram Discord Community](https://discord.gg/xWRaCDBtW4)

## Further Reading

Check out the Developer Documentation at [https://developers.deepgram.com/](https://developers.deepgram.com/)
