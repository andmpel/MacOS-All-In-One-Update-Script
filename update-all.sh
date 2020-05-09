#!/bin/zsh
# To execute run:
#
#    zsh update-all.sh
# 
# To source and then use individual update-* functions
# first comment out the command at the bottom of the file
# and run:
# 
#    source ./update-all.sh
#
# If you want to use this command often copy it to directory
# that you have in PATH (check with `echo $PATH`) like this:
#
#     USER_SCRIPTS="${HOME}/.local/bin"  # change this
#     cp ./update-all.sh $USER_SCRIPTS/update-all
#     chmod +x $USER_SCRIPTS/update-all
#
#  and now you can call the script any time :)

# Text Color Variables
GREEN='\033[32m' # Green
CLEAR='\033[0m'  # Clear color and formatting

update-brew() {
    if ! which brew &>/dev/null; then return; fi

    echo -e "${GREEN}Updating Brew Formula's${CLEAR}"
    brew update
    brew upgrade
    brew cleanup -s

    echo -e "\n${GREEN}Updating Brew Casks${CLEAR}"
    brew cask outdated
    brew cask upgrade
    brew cleanup -s

    echo -e "\n${GREEN}Brew Diagnostics${CLEAR}"
    brew doctor
    brew missing
}

update-atom() {
    if ! which apm &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Atom${CLEAR}"
    apm upgrade -c false
}

update-npm() {
    if ! which npm &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating npm Packages${CLEAR}"
    npm update -g
}

update-gem() {
    if ! which gem &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating gems${CLEAR}"
    gem update --user-install
    gem cleanup --user-install
}

update-yarn() {
    if ! which yarn &>/dev/null; then return; fi

    echo -e "${GREEN}Updating Brew Formula's${CLEAR}"
    yarn upgrade --latest
}

update-pip2() {
    if ! which pip2 &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Python 2.7.X pips${CLEAR}"
    pip2 freeze - local | grep -v ‘^\-e’ | cut -d = -f 1 | xargs pip2 install -U
}

update-pip3() {
    if ! which pip3 &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Python 3.X pips${CLEAR}"
    pip3 freeze - local | grep -v ‘^\-e’ | cut -d = -f 1 | xargs pip3 install -U
}

update-app_store() {
    if ! which mas &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating App Store Applications${CLEAR}"
    mas outdated
    mas upgrade
}

update-macos() {
    echo -e "\n${GREEN}Updating Mac OS${CLEAR}"
    softwareupdate -i -a
}

update-all() {
    update-brew
    update-atom
    update-npm
    update-gem
    update-yarn
    update-pip2
    update-pip3
    update-app_store
    update-macos
}

# COMMENT OUT IF SOURCING
update-all

