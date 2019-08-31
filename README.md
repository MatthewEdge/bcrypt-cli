# bcrypt-cli

Small utility to enable bcrypt from the CLI

## Usage

`bcrypt -help` to see flags

`echo -n 'STR' | bcrypt` to create a bcrypt-hashed string

Make sure to use single quotes `'` and the `-n` flag for echo - otherwise the newline will be hashed as well
