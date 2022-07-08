# demo-wal
This is a demo of wal performance test.


# Test Result
```
2022/07/08 18:49:02 [0]start
2022/07/08 18:49:02 [0]wal.1
2022/07/08 18:49:02 [0]wal.lastIndex: 200000
2022/07/08 18:49:02 [0]wal.write1: 11111 /s
2022/07/08 18:49:02 [0]wal.write2: 10526 /s
2022/07/08 18:49:02 [0]wal.read 12500 /s
2022/07/08 18:49:02 [0]wal.4
2022/07/08 18:49:02 [0]wal.5

```

# desc
The first write is about 60-100 per second, 
but the modified write can reach more than 10000
