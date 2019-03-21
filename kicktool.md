## To Deploy:
      1. in client directory gopherjs build
      2. ls
      4. gopherjs build -m
      5.  mv *.js ../static
      6.   build file by referencing the root directory
             * run the app from the root of the project  


# Kick Tool
Kickstart Go web server upon modification of source file

*including subfiles


1. install kick

   ```go get -u github.com/isomorphicgo/kick```
   
2. add ENV variable 

vi ~/.bash_profile

```
export KICKSTART_APP_ROOT=${GOPATH}/src/github.com/JayneJacobs/FullStackWebDev/kickstart/
```

kick --appPath=$KICKSTART_APP_ROOT --gopherjsAppPath=$KICKSTART_APP_ROOT/client --mainSourceFile=gopherface.go
