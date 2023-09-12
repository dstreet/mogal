# Mogal

The Movie Gallery

## Getting Started

This project uses [Task](https://taskfile.dev/) as its task runner. Follow the [installation instructions](https://taskfile.dev/installation/) for your system to install.

**Run the initial project setup**:
```sh
task setup
```

**Add .env file**:
```sh
cp .env.example .env
```

## Running

To the run the application you must have Docker and Docker Compose installed.

```sh
task up
```

If this is the first run of the application, you will need to apply the database schema:

```sh
task db.migrate.up
```

Once the application is up and running, you can view it at: http://localhost:3000

## Development

### Starting the Database

```sh
task db
```

### Starting the API

```sh
task api.dev
```

### Starting the UI

```sh
task ui.dev
```

## Tests

```sh
task test
```