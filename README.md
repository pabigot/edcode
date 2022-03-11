# edcode

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Doc](https://pkg.go.dev/badge/github.com/pabigot/edcode.svg)](https://pkg.go.dev/github.com/pabigot/edcode)
[![Go Report Card](https://goreportcard.com/badge/github.com/pabigot/edcode)](https://goreportcard.com/report/github.com/pabigot/edcode)
[![Build Status](https://github.com/pabigot/edcode/actions/workflows/core.yml/badge.svg)](https://github.com/pabigot/edcode/actions/workflows/core.yml)
[![Coverage Status](https://coveralls.io/repos/github/pabigot/edcode/badge.svg)](https://coveralls.io/github/pabigot/edcode)

Package edcode provides wrapper types that implement
encoding.TextMarshaler and encoding.TextUnmarshaler for standard types
where Go1 doesn't do this.

Why?  Because I got tired of copying this into every application that
needs durations in its config files.
