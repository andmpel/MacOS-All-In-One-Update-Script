# All-In-One Mac Update Script 🍎🖥️

> Inspired by the article
[Keeping macOS clean](https://waxzce.medium.com/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

This Zsh script simplifies the process of updating all your macOS software directly from the Terminal. While it covers many updates, you may want to install [`mas`](https://github.com/mas-cli/mas) to manage App Store applications.

## Getting Started

### Install Script as Alias

To Download & Execute, Run the following command in your terminal:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/install.sh | zsh
```

**Now you can run the script anytime by simply typing `update` in your terminal.**

### Running the Script

To perform a full update, run:

```sh
zsh update-all.sh
```

If you want to use individual update functions, first comment out the last line of `update-all.sh`, then source the script:

```sh
source ./update-all.sh
```

### Setting Up for Frequent Use

For easy access, copy the script to a directory included in your `PATH`. Here’s how:

```sh
USER_SCRIPTS="${HOME}/.local/bin"  # Modify as needed
cp ./update-all.sh $USER_SCRIPTS/update-all
chmod +x $USER_SCRIPTS/update-all
```

Now you can run the script anytime by simply typing `update-all` in your terminal.

## Supported Updates

This script currently updates the following:

- 🍺 **Homebrew** formulas and casks (`brew`)
- 📚 **Microsoft Office** applications (`msupdate`)
- 🧑‍💻 **VSCode** extensions (`code`)
- 📦 **Node Package Manager** packages (`npm`)
- 💎 **RubyGems** (`gem`)
- 🧶 **Yarn** packages (`yarn`)
- 🐍 **Python3** packages (`pip`)
- 🧶 **Rust** packages (`cargo`)
- 🔵 **App Store** applications (`mas`)
- 🖥  **MacOS** system updates and patches (`softwareupdate`)

Feel free to contribute or customize the script to suit your needs! Happy updating! 🎉
