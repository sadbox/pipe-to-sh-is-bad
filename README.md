pipe-to-sh-is-bad
=================

A proof of concept showing why piping to sh is a bad idea.

How to test
-----------
All you need to do is build the app. If it detects the curl user-agent it will serve up different content.

```
jmcguire@hyper-chicken:~$ wget -O - localhost:8080/install.sh 2>/dev/null 
Regular script for people to look at
jmcguire@hyper-chicken:~$ curl localhost:8080/install.sh
{INSERT BAD THINGS HERE}
```
