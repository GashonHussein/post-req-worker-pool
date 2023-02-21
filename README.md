## post-req-worker-pool

Multi-threaded worker pool script that sends a series of HTTP POST requests to a given domain using TCP sockets. The program creates `N_WORKERS` worker goroutines that simultaneously open and close a single TCP connection, and each worker pulls requests from a shared pool of `N_REQUSTS` requests.

