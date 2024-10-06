# rater

![Test](https://github.com/thetechpanda/mutex/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thetechpanda/rater)](https://goreportcard.com/report/github.com/thetechpanda/rater)
[![Go Reference](https://pkg.go.dev/badge/github.com/thetechpanda/rater.svg)](https://pkg.go.dev/github.com/thetechpanda/rater)
[![Release](https://img.shields.io/github/release/thetechpanda/rater.svg?style=flat-square)](https://github.com/thetechpanda/rater/releases)
![Dependencies](https://img.shields.io/badge/Go_Dependencies-_None_-green.svg)

The `rater` package provides a way to monitor the rate of events or operations over a specified period.

The `Monitor` struct tracks the number of times its `Rate()` method is called and provides the current count through the `Value()` method.
At every interval the current count is sent through the `C` channel before being reset.

The caller can stop the monitor by cancelling the context passed to the `NewMonitor()` function.

## Examples

Check the test files for examples on how to use the `rater` package.

A simple command line example is available in the [`example`](./example/example.go) file.

## Contributing

Contributions are welcome and very much appreciated!

Feel free to open an issue or submit a pull request.

## License

The `rater` package is released under the MIT License. See the [LICENSE](LICENSE) file for details.
