#!/bin/bash

set -x

main() {
    install_aws_cli
    deploy_hello_zip
}

deploy_hello_zip() {
    GOOS=linux go build hello.go
    zip hello.zip ./hello

    echo $PWD
    aws lambda update-function-code --function-name AluminumFalcon --zip-file fileb://./hello.zip
}

install_aws_cli() {
    sudo apt-get update
    sudo apt-get -y install python-dev python-setuptools
    sudo easy_install --upgrade pip six
    sudo pip install urllib3==1.21.1 ;# to satisfy version dependency conflict
    sudo pip install awscli
}

[ "${BASH_SOURCE[0]}" == "$0" ] && main "$@"
