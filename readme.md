#### Instruction
This is practice project for golang.
The repo is simple log collector, which the `log agent` receives the monitor task from etcd, then, watches the local log files, and send to kafka, in the same time, the  `log_transfer` consumers the message from the kafka and sends to elasticsearch, and finally you can view all log information on the kibana platform.

#### repo
- etcd
- kafka
- es
- kibana
