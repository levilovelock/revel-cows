# revel-cows module

A sample module for revel.

To add to your app simply run ```go get github.com/levilovelock/revel-cows``` then add the follow line to your ```conf/app.conf``` file:
```golang
module.cows = github.com/levilovelock/revel-cows
```

and this line to your ```conf/routes``` file:
```golang
module:cows
```

launch your app and find the magic of the cows at yourapp:port/moo ;)
