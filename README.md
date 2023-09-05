# DCA IOT system

# Introduction

DCA IoT System is a rule-based device monitor and message processing system. This project is based on 

## Motivation

The system is designed for manage the IoT devices and devices' messages.

## Features

- Make decision by writing DCA rules.
- Allow to write multi-devices rules.
- High speed & easy to use.
- Allow for sharing device data among users.

## Outline
- [DCA IOT system](#dca-iot-system)
- [Introduction](#introduction)
  - [Motivation](#motivation)
  - [Features](#features)
  - [Outline](#outline)
- [Quick Start](#quick-start)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Basic Use](#basic-use)
- [Documentation](#documentation)
  - [Architecture](#architecture)
  - [Key Concepts](#key-concepts)
  - [DCA Rule](#dca-rule)
    - [Rule Syntax](#rule-syntax)
    - [Message Accepter](#message-accepter)
    - [Rule Matcher](#rule-matcher)
    - [Action Executor](#action-executor)
  - [WebSocket Connector](#websocket-connector)
- [License](#license)



# Quick Start

## Requirements

[MySQL service](#https://www.mysql.com) and [EMQX service](#https://www.emqx.io/) is required to running this system.

- Deploy MySQL service 
- Deploy EMQX service

Go version: 

Install [python](#https://www.python.org/) if you'd like to use the testing message sender script.

## Installation

1. Clone this project

```
git clone https://github.com/ChenKen9869/DCA-IOT-system.git
```

2. Build it locally

```shell
cd DCA-IOT-system
go build -o ./build/dca
```

3. Config the project at /config/application.yaml

```
vim ./config/application.yaml
```

**Deploy this project in server**

1. Use scp command to send the executable file to the server.

   ```
   scp ./build/dca another-server-directory
   ```

2. Use [go-swagger](#https://github.com/go-swagger/go-swagger) to generate the swagger api file. And also use scp command to send swagger files to the server.

   ```
   swag init
   scp -r ./docs another-server-directory
   ```

3. Use scp command to send the config file to the server

   ```
   scp ./config/application.yaml another-server-directory/config/
   ```

4. Use scp command to send the example device python file to the server for testing rules

   ```
   scp ./scripts/example_device.py another-server-directory/scripts/
   ```

5. Start the service

   ```
   ./dca
   ```

## Basic Use

After deployed and started service, test the system at swagger api webpage : <http://localhost:5930/swagger/index.html#/>

1. Create an account

2. Create a company

3. Create a biology

4. Create a Portable Device or a Fixed Device

5. Create a DCA rule

6. Start the DCA rule

7. Use websocket client to connect to the monitor center. Since the Authorism header is required while visit the websocket api. So, recommend for using [postman](#https://www.postman.com/) as the websocket client.

8. Use MQTT client to listen and get the message send by DCA runtime object. Recommend using [MQTTX desktop](#https://mqttx.app/) as MQTT client.

9. Run test python script to send data to the system. This script simulate sensor data.

   ```
   python ./script/example_device.py
   ```

   And send the following test message to the system.

   ```
   0000001, collar, temperature, 25.6
   ```

   "0000001" stands for device id; 

   "collar" stands for device type; 

   "temperature" stands for message attribution; 

   "25.6" stands for current data of attribution "temperature". 

   You can change the data in this message form. But when you do this, **change the device's information and rule description at the same time.**

10. When finish testing, remember to end the rule. 

    If you just close the service without end the rule, you should **change the rule's status in MySQL manually**. 

    ```sql
    UPDATE rules SET stat='Inactive' WHERE id=id_of_the_test_rule;
    ```
    
    If rule's status equals "Active" of "Scheduled", error will occur when you start or schedule the rule in the next time.

# Documentation

## Architecture

The system is made up by following architecture.



## Key Concepts

- user
- manager/visitor
- company
- biology
- device
- rule

## DCA Rule

### Rule Syntax

Datasource syntax: 

​	**Name{id, type, attribute}**
Condition syntax: 

​	**Expression of Name**
Action syntax: 

​	**ActionType : params; ActionType : params**



How to use in webpage?

test html

### Message Accepter

Write about how to create message accepter for developers.

1. inner accepter

2. outside accepter: update Datasource Management by http api or rpc call

### Rule Matcher

Write about how to create function rules for developers

create new matcher function

create the function condition type symbol

initial it in api/rule/init.go

### Action Executor

Write about how to create action executor for developers

create executor function and it's param channel

define the params list

initial it in api/rule/init.go

## WebSocket Connector

How it works?

connect

disconnect

use it inside the system

Or, use it outside by http api or rpc call

# License

