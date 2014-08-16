[![Build Status](https://travis-ci.org/haukurk/occasions.svg?branch=master)](https://travis-ci.org/haukurk/occasions)

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

 # port to listen to
 port=81
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

To run the REST interface you run the command ```occasions rest```.
You will then run a HTTP server listening on a port defined in your ```/etc/occasions/occasions.conf```.
The following resources are available:
* [GET] /api/occasions
  (Get all occasions in you vCal files)
* [GET] /api/occasions/upcoming
  (Get upcoming occasions from your vCal files)

# Tips!

I recommend you to add this as a greeting message to your shell. 
So that you will be notified when you login to your box.

For example if you are using Bash, you can simply add the following lines
to the bottom of your ```.bashrc```
```
# If not running interactively, don't do anything
if [[ $- != *i* ]]; then return; else occasions; fi     
```

Thus getting a greeting message like this:
```
Last login: Fri Aug 15 18:40:10 2014 from 69.69.69.69
[Occasions] Hi! Today is Fri, 08/15/14
[Occasions] No upcoming occasions found. Sorry :-(
[18:46:28] (root@orange:~)> 
```

