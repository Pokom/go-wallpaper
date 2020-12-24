# Go Muzei 

![hippo](https://media.giphy.com/media/MPcr38SwPlwcYL2e4s/giphy.gif)

[Muzei](https://github.com/romannurik/muzei) is an Android application that will periodically refreshes the users background to a famous artwork. The goal of this project is to replicate this for terminals and desktop backgrounds. 

## Installing

```sh
go get -u github.com/pokom/go-muzei
go-muzei 
```

## Building Locally

```sh
go build
./go-muzei
```

## Feature Roadmap

- [ ] iterm background 
    - [x] cli: `go-muzei iterm`
    - [ ] on launch of iterm
    - [ ] periodically(~24 hours default?)
- [ ] MacOS Desktop Wallpaper
    - [x] cli: `go-muzei wallpaper`
    - [ ] on login
    - [ ] periodically
- [ ] Windows Desktop Wallpaper
    - [ ] cli
    - [ ] on login
    - [ ] periodically
- [x] display info for current background image
- [ ] open link to info page in system browser
- [ ] multiple providers
    - [ ] reddit
    - [ ] local images
    - [ ] bucket storage?
