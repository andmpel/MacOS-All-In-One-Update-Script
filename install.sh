#!/bin/sh

###################################################################################################
# File: install.sh
# Description: This script create an alias by downloading a
#              predefined update-all.sh file from a GitHub repository and
#              integrating it into the user's shell configuration
#              files (~/.zshrc)
###################################################################################################

###################################################################################################
# Global Variables & Constants
###################################################################################################
# Exit the script immediately if any command fails
set -e

readonly FILE_NAME="update-all.sh"
readonly UPDATE_SCRIPT_SOURCE_URL="https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/${FILE_NAME}"

UPDATE_SOURCE_STR=$(
    cat <<EOF

# System Update
update() {
    # Check if curl is available
    if ! command -v curl >/dev/null 2>&1; then
        echo "Error: curl is required but not installed. Please install curl."
        exit 1
    fi

    readonly TEST_URL="https://www.google.com"
    readonly TIMEOUT=2

    # Check if the internet is reachable
    if ! curl -s --max-time \${TIMEOUT} --head --request GET \${TEST_URL} | grep "200 OK" >/dev/null; then
        echo "Internet Disabled!!!"
        exit 1
    fi

    curl -fsSL ${UPDATE_SCRIPT_SOURCE_URL} | zsh
}
EOF
)

###################################################################################################
# Functions
###################################################################################################

# Function: println
# Description: Prints each argument on a new line, suppressing any error messages.
println() {
    printf "%s\n" "$*" 2>/dev/null
}

# Function: print_err
# Description: Prints each argument as error messages.
print_err() {
    printf "%s\n" "$*" >&2
}

# Function: update_rc
# Description: Update shell configuration files
update_rc() {
    _rc=""
    case $ADJUSTED_ID in
    darwin)
        _rc="${HOME}/.zshrc"
        ;;
    *)
        print_err "Error: Unsupported or unrecognized distribution ${ADJUSTED_ID}"
        exit 1
        ;;
    esac

    # Check if `alias update='sudo sh ${HOME}/.update.sh'` is already defined, if not then append it
    if [ -f "${_rc}" ]; then
        if ! grep -qxF "${UPDATE_SOURCE_STR}" "${_rc}"; then
            println "==> Updating ${_rc} for ${ADJUSTED_ID}..."
            println "${UPDATE_SOURCE_STR}" >>"${_rc}"
        fi
    else
        # Notify if the rc file does not exist
        println "==> Profile not found. ${_rc} does not exist."
        println "==> Creating the file ${_rc}... Please note that this may not work as expected."
        # Create the rc file
        touch "${_rc}"
        # Append the sourcing block to the newly created rc file
        println "${UPDATE_SOURCE_STR}" >>"${_rc}"
    fi

    println ""
    println "==> Close and reopen your terminal to start using 'update' alias"
    println "    OR"
    println "==> Run the following to use it now:"
    println ">>> source ${_rc} # This loads update alias"
}

###################################################################################################
# Main Script
###################################################################################################

OS=$(uname)

case ${OS} in
Darwin)
    ADJUSTED_ID="darwin"
    ;;
*)
    print_err "Error: Unsupported or unrecognized OS distribution: ${OS}"
    exit 1
    ;;
esac

# Check if curl is available
if ! command -v curl >/dev/null 2>&1; then
    print_err "Error: curl is required but not installed. Please install curl."
    exit 1
fi

# Update the rc (.zshrc) file for `update` alias
update_rc
