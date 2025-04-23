# An in-memory Key-Value database made with Go

An in-memory Key-Value database made with Go, for no particular reason other than to learn basic things about the
language (really I just thought it would be fun, and it actually was).

## Usage

```bash
go run main.go
```

## Commands

The commands below are also recognized even when they are not in uppercase.

- `HELP`: Display all commands
- `ALL`: Display all entries in the database
- `GET <key>`: Get the value of a key
- `CREATE <key> <value>`: Create a new entry in the database
- `DELETE <key>`: Delete an entry from the database
- `LEVEL`: Display the current level structure
- `DOWN`: Move down to the next level
- `APPLY`: Apply changes to the parent level
- `DISCARD`: Discard changes and move up to the parent level
- `EXIT`: Exit the program

## What are levels?

A level is the basic data type structure in this database. It essentially defines a map (which is actually where the
data is stored), a reference to the parent level, and a depth.

Each level can have its own changes made to the database, and they are only applied/discarded with the commands `APPLY`
and `DISCARD`. Both commands are very similar: they terminate the current level and move up to the parent level. The
difference is that `APPLY` applies the changes to the parent level, while `DISCARD` discards it.

You can always check where you are with the `LEVEL` command.
