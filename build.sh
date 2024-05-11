#!/bin/bash

RUN_NAME="tticket"
mkdir -p /etc/${RUN_NAME}
sudo cp tticket.toml /etc/${RUN_NAME}/tticket.toml.tpl
go build -o /etc/${RUN_NAME}/${RUN_NAME}