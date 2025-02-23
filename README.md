# All-In-One Mac Update Script 🍎🖥️

> Inspired by the article
[Keeping MacOS Clean](https://waxzce.medium.com/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

This Zsh script simplifies the process of updating all your macOS software directly from the Terminal. While it covers many updates, you may want to install [`mas`](https://github.com/mas-cli/mas) to manage App Store applications.

## Getting Started

### Install Script as Alias for Repeat Use

To Download & Execute, Run the following command in your terminal:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/install.sh | zsh
```

### Manually Configuring Alias for Repeat Use

For easy access, copy the `update-all.sh` script to a directory included in your `PATH`. Here’s how:

```sh
USER_SCRIPTS="${HOME}/.local/bin"  # Modify as needed
cp ./update-all.sh $USER_SCRIPTS/update
chmod +x $USER_SCRIPTS/update
```

**Now you can run the script anytime by simply typing `update` in your terminal.**

## Supported Updates

This script currently updates the following:

- 🍺 **Homebrew** formulas and casks (`brew`)
- 🧑‍💻 **VSCode** extensions (`code`)
- 📦 **Node Package Manager** packages (`npm`)
- 💎 **RubyGems** (`gem`)
- 🧶 **Yarn** packages (`yarn`)
- 🚚 **Rust** packages (`cargo`)
- 🔵 **App Store** applications (`mas`)
- 🖥  **MacOS** system updates and patches (`softwareupdate`)

Feel free to contribute or customize the script to suit your needs! Happy updating! 🎉
