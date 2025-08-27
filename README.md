# All-In-One Mac Update Script 🍎🖥️

[![Dependabot Updates](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/dependabot/dependabot-updates)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/andmpel/MacOS-All-In-One-Update-Script/master.svg)](https://results.pre-commit.ci/latest/github/andmpel/MacOS-All-In-One-Update-Script/master)
[![Test](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/test.yaml/badge.svg)](https://github.com/andmpel/MacOS-All-In-One-Update-Script/actions/workflows/test.yaml)

> Inspired by the article
[Keeping MacOS Clean](https://waxzce.medium.com/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

🔔 This script has recently been ported from Zsh/Bash to Go. If you prefer the old script [here are the instructions](https://github.com/andmpel/MacOS-All-In-One-Update-Script/blob/ReadMe-Revise/zsh-instructions.md)

This Go script simplifies the process of updating all your macOS software directly from the Terminal. While it covers many updates, you may want to install [`mas`](https://github.com/mas-cli/mas) to manage App Store applications.

## Getting Started

1. Head to releases and download `macup` or `macup_<cpu-arch>` for your mac
2. Run the `./macup` from terminal, then select your updates
3. To skip the selection menu, run `macup --yes` to use your previous selections

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
