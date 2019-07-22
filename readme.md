### MySQL slave status

Cli util that just print to stdout short status of mysql slave instance
    
* Seconds behind master
* Status
* Error (if has)
    
### Installation

Run `go get github.com/bald2b/mysql-slave-status`

Minimal requirements:
* golang 1.11+

### Usage

* **-h** mysql **h**ost:port (default 127.0.0.1:3306)
* **-u** mysql **u**ser name (default root)
* **-p** mysql **p**assword, ask if empty
* **-r** **r**epeat every N seconds, no repeat if 0
* **-s** print only **s**econds behind master value

### Example

```
mysql-slave-status -h 172.17.0.1:3307 -u slave_user -p
```

Output:
```
Enter password: ******
Seconds behind master: 0
State: Slave has read all relay log; waiting for more updates
```

```
mysql-slave-status -h 172.17.0.1:3307 -u slave_user -p passw0rd -s 
```

Output:
```
341
```
