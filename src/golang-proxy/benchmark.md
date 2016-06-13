
# Benchmark

With Reverse Proxy (Hello-World + time.Nanos):
```
ab -n 100 -c 100 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        20 bytes

Concurrency Level:      100
Time taken for tests:   0.026 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      13700 bytes
HTML transferred:       2000 bytes
Requests per second:    3871.32 [#/sec] (mean)
Time per request:       25.831 [ms] (mean)
Time per request:       0.258 [ms] (mean, across all concurrent requests)
Transfer rate:          517.94 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    3   0.9      3       4
Processing:     3   15   3.7     14      21
Waiting:        3   14   3.8     14      21
Total:          7   17   3.5     17      24

Percentage of the requests served within a certain time (ms)
  50%     17
  66%     19
  75%     20
  80%     21
  90%     22
  95%     23
  98%     24
  99%     24
 100%     24 (longest request)
 ```

 No Proxy (Hello-World + time.Nanos):

 ```
 ab -n 100 -c 100 http://localhost:8080/
 This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
 Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
 Licensed to The Apache Software Foundation, http://www.apache.org/

 Benchmarking localhost (be patient).....done


 Server Software:
 Server Hostname:        localhost
 Server Port:            8080

 Document Path:          /
 Document Length:        19 bytes

 Concurrency Level:      100
 Time taken for tests:   0.008 seconds
 Complete requests:      100
 Failed requests:        0
 Total transferred:      13600 bytes
 HTML transferred:       1900 bytes
 Requests per second:    12858.43 [#/sec] (mean)
 Time per request:       7.777 [ms] (mean)
 Time per request:       0.078 [ms] (mean, across all concurrent requests)
 Transfer rate:          1707.76 [Kbytes/sec] received

 Connection Times (ms)
               min  mean[+/-sd] median   max
 Connect:        2    3   0.4      3       4
 Processing:     1    2   0.5      2       3
 Waiting:        1    2   0.3      2       2
 Total:          3    5   0.8      5       6

 Percentage of the requests served within a certain time (ms)
   50%      5
   66%      5
   75%      5
   80%      6
   90%      6
   95%      6
   98%      6
   99%      6
  100%      6 (longest request)
  ```