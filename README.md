# bcrypt-cli

Small utility to enable bcrypt from the CLI

If using zsh the `install.sh` script will build the binary and install an entry to your PATH

## Usage

`bcrypt -help` to see flags

`echo -n 'STR' | bcrypt` to create a bcrypt-hashed string

Make sure to use single quotes `'` and the `-n` flag for echo - otherwise the newline will be hashed as well
