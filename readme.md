#### Instruction
This is a practice project for golang.  
The repo is a simple log collector, which the `log agent` receives the monitor task from etcd, then, watches the local log files, and send log lines to kafka, in the same time, the  `log_transfer` consumers the message from the kafka and sends to elasticsearch, and finally you can view all log information on the kibana platform.

#### What you can learn about?
- etcd
- kafka
- es
- kibana
- docker 
