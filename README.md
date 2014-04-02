README.md
========
simple web server that always returns the specified return code  

```
./fakeWeb -h  
Usage of ./fakeWeb:  
  -address="0.0.0.0:8082": address to listen on  
  -log=false: enable request logging to stdout  
  -status=200: status code to return  
```
```
Running it: 
nohup ./fakeWeb &                    # defaults - no logging, listen on 8082 and return 200s
nohup ./fakeWeb -log > access.log &  # log requests
nohup ./fakeWeb -status=500          # return 500s
```

