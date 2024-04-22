# ooda (who-the) - AWS Profile Switcher in Go

```ascii
   ____  ____  ____  ___ 
  / __ \/ __ \/ __ \/   |        name:    ooda
 / / / / / / / / / / /| |        desc:    AWS Profile Switcher on the command-line
/ /_/ / /_/ / /_/ / ___ |        version: 0.1.0
\____/\____/_____/_/  |_|

```                         
ooda (who-the) is a fork of https://github.com/radiusmethod/awsd I created to personally test experiental features.
awsd by radiusmethod is a command-line utility that allows you to easily switch between AWS Profiles.

<img src="assets/demo.gif" width="500">

## Installation

Make sure you have Go installed. You can download it from [here](https://golang.org/dl/).

### Homebrew

```sh
brew tap radiusmethod/awsd
brew install awsd
```

### Makefile

```sh
make install
```

### To Finish Installation
Add the following to your bash profile or zshrc then open new terminal or source that file

```sh
alias awsd="source _awsd"
```

Ex. `echo 'alias awsd="source _awsd"' >> ~/.zshrc`

## Usage

### Switching AWS Profiles

It is possible to shortcut the menu selection by passing the profile name you want to switch to as an argument.

```bash
> awsd work
Profile work set.
```

To switch between different profiles files using the menu, use the following command:

```bash
awsd
```

This command will display a list of available profiles files in your `~/aws/config` file. Select the one you want to use.

## Persist Profile across new shells
To persist the set profile when you open new terminal windows, you can add the following to your bash profile or zshrc.

```bash
export AWS_PROFILE=$(cat ~/.awsd)
```

## Contributing

TBA: UNDER CONSTRUCTION

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


Inspired by & forked from [awsd](https://github.com/radiusmethod/awsd)
