#!/bin/sh
go build -o $(PWD)/build/bcrypt

echo "\n\n# bcrypt-cli" >> ~/.zshrc
echo "export PATH=\$PATH:$(PWD)/build" >> ~/.zshrc
source ~/.zshrc
    