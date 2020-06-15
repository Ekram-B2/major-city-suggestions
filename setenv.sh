#!/bin/bash

export DEPLOYMENT_TYPE=1
export DEVELOPMENT_CONFIG_PATH=config/files/config.development.json
export PRODUCTION_CONFIG_PATH=config/files/config.production.json

export DEVELOPMENT_SERVICE_PATH="http://127.0.0.1:8082"
export PRODUCTION_SERVICE_PATH="https://shielded-refuge-18793.herokuapp.com"
export PORT=8080

export CONFIG_OPERATION_TYPE="default"

