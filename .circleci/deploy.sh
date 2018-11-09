#!/bin/bash

set -x

main() {
    set_up_aws_credentials
    install_aws_cli
    deploy_main_zip
}

set_up_aws_credentials() {
  mkdir "${HOME}/.aws"
  printf "[default]\naws_access_key_id = ${AWS_ACCESS_KEY_ID}\naws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}\nregion = us-east-2" > "${HOME}/.aws/config"
}

install_aws_cli() {
    sudo apt-get update
    sudo apt-get -y install python-dev python-setuptools
    sudo easy_install --upgrade pip six
    sudo pip install urllib3==1.21.1 ;# to satisfy version dependency conflict
    sudo pip install awscli
}

deploy_main_zip() {
    mkdir bin
    cd bin
    env GOOS=linux GOARCH=amd64 go build -o ./main github.com/cbusby/Starish-Wars-Backend/cmd/main
    zip ./main.zip ./main

    aws lambda update-function-code --function-name swb-lambda --zip-file fileb://./main.zip
}

[ "${BASH_SOURCE[0]}" == "$0" ] && main "$@"
