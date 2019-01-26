# Pushover

## Usage

Ensure that the following environment variables are set:
* `PUSHOVER_TOKEN`
* `PUSHOVER_USER`

These can be found in the Pushover settings page.

From there, call the executable with all applicable flags (`--message` is required)

```
pushover --title "File Transfers" --message "Your files have finished transferring" --url "https://www.dropbox.com"
```

## Building

Ensure that Docker is installed on your computer and run `./build.sh` to build Linux and MacOS binaries.
