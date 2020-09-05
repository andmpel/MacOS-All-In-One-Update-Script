# All-In-One Mac Update script 🍎

> Inspired by the article
[Keeping macOS clean](https://medium.com/@waxzce/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

This is a zsh Mac update script that updates all software I could find to be updated via Terminal on macOS.

Missing commands are not updated, but you might want
to install [`mas`](https://github.com/mas-cli/mas) to update applications from Appstore.

## Run

To execute just run:

```sh
zsh update-all.sh
```

To source and then use individual update-* functions first
comment out the command at the bottom of the file and run:

```sh
source ./update-all.sh
```

If you want to use this command often copy it to directory that you
have in PATH variable (check with `echo $PATH`) like this:

```sh
USER_SCRIPTS="${HOME}/.local/bin"  # change this
cp ./update-all.sh $USER_SCRIPTS/update-all
chmod +x $USER_SCRIPTS/update-all
```

and now you can call the script any time :)


## Updates

Currently including:

- 🍺 Homebrew formula's and casks (`brew`)
- ⚛️ Atom (`apm`)
- 📦 Node Package Manager (`npm`)
- 💎 RubyGems (`gem`)
- 🧶 Yarn (`yarn`)
- 🐍 Python 2.7 and 3 (`pip`)
- 🔵 Applications in the Appstore (`mas`)
- 📚 Microsoft Office (`msupdate`)
- 🖥 MacOS Operating System Updates/Patches (`softwareupdate`)

