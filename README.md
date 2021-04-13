# golembar
Golang waybar module for tracking the status and value of your golem miner.

<img src="https://github.com/bwoff11/golembar/blob/master/img/sample.png" align="center"
     alt="Size Limit logo by Anton Lovchikov" width="303" height="18">

## Installation

1. Download and run the latest version of the Golem provider software available [here](https://handbook.golem.network/provider-tutorials/provider-tutorial).
2. Download and install [Waybar](https://github.com/Alexays/Waybar).
3. Get an API key from coinmarketcap [here](https://coinmarketcap.com/api/).
4. Clone this repository, replace "YOUR_API_KEY" with your API key and build the software:```go build golembar.go```
5. Add the contents of [config_sample](https://github.com/bwoff11/golembar/blob/master/config_sample) to your waybar config.
6. Add the module as you would any other module to an orientation (e.g. "modules-right").
7. Reload your environment (Default: $mod+shift+R).
