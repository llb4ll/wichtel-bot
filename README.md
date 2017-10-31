# wichtel-bot
A small wichtel bot to send out secret santa emails.

[![Build Status](https://travis-ci.org/llb4ll/wichtel-bot.svg?branch=master)](https://travis-ci.org/llb4ll/wichtel-bot)

## Features
- Add email addresses and the wichtel bot will randomly draw secret santas and send them emails containing the name and email of their wichtel
- You can dry-run and run the wichtel bot using Docker
- Configure smtp settings in a config file
- Provide secret santa email addresses using a config file
- Define exclusions for each secret santa (this person will not be in the draw). Perfect to prevent couples from gifting each other.
- Chirstmas style the email you sent out

## Technologies
- Golang: https://golang.org/
- Docker: https://www.docker.com/

## Build & Run
The easiest way to get started is to use Docker to build and run the project:

To build the project with docker:
`docker build -t "wichtel-bot" .`

To run the wichtel bot:
`docker run wichtel:latest`