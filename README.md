# All-In-One Mac Update Script 🍎🖥️

[![Dependabot Updates](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/dependabot/dependabot-updates)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/andmpel/MacOS-All-In-One-Update-Script/master.svg)](https://results.pre-commit.ci/latest/github/andmpel/MacOS-All-In-One-Update-Script/master)
[![Test](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/test.yaml/badge.svg)](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/test.yaml)

> Inspired by the article
[Keeping MacOS Clean](https://waxzce.medium.com/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

This Zsh script simplifies the process of updating all your macOS software directly from the Terminal. While it covers many updates, you may want to install [`mas`](https://github.com/mas-cli/mas) to manage App Store applications.

## Getting Started

### Option 1: Install `update` as an Alias

This method ensures that the script always points to the latest available version. To install:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/install.sh | zsh
```

### Option 2: Install `update` as an Executable

This method installs the script so that the update command will point to the current version at the time of installation. To proceed, run:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/update-installer.sh | zsh
```

The script will automatically be placed as an executable file in `${HOME}/.local/bin`

#### **After installation, you can run the update script anytime by simply typing `update` in your terminal.**

## Supported Updates

This script currently updates the following:

- 🍺 **Homebrew** formulas and casks (`brew`)
- 🧑‍💻 **VSCode** extensions (`code`)
- 📦 **Node Package Manager** packages (`npm`)
- 💎 **RubyGems** (`gem`)
- 🧶 **Yarn** packages (`yarn`)
- 🚚 **Rust** packages (`cargo`)
- 🛍 **App Store** applications (`mas`)
- 🖥 **MacOS** system updates and patches (`softwareupdate`)

Feel free to contribute or customize the script to suit your needs! Happy updating! 🎉
