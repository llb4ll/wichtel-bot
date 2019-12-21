# wichtel-bot
A small wichtel bot to send out secret santa emails.

[![Build Status](https://travis-ci.org/llb4ll/wichtel-bot.svg?branch=master)](https://travis-ci.org/llb4ll/wichtel-bot)

## Features
- Add email addresses and the wichtel bot will randomly draw secret santas and send them emails containing the name and email of their wichtel.
- You can dry-run and run the wichtel bot using Docker if you don't change the `settings.json`.
- Configure smtp settings in a config file `settings.json`.
- Provide secret santa email addresses using a config file `wichtel.json`.
- Define exclusions for each secret santa (this person will not be in the draw). Perfect to prevent couples from gifting each other.

## Technologies
- Golang: https://golang.org/
- Docker: https://www.docker.com/

## Build & Run
You can either run the go program directly or use docker.

### Using Docker
The easiest way to get started is to use Docker to build and run the project:

0. Build the project with docker: `docker build -t "wichtel-bot" .`

1. Copy the folder `config`. Let's name the copied folder `my-config`.
 
2. Adapt `my-config/settings.json` to add your smtp server settings.

2. Add participating wichtel/santas to `my-config/wichtel.json`.

4. To run the wichtel bot:
	`docker run -v <ABSOLUTE-PATH-TO-MY-CONFIG>:/go/src/wichtel-bot/config wichtel-bot:latest`
	
### Using Go
Run the wichtelbot with `go run src/wichtelbot.go` - In this case you will need to edit the config files directly.

## Algorithm
The first implementation will be a greedy shuffle that is performed for each secret santa. 
In case a secret santa is left without a matching wichtel the draw is repeated.
The maximum amount of tries is set to 100.
Of course this implementation has the disadvantage of not beeing deterministic.
An improvement idea would be to use an algorithm finding a perfect matching in a bipartite graph.

Links:
- https://en.wikipedia.org/wiki/Hall%27s_marriage_theorem
