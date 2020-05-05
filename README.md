# All-In-One Update script

> Inspired by the article
[Keeping macOS clean](https://medium.com/@waxzce/keeping-macos-clean-this-is-my-osx-brew-update-cli-command-6c8f12dc1731).

This is a bash MacOS Update script that updates all software I could find to be updated via Terminal on MacOS.

If e.g. `brew` command is missing, than it is not updated.
You might want to install [`mas`](https://github.com/mas-cli/mas) to update applications from Appstore, though.


## Run

To execute run:

   zsh update-all.sh

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

- 🍺 Homebrew formula's and casks
- 📦 Atom
- 📦 npm
- 📦 gem
- 📦 yarn
- 📦 Python 2.7.X and 3.X pip
- 🍏 Applications in the Appstore
- 🍎 MacOS Operating System Updates/Patches.

