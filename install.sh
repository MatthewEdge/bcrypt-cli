#!/bin/sh
go build -o $(PWD)/build/bcrypt

grep -qxF '# bcrypt-cli' ~/.zshrc
if [ $? -eq 0 ]; then
    exit 0;
fi

echo "\n\n# bcrypt-cli" >> ~/.zshrc
echo "export PATH=\$PATH:$(PWD)/build" >> ~/.zshrc
source ~/.zshrc
    