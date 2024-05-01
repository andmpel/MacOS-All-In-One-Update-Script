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
readonly RED='\033[31m'   # Red
readonly GREEN='\033[32m' # Green
readonly CLEAR='\033[0m'  # Clear color and formatting

update-brew() {
    if ! which brew &>/dev/null; then return; fi

    echo -e "${GREEN}Updating Brew Formula's${CLEAR}"
    brew update
    brew upgrade
    brew cleanup -s

    echo -e "\n${GREEN}Updating Brew Casks${CLEAR}"
    brew outdated --cask
    brew upgrade --cask
    brew cleanup -s

    echo -e "\n${GREEN}Brew Diagnostics${CLEAR}"
    brew doctor
    brew missing
}

update-gem() {
    if ! which gem &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Gems${CLEAR}"
    gem update --user-install
    gem cleanup --user-install
}

update-npm() {
    if ! which npm &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Npm Packages${CLEAR}"
    npm update -g
}

update-yarn() {
    if ! which yarn &>/dev/null; then return; fi

    echo -e "${GREEN}Updating Yarn Packages${CLEAR}"
    yarn upgrade --latest
}

update-pip2() {
    if ! which python2 &>/dev/null; then return; fi
    if ! which pip &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Python 2.x pips${CLEAR}"
    # python2 -c "import pkg_resources; from subprocess import call; packages = [dist.project_name for dist in pkg_resources.working_set]; call('pip install --upgrade ' + ' '.join(packages), shell=True)"
    pip list --outdated --format=columns | grep -v '^\-e' | cut -d = -f 1 | xargs -n1 pip install -U
}

update-pip3() {
    if ! which python3 &>/dev/null; then return; fi
    if ! which pip3 &>/dev/null; then return; fi

    echo -e "\n${GREEN}Updating Python 3.x pips${CLEAR}"
    # python3 -c "import pkg_resources; from subprocess import call; packages = [dist.project_name for dist in pkg_resources.working_set]; call('pip3 install --upgrade ' + ' '.join(packages), shell=True)"
    pip3 list --outdated --format=columns | grep -v '^\-e' | cut -d = -f 1 | xargs -n1 pip3 install -U
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

update-office() {
    local MS_OFFICE_UPDATE='/Library/Application\ Support/Microsoft/MAU2.0/Microsoft\ AutoUpdate.app/Contents/MacOS/msupdate'
    if [ ! -f $MS_OFFICE_UPDATE ]; then return; fi

    echo -e "\n${GREEN}Updating MS-Office${CLEAR}"
    $MS_OFFICE_UPDATE --install
}

update-all() {
    local PING_IP=8.8.8.8
    if ping -q -W 1 -c 1 $PING_IP &> /dev/null; then
        update-brew
        update-gem
        update-npm
        update-yarn
        update-pip2
        update-pip3
        update-app_store
        update-macos
        update-office # Enable only if MS-Office is installed in your system.
    else
        echo -e "${RED}Internet Disabled!!!${CLEAR}"
    fi
}

# COMMENT OUT IF SOURCING
update-all
