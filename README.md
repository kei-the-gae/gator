# Gator

This is a [guided project] (https://www.boot.dev/courses/build-blog-aggregator-golang) I completed as a part of the Boot.dev back-end developer curriculum. This is an RSS feed aggregator that allows users to subscribe to various RSS feeds, view the latest articles from those feeds, and manage their subscriptions. The project focuses on practicing SQL queries and database design.

[![golang](https://badgen.net/badge/go/1.24.1/cyan?icon=https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_LightBlue.svg)](https://go.dev/)

## Features

- Follow and unfollow RSS feeds
- Fetch and parse RSS posts
- View the latest articles from followed feeds
- Store user subscriptions
- Multi-user support

## Usage

Go version 1.24.1 or higher and PostgreSQL 17 is required to run this project.

Run the following command to install the application:

```bash
go install github.com/kei-the-gae/gator@latest
```

Run the following command to create the config file:

```bash
cat > ~/.gatorconfig.json << EOF
{"db_url":"postgres://$(whoami):@localhost:5432/gator?sslmode=disable"}
EOF
```

Run the following command to create the PostgreSQL database:

```bash
createdb gator
```

Now you can run the application from the command line.
`gator register <name>` registers a new user

`gator login <name>` logs in an existing user

`gator feeds` lists all the RSS feeds available in the database

`gator addfeed <name> <feed_url>` adds a new feed to the database and allows the logged-in user to follow it

`gator follow <feed_url>` allows the logged-in user to follow an RSS feed

`gator unfollow <feed_url>` allows the logged-in user to unfollow an RSS feed

`gator following` lists all the RSS feeds that the logged-in user is currently following

`gator browse <limit>` lists the latest articles from all the RSS feeds that the logged-in user is following. You can specify an optional limit to the number of articles to display (default is 2).

`gator agg <time_between_reqs>` fetches the latest articles from all the RSS feeds in the database and aggregates them into a single list. This command will parse the feeds and store the latest articles in the database. The argument is a time duration string that specifies the time to wait between requests to fetch the feeds. This is to avoid overwhelming the RSS feed servers. For example, `gator agg 5s` will wait 5 seconds between each request to fetch the feeds.
