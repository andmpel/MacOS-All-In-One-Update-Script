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
    echo -e "${GREEN}Updating Brew Formula's${CLEAR}"

    if ! command -v brew &>/dev/null; then
        echo -e "${RED}Brew is not installed.${CLEAR}"
        return
    fi

    brew update && brew upgrade && brew cleanup -s

    echo -e "\n${GREEN}Updating Brew Casks${CLEAR}"
    brew outdated --cask && brew upgrade --cask && brew cleanup -s

    echo -e "\n${GREEN}Brew Diagnostics${CLEAR}"
    brew doctor && brew missing
}

update-vscode() {
    echo -e "\n${GREEN}Updating VSCode Extensions${CLEAR}"

    if ! command -v code &>/dev/null; then
        echo -e "${RED}VSCode is not installed.${CLEAR}"
        return
    fi

    code --update-extensions
}

update-office() {
    echo -e "\n${GREEN}Updating MS-Office${CLEAR}"

    local MS_OFFICE_UPDATE='/Library/Application Support/Microsoft/MAU2.0/Microsoft AutoUpdate.app/Contents/MacOS/msupdate'
    if [ ! -f "$MS_OFFICE_UPDATE" ]; then
        echo -e "${RED}MS-Office update utility is not installed.${CLEAR}"
        return
    fi

    "$MS_OFFICE_UPDATE" --install
}

update-gem() {
    echo -e "\n${GREEN}Updating Gems${CLEAR}"

    if ! command -v gem &>/dev/null; then
        echo -e "${RED}Gem is not installed.${CLEAR}"
        return
    fi

    gem update --user-install && gem cleanup --user-install
}

update-npm() {
    echo -e "\n${GREEN}Updating Npm Packages${CLEAR}"

    if ! command -v npm &>/dev/null; then
        echo -e "${RED}Npm is not installed.${CLEAR}"
        return
    fi

    npm update -g
}

update-yarn() {
    echo -e "\n${GREEN}Updating Yarn Packages${CLEAR}"

    if ! command -v yarn &>/dev/null; then
        echo -e "${RED}Yarn is not installed.${CLEAR}"
        return
    fi

    yarn upgrade --latest
}

update-pip3() {
    echo -e "\n${GREEN}Updating Python 3.x pips${CLEAR}"

    if ! command -v python3 &>/dev/null || ! command -v pip3 &>/dev/null; then
        echo -e "${RED}Python 3 or pip3 is not installed.${CLEAR}"
        return
    fi

    # python3 -c "import pkg_resources; from subprocess import call; packages = [dist.project_name for dist in pkg_resources.working_set]; call('pip3 install --upgrade ' + ' '.join(packages), shell=True)"
    pip3 list --outdated --format=columns | grep -v '^\-e' | cut -d = -f 1 | xargs -n1 pip3 install -U
}

update-app_store() {
    echo -e "\n${GREEN}Updating App Store Applications${CLEAR}"

    if ! command -v mas &>/dev/null; then
        echo -e "${RED}mas is not installed.${CLEAR}"
        return
    fi

    mas outdated | while read -r app; do mas upgrade "$app"; done
}

update-macos() {
    echo -e "\n${GREEN}Updating Mac OS${CLEAR}"
    softwareupdate -i -a
}

update-all() {
    local PING_IP=8.8.8.8
    if ping -q -W 1 -c 1 $PING_IP &> /dev/null; then
        update-brew
        update-office
        update-vscode
        update-gem
        update-npm
        update-yarn
        update-pip3
        update-app_store
        update-macos
    else
        echo -e "${RED}Internet Disabled!!!${CLEAR}"
    fi
}

# COMMENT OUT IF SOURCING
update-all
