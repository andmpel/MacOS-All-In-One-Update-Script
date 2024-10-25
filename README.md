# All-In-One Mac Update script 🍎🖥️

> Inspired by the article
[Keeping macOS clean](https://waxzce.medium.com/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

This is a zsh Mac update script that updates all software I could find to be updated via Terminal on macOS.

Missing commands are not updated, but you might want
to install [`mas`](https://github.com/mas-cli/mas) to update applications from Appstore.

## Run

To Download & Execute, Run the following command in your terminal:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/update-all.sh | zsh
```

To execute just run:

```sh
zsh update-all.sh
```

To source and then use individual update-* functions first
comment out the command at the bottom of the file and run:

```sh
source ./update-all.sh
```

If you want to use this command often, run the below command
to download the install script:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/install.sh | zsh
```

Now you can call the script any time by running ```update``` 
in your zsh shell


## Updates

Currently including:

- 🍺 Homebrew formula's and casks (`brew`)
- 📚 Microsoft Office (`msupdate`)
- 🧑‍💻 VS Code Extensions (`code`)
- 📦 Node Package Manager (`npm`)
- 💎 RubyGems (`gem`)
- 🧶 Yarn (`yarn`)
- 🐍 Python3 (`pip`)
- 🚚 Rust Cargo Crates (`cargo`)
- 🔵 Applications in the Appstore (`mas`)
- 🖥 MacOS Operating System Updates/Patches (`softwareupdate`)

