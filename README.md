# Fun Banking

The official Fun Banking rewrite with Go, HTMX, and SQLite. This inspiration came from needed
server-side rendering to boost performance and SEO, as well as needing to keep our
cost down so that we can be in business; thus the "boring" architecture.

# Getting Started

1. Create your `.env`

```.env
TEMPLATES_PATH=templates
DATABASE_URL=fun_banking.db
```

# Local Environment

We recommend that you use Air. It's a dev server the will listen to your files.

```shell
air init
```

Make sure that you update this in your `.air.toml` file

```diff
+ cmd = "go build -o ./tmp/main cmd/fun-banking/main.go"
```

Now, you can just run the following command and your server will automatically launch in watch mode.

```shell
air
```
