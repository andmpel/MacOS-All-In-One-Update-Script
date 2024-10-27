#!/bin/sh

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

println() {
    printf "\n${GREEN}%s${CLEAR}\n" "$*" 2>/dev/null
}

print_err() {
    printf "\n${RED}%s${CLEAR}\n" "$*" >&2
}

update_brew() {
    println "Updating Brew Formula's"

    if ! command -v brew >/dev/null 2>&1; then
        print_err "Brew is not installed."
        return
    fi

    brew update && brew upgrade && brew cleanup -s

    println "Brew Diagnostics"
    brew doctor && brew missing
}

update_vscode() {
    println "Updating VSCode Extensions"

    if ! command -v code >/dev/null 2>&1; then
        print_err "VSCode is not installed."
        return
    fi

    code --update-extensions
}

update_office() {
    println "Updating MS-Office"

    readonly MS_OFFICE_UPDATE='/Library/Application Support/Microsoft/MAU2.0/Microsoft AutoUpdate.app/Contents/MacOS/msupdate'
    if [ ! -f "${MS_OFFICE_UPDATE}" ]; then
        print_err "MS-Office update utility is not installed."
        return
    fi

    "${MS_OFFICE_UPDATE}" --install
}

update_gem() {
    println "Updating Gems"

    if ! command -v gem >/dev/null 2>&1; then
        print_err "Gem is not installed."
        return
    fi

    gem update --user-install && gem cleanup --user-install
}

update_npm() {
    println "Updating Npm Packages"

    if ! command -v npm >/dev/null 2>&1; then
        print_err "Npm is not installed."
        return
    fi

    npm update -g
}

update_yarn() {
    println "Updating Yarn Packages"

    if ! command -v yarn >/dev/null 2>&1; then
        print_err "Yarn is not installed."
        return
    fi

    yarn upgrade --latest
}

update_pip3() {
    println "Updating Python 3.x pips"

    if ! command -v python3 >/dev/null 2>&1 || ! command -v pip3 >/dev/null 2>&1; then
        print_err "Python3 or pip3 is not installed."
        return
    fi

    pip3 list --outdated --format=columns | grep -v '^\-e' | cut -d = -f 1 | xargs -n1 pip3 install -U
}

update_cargo() {
    println "Updating Rust Cargo Crates"

    if ! command -v cargo >/dev/null 2>&1; then
        print_err "Rust/Cargo is not installed."
        return
    fi

    cargo install "$(cargo install --list | grep -E '^[a-z0-9_-]+ v[0-9.]+:$' | cut -f1 -d' ')"
}

update_app_store() {
    println "Updating App Store Applications"

    if ! command -v mas >/dev/null 2>&1; then
        print_err "mas is not installed."
        return
    fi

    mas outdated | while read -r app; do mas upgrade "${app}"; done
}

update_macos() {
    println "Updating MacOS"
    softwareupdate -i -a
}

update_all() {
    readonly TEST_URL="https://www.google.com"
    readonly TIMEOUT=2

    # Check if curl is available
    if command -v curl >/dev/null 2>&1; then
        # Check if the internet is reachable
        if ! curl -s --max-time ${TIMEOUT} --head --request GET ${TEST_URL} | grep "200 OK" >/dev/null; then
            print_err "Internet Disabled!!!"
            exit 1
        fi
    fi

    update_brew
    update_office
    update_vscode
    update_gem
    update_npm
    update_yarn
    update_pip3
    update_cargo
    update_app_store
    update_macos
}

# COMMENT OUT IF SOURCING
update_all
