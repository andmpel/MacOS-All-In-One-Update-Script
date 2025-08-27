# Instructions for Zsh/Bash script

### Install Script as Alias for Repeat Use

To Download & Execute, Run the following command in your terminal:

```sh
curl -fsSL https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/install.sh | zsh
```

### Manually Downloading and Running Script

For easy access, save the `update-all.sh` script to your Mac user's home folder, make it executable, and then run it.

```sh
USER_SCRIPTS="${HOME}/"
curl -fsSLo "$USER_SCRIPTS/update" https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/update-all.sh
chmod +x "$USER_SCRIPTS/update"
```

**Now you can run the script anytime by simply typing `./update` from your home directory in your terminal.**
