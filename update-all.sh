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

update_brew() {
    echo "${GREEN}Updating Brew Formula's${CLEAR}"

    if ! command -v brew >/dev/null 2>&1; then
        echo "${RED}Brew is not installed.${CLEAR}"
        return
    fi

    brew update && brew upgrade && brew cleanup -s

    echo "\n${GREEN}Updating Brew Casks${CLEAR}"
    brew outdated --cask && brew upgrade --cask && brew cleanup -s

    echo "\n${GREEN}Brew Diagnostics${CLEAR}"
    brew doctor && brew missing
}

update_vscode() {
    echo "\n${GREEN}Updating VSCode Extensions${CLEAR}"

    if ! command -v code >/dev/null 2>&1; then
        echo "${RED}VSCode is not installed.${CLEAR}"
        return
    fi

    code --update-extensions
}

update_office() {
    echo "\n${GREEN}Updating MS-Office${CLEAR}"

    readonly MS_OFFICE_UPDATE='/Library/Application Support/Microsoft/MAU2.0/Microsoft AutoUpdate.app/Contents/MacOS/msupdate'
    if [ ! -f "$MS_OFFICE_UPDATE" ]; then
        echo "${RED}MS-Office update utility is not installed.${CLEAR}"
        return
    fi

    "$MS_OFFICE_UPDATE" --install
}

update_gem() {
    echo "\n${GREEN}Updating Gems${CLEAR}"

    if ! command -v gem >/dev/null 2>&1; then
        echo "${RED}Gem is not installed.${CLEAR}"
        return
    fi

    gem update --user-install && gem cleanup --user-install
}

update_npm() {
    echo "\n${GREEN}Updating Npm Packages${CLEAR}"

    if ! command -v npm >/dev/null 2>&1; then
        echo "${RED}Npm is not installed.${CLEAR}"
        return
    fi

    npm update -g
}

update_yarn() {
    echo "\n${GREEN}Updating Yarn Packages${CLEAR}"

    if ! command -v yarn >/dev/null 2>&1; then
        echo "${RED}Yarn is not installed.${CLEAR}"
        return
    fi

    yarn upgrade --latest
}

update_pip3() {
    echo "\n${GREEN}Updating Python 3.x pips${CLEAR}"

    if ! command -v python3 >/dev/null 2>&1 || ! command -v pip3 >/dev/null 2>&1; then
        echo "${RED}Python 3 or pip3 is not installed.${CLEAR}"
        return
    fi

    # python3 -c "import pkg_resources; from subprocess import call; packages = [dist.project_name for dist in pkg_resources.working_set]; call('pip3 install --upgrade ' + ' '.join(packages), shell=True)"
    pip3 list --outdated --format=columns | grep -v '^\-e' | cut -d = -f 1 | xargs -n1 pip3 install -U
}

update_app_store() {
    echo "\n${GREEN}Updating App Store Applications${CLEAR}"

    if ! command -v mas >/dev/null 2>&1; then
        echo "${RED}mas is not installed.${CLEAR}"
        return
    fi

    mas outdated | while read -r app; do mas upgrade "$app"; done
}

update_macos() {
    echo "\n${GREEN}Updating Mac OS${CLEAR}"
    softwareupdate -i -a
}

update_all() {
    readonly PING_IP=8.8.8.8
    if ping -q -W 1 -c 1 $PING_IP >/dev/null 2>&1; then
        update_brew
        update_office
        update_vscode
        update_gem
        update_npm
        update_yarn
        update_pip3
        update_app_store
        update_macos
    else
        echo "${RED}Internet Disabled!!!${CLEAR}"
    fi
}

# COMMENT OUT IF SOURCING
update_all
