# file-analizer
> This application is used to check the disk usage in bytes of all files in a given directory.

## Running / Getting started

```shell
go run cmd/app/main.go dir
```

Will run the application for the given directory

### Initial Configuration

This application can generate logs depending on the configuration settings, allowing
you to specify the level of logging used by the application. You can specify it by seeting 
a system variable to `LOG_LEVEL`. By default the log level is "info".


### Building

To generate an executable of the application, use the following command:

```shell
go CGO_ENABLED=0 build -o file-analizer cmd/app/main.go
```

This command will generate an executable of the application named `file-analizer`

## Features

This application will take a mount point as an argument and return a list of all
the files on the mountpoint and their disk usage in bytes json format.

Eg:

etdiskusage.py /tmp 
 
{ 
    "files": 
    [ 
      {"/tmp/foo", 1000}, 
      {"/tmp/bar", 1000000}, 
      {"/tmp/buzzz", 42}, 
    ], 
} 

What's all the bells and whistles this project can perform?
* What's the main functionality
* You can also do another thing
* If you get really randy, you can even do this

#### Argument 1
Type: `String`  
Default: `'default value'`

State what an argument does and how you can use it. If needed, you can provide
an example below.

Example:
```bash
awesome-project "Some other value"  # Prints "You're nailing this readme!"
```

## Contributing

If you'd like to contribute, please fork the repository and make changes as you'd like. Pull requests are warmly welcome.

## Licensing

"The code in this project is licensed under MIT license."
