# Using Modules locally
To use modules locally create for each module a directory and run:

```zsh
go mod init <example/module>
```

To reference the local module in your main run following commands:
```zsh
go init <example/main_module>
go mod edit -replace <example/module>=../module
go mod tidy
```
It is assumed that the main module and the custom module are nabouring in a folder tree.