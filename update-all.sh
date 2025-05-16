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

check_command() {
    command_name="$1"

    if ! command -v "${command_name}" >/dev/null 2>&1; then
        print_err "${command_name} is not installed."
        return 1
    fi

    return 0
}

update_brew() {
    println "Updating Brew Formula's"

    if ! check_command brew; then
        return
    fi

    brew update && brew upgrade && brew cleanup -s

    println "Brew Diagnostics"
    brew doctor && brew missing
}

update_vscode() {
    println "Updating VSCode Extensions"

    if ! check_command code; then
        return
    fi

    code --update-extensions
}

update_gem() {
    println "Updating Gems"

    # Get the path of the `gem` command
    GEM_PATH=$(which gem)

    # Check if the path does not match the expected path
    if [ "$GEM_PATH" = "/usr/bin/gem" ]; then
        print_err "gem is not installed."
        return
    fi

    gem update --user-install && gem cleanup --user-install
}

update_npm() {
    println "Updating Npm Packages"

    if ! check_command npm; then
        return
    fi

    npm update -g
}

update_yarn() {
    println "Updating Yarn Packages"

    if ! check_command yarn; then
        return
    fi

    yarn upgrade --latest
}

update_cargo() {
    println "Updating Rust Cargo Crates"

    if ! check_command cargo; then
        return
    fi

    cargo install --list | grep -E '^[a-z0-9_-]+ v[0-9.]+:$' | cut -f1 -d' ' | xargs cargo install
}

update_app_store() {
    println "Updating App Store Applications"

    if ! check_command mas; then
        return
    fi

    mas upgrade
}

update_macos() {
    println "Updating MacOS"
    softwareupdate -i -a
}

check_internet() {
    if ! check_command curl; then
        print_err "Error: curl is required but not installed. Please install curl."
        return 1
    fi

    # Check internet connection by pinging a reliable server
    TEST_URL="https://www.google.com"

    # Use curl to check the connection
    TEST_RESP=$(curl -Is --connect-timeout 5 --max-time 10 "${TEST_URL}" 2>/dev/null | head -n 1)

    # Check if response is empty
    if [ -z "${TEST_RESP}" ]; then
        print_err "No Internet Connection!!!"
        return 1
    fi

    # Check for "200" in the response
    if ! printf "%s" "${TEST_RESP}" | grep -q "200"; then
        print_err "Internet is not working!!!"
        return 1
    fi

    return 0
}

update_all() {
    # Check if internet is available
    if ! check_internet; then
        exit 1
    fi

    update_brew
    update_vscode
    update_gem
    update_npm
    update_yarn
    update_cargo
    update_app_store
    update_macos
}

# COMMENT OUT IF SOURCING
update_all
