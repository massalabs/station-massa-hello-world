# station-massa-hello-world

This repository show case how to create a plugin for Massa Station.

It is a simple plugin that only display "Hello world".

It contains a go package user by the other plugins to register them-self to Massa Station plugin manager: `pkg/plugin/register.go`
Here is how to use it in your plugin:

```golang
plugin.RegisterPlugin(listener, plugin.Info{
    Name: PluginName, Author: PluginAuthor,
    Description: PluginDescription, APISpec: "", Logo: logoFile,
})
```

These commands will help you build and manually install this plugin (for development purpose only):

```shell
go build -o hello-world station-massa-hello-world.go
mkdir -p ~/.config/massastation/plugins/hello-world
mv hello-world ~/.config/massastation/plugins/hello-world
```