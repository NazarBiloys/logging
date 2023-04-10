# logging
Test ELK, GrayLog2 with different long_query_time for slow log query in mysql

### Small desc
I used `logstash` for export metrics to ES.
`Graylog` and `Kibana` got data from ES

### Result of benchmark with different long_query_time

Configured MySql with slow_query_log, set different global value -> `SET GLOBAL long_query_time = $value;`

For test run `siege` with 50 concurrency and 2000 requests per user in insert mode.
For test run `siege` with 5 concurrency and 5000 requests per user in read mode.

`  CREATE TABLE users (
id INT NOT NULL AUTO_INCREMENT,
firstname VARCHAR(50) NOT NULL,
lastname VARCHAR(50) NOT NULL,
email VARCHAR(255) NOT NULL,
password VARCHAR(255) NOT NULL,
date_of_birth DATE NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (id)
);
` Make a table of users.

#### GrayLog

* Insert to database


|         long_query_time:      | **0**  | **5** | **10** |
|:------------------------------:|:------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 100.00  | 100.00  |
|**Transaction rate:**, trans/sec |  666.71  |  765.76 |  751.60   |
|    **Response time**, secs    |  0.07  |  0.07   |  0.07   |
|  **Successful transactions**  | 100000   | 100000  | 100000  |
|    **Failed transactions**    |   0    |    0    |    0    |
| **Longest transaction**, secs |  0.88  |  0.38   |  0.46   |


* Read from database

|         long_query_time:      | **0**  | **5** | **10** |
|:------------------------------:|:------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 100.00  | 100.00  |
|**Transaction rate:**, trans/sec |  553.34  |  606.35 |  621.27   |
|    **Response time**, secs    |  0.01  |  0.01   |  0.01   |
|  **Successful transactions**  | 25000   | 25000  | 25000  |
|    **Failed transactions**    |   0    |    0    |    1    |
| **Longest transaction**, secs |  0.07  |  0.04   |  0.03   |

#### Kibana

* Insert to database

|         long_query_time:      | **0**  | **5** | **10** |
|:------------------------------:|:------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 100.00  | 100.00  |
|**Transaction rate:**, trans/sec |  666.13  |  765.76 |  755.06   |
|    **Response time**, secs    |  0.07  |  0.07   |  0.07   |
|  **Successful transactions**  | 100000   | 50000  | 50000  |
|    **Failed transactions**    |   0    |    0    |    0    |
| **Longest transaction**, secs |  0.89  |  0.43   |  0.41   |


* Read from database

|         long_query_time:      | **0**  | **5** | **10** |
|:------------------------------:|:------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 100.00  | 100.00  |
|**Transaction rate:**, trans/sec |  522.03  |  620.04 |  597.66   |
|    **Response time**, secs    |  0.01  |  0.01   |  0.01   |
|  **Successful transactions**  | 25000   | 25000  | 25000  |
|    **Failed transactions**    |   0    |    0    |    0    |
| **Longest transaction**, secs |  0.06  |  0.03   |  0.04   |
