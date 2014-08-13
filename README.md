# Occasions
Occasions is a program that handles special occasions dates. 
It collects the dates from a vcal files. 

# Install

Create paths
```
mkdir -p /etc/occasions/ics
```

Create a configuration file
```
cat<< EOF> /etc/occasions/occasions.conf
 # 
 # Configuration file for occasions
 # http://github.com/haukurk/occasions
 
 [general]
 
 # ics folder location 
 ics=/etc/occasions/ics
EOF

```

Install with go
``` 
go install github.com/haukurk/occasions
```

Then upload vCal files to /etc/occasions/ics.

I recommend to get your files at
https://www.mozilla.org/en-US/projects/calendar/holidays/

# CLI

To run occasions manually just run ```occasions```. 
It will then notify you if there is any occasions upcoming based on your vCal files.

# REST API

Not implemented yet.

# 
