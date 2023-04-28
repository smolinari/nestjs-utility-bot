# Utility Bot

[![build-and-test](https://github.com/nestjs-discord/utility-bot/actions/workflows/build-and-test.yaml/badge.svg)](https://github.com/nestjs-discord/utility-bot/actions/workflows/build-and-test.yaml)

A Discord bot designed to streamline the support process on the official NestJS Discord server.

As people who usually answer questions on [the official NestJS Discord server](https://discord.gg/nestjs), we've experienced that sometimes users ask questions that have already been answered many times before.

There are some common cases, like when they post a new support request, they don't provide a [minimal reproduction code](https://minimum-reproduction.wtf/), or sometimes they don't even share their code, and even if they do, they don't know how to put them in [code blocks](https://gist.github.com/matthewzring/9f7bbfd102003963f9be7dbcf7d40e51#code-blocks) properly.

So we devised the idea of having a [Discord bot](https://discord.com/developers/docs/intro#bots-and-apps) with predefined and well-written Markdown content as [slash commands](https://discord.com/developers/docs/interactions/application-commands) to reply to users instead of repeatedly writing and explaining.

## Todo

- [ ] NPM related (before interacting with the npm registry API)
  - [ ] Validate npm package names
  - [ ] Validate version numbers

- [ ] Features
  - [ ] Mark slash commands as `protected`
  - [ ] `npm > inspect` slash command https://registry.npmjs.org/@nestjs/core/latest

- [ ] Refactor
  - [ ] Wrap errors

## Configuration

```shell
cp .env.example .env
```

Three `DISCORD_APP_ID`, `DISCORD_BOT_TOKEN`, and `DISCORD_GUILD_ID` environment variables are required, and rest of
the configuration is located in `config.yml`.

Currently, the bot doesn't support hot-reloading. Instead, the application should restart to apply the changes.

## Build

To build this project, you must [install Golang](https://go.dev/doc/install) in your system
and execute the following command.

```shell
go build -trimpath -buildvcs=false -ldflags "-w" -o ./bin/utility-bot ./main.go
```

## Adding your bot to server

Please add the bot to the server using the invite link created by the `discord:invite` CLI command,
as the permissions within the link will likely change based on the features we may add.

## Docker usage

```shell
docker compose up -d

# Generates an invite link to add the bot to servers
docker compose exec utility-bot ./utility-bot discord:invite

# Graceful shutdown
# docker compose down

# Removing the registered slash commands (in case of inconsistency between the configuration and the production)
# docker compose exec utility-bot ./utility-bot discord:clean
```

## Notes

- Slash commands
  - Once slash commands are registered, Discord will sort them alphabetically, regardless of their initial order in `config.yml`.
  - Once slash commands are registered or removed, they get updated instantly for the end-user because this project uses [guild commands](https://discord.com/developers/docs/interactions/application-commands#registering-a-command) instead of global commands.
  - Discord has a global rate limit of [200 application command creations per day, per guild](https://discord.com/developers/docs/interactions/application-commands#registering-a-command).
  - Bot will automatically register non-registered slash commands on bootstrap.
  - Registered slash commands can be removed by `discord:clean` command.
- Markdown content
  - Content within the slash commands can have a maximum of 3500 characters.
  - The bot will cache Markdown content on memory to avoid spamming I/O.
- Moderators
  - They can be defined by their unique Discord ID in `config.yml`.
  - They bypass rate-limit policies.
  - They can execute `protected` commands in `config.yml`.

## Running tests

```shell
go test -v ./...
```

## Dependencies overview

- [DiscordGo](https://github.com/bwmarrin/discordgo) - Provides low level bindings to the Discord chat client API
- [Cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions
- [Viper](https://github.com/spf13/viper) - Complete configuration solution for Go applications
- [Validator](https://github.com/go-playground/validator) - Implements value validations for structs based on tags
- [Zerolog](https://github.com/rs/zerolog) - Zero allocation JSON logger
- [Go-humanize](https://github.com/dustin/go-humanize) - Formatters for units to human friendly sizes
- [Testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks