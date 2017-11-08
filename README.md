![Klingon Flag](https://github.com/fegoulart/fegoulartAssets/blob/master/klingonFlag.jpg)

## Fernando Spock Project

This is a command line application to translate Star Trek character names from **English** to **Klingon**  and find out its species.

## Architecture decisions

This application is built in GoLang and I explain why

* I love to learn new tricks (it's my first Go experience)
* Go is from Google and Google rocks
* Go is so Hipster
* Go is perfect for command line apps
* Go is fun 

## Business rules assumptions

1. English characters not considered in Klingon alphabet (both lower and uppercase)

    * c
    * f
    * g
    * k
    * x
    * z
    
2. Namesakes
    * First occurrence will be considered (eg. **Kellin**)

## Build and Execution

`go build fernandoSpock`

`./fernandoSpock <characterName>`

Example

`./fernandoSpock Uhura`<br>
`0xF8E5 0xF8D6 0xF8E1 0xF8D0`<br>
`Human`

## Tests

I used **testify** library to help me with automated tests

If you haven't installed it yet, just run the following command

`go get github.com/stretchr/testify/assert`

Test cases are in fernandoSpock_test.go file.
Just run the code below to run it

`go test -v`

## License

Feel free to use it whenever you want. Pull requests are more than welcome.<br>

![Spock greeting](https://github.com/fegoulart/fegoulartAssets/blob/master/spock.jpg)
<br>
Bedankt