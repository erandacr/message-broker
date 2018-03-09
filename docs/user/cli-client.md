# Broker Admin Command Line Interface (CLI)

Broker Command Line Interface is used to perform all kinds of administrative operations in the Message Broker. 

Following are are the list of possible operations supported by the Broker CLI 

1. Initialize the Broker connection settings
2. Exchanges related administrative operations
    i. Exchange creation
    ii. Exchange information retrieval
    iii. Exchange deletion
3. Queues related administrative operations    
    i. Queue creation
    ii. Queue information retrieval
    iii. Queue deletion
4. Bindings related administrative operations
    i. Binding creation
    ii. Binding information retrieval
5. Consumer related administrative operations
    i. Consumer information retrieval


##Specification

###General command format

`./broker-admin.sh [action] [resource-type]? [resource-name]? [flag]*`

###CLI actions

init
create
list
delete
 
###CLI resource types

exchange
queue
binding
consumer