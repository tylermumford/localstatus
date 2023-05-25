# Changelog

All notable changes to LocalStatus will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0).
The project strives to follow [Semantic Versioning](https://semver.org/#semantic-versioning-200).

## Unreleased

- (nothing yet)

## 0.2.2-beta - 2023-05-24

- Fixed #5 (tcp.open: program crashes without "label" parameter)

## 0.2.1-beta - 2023-05-24

- Fixed a memory leak
- Added CHANGELOG.md

## 0.2-beta - 2023-05-04

- Added --watch mode
- Added Homebrew tap & formula
- Added docs folder, generated from source comments
- Added red/green color in output text
- Added check: command (experimental)
- Added check: const
- Added check: git.branch
- Added check: npm.install

## 0.1-alpha - 2023-04-21

- First release
- Added TOML configuration loading from hardcoded path: `~/.config/localstatus.toml`
- Added check: file.exists
- Added check: env
- Added check: tcp.open
- Added check: http.ok
