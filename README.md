Nameifier - Because naming things is hard
=========================================

Nameifier is a tool for generating consistent, stable sets of names, using only
a seed string and a counter. It's nice for making lists of hostnames or other
data where a name is easier to remember than a number.

You simply hit the endpoint and you will get text blob returned that contains
the list of names. Samples follow

This is available, publicly-hosted at: https://nameifier.herokuapp.com

## URL Format

```
/nameifier/<seed string>/<count>/
```

## Generate a Hostname for an AWS Instance

http://nameifier.herokuapp.com/nameifier/i-12deadbeef345/1

## Generate 100 Hostnames for a Cluster

http://nameifier.herokuapp.com/nameifier/my-host-cluster/100

## Generate 1000 Names for Objects

http://nameifier.herokuapp.com/nameifier/random_objects/1000 
