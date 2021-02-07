# Go Muzei 

![giphy](https://media.giphy.com/media/MPcr38SwPlwcYL2e4s/giphy.gif)

[Muzei](https://github.com/romannurik/muzei) is an Android application that will periodically refreshes the users background to a famous artwork. The goal of this project is to replicate this for terminals and desktop backgrounds. 

:fire: This project has only been tested to work on MacOS. :fire:

## Installing

```sh
go get -u github.com/pokom/go-wallpaper
go-wallpaper 
```

## Running

To see a list of commands, run:

```sh
go-wallpaper --help
```

### Muzei


```sh
go-wallpaper desktop --provider muzei
```

### Reddit

Set the desktop to a random image from the top 10 of `reddit.com/r/earthporn`

```sh
go-wallpaper desktop --provider reddit
```

Set the desktop to random image from a different subreddit:

:warning: Not gauranteed to work. Need to filter out posts that are not a single image :warning:

```sh
go-wallpaper desktop --provider reddit --subreddit cityporn
```

## Building Locally

```sh
go build
./go-wallpaper
```

## Feature Roadmap

- [ ] iterm background 
    - [x] cli: `go-wallpaper iterm`
    - [ ] on launch of iterm
    - [ ] periodically(~24 hours default?)
- [ ] MacOS Desktop Wallpaper
    - [x] cli: `go-wallpaper wallpaper`
    - [ ] on login
    - [ ] periodically
- [ ] Windows Desktop Wallpaper
    - [ ] cli
    - [ ] on login
    - [ ] periodically
- [x] display info for current background image
- [ ] open link to info page in system browser
- [ ] multiple providers
    - [x] reddit
    - [ ] local images
    - [ ] bucket storage?
