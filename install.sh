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
# Exit the script immediately if any command fails or if an unset variable is used
set -eu

readonly FILE_NAME="update-all.sh"
readonly UPDATE_SCRIPT_SOURCE_URL="https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/${FILE_NAME}"

UPDATE_SOURCE_STR=$(
    cat <<EOF

# System Update
update() {
    # Check if curl is available
    if ! command -v curl >/dev/null 2>&1; then
        echo "Error: curl is required but not installed. Please install curl." >&2
        return
    fi

    # Check internet connection by pinging a reliable server
    TEST_URL="https://www.google.com"

    # Use curl to check the connection
    TEST_RESP=\$(curl -Is --connect-timeout 5 --max-time 10 "\${TEST_URL}" 2>/dev/null | head -n 1)

    # Check if response is empty
    if [ -z "\${TEST_RESP}" ]; then
        echo "No Internet Connection!!!" >&2
        return
    fi

    # Check for "200" in the response
    if ! printf "%s" "\${TEST_RESP}" | grep -q "200"; then
        echo "Internet is not working!!!" >&2
        return
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
    case "${ADJUSTED_ID}" in
    darwin)
        _rc="${HOME}/.zshrc"
        ;;
    *)
        print_err "Error: Unsupported or unrecognized distribution ${ADJUSTED_ID}"
        exit 1
        ;;
    esac

    # Check if `update` function is already defined, if not then append it
    if [ -f "${_rc}" ]; then
        if ! awk '/^update\(\) {/,/^}/' "${_rc}" | grep -q 'curl'; then
            println "==> Updating ${_rc} for ${ADJUSTED_ID}..."
            println "${UPDATE_SOURCE_STR}" >>"${_rc}"
        fi
    else
        # Notify if the rc file does not exist
        println "==> Profile not found. ${_rc} does not exist."
        println "==> Creating the file ${_rc}... Please note that this may not work as expected."
        # Create the rc file
        if ! touch "${_rc}"; then
            print_err "Error: Failed to create ${_rc}."
            exit 1
        fi
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

# Check for required commands
for cmd in awk cat chmod curl mkdir touch uname; do
    if ! command -v "${cmd}" >/dev/null 2>&1; then
        print_err "Error: Required command '%s' not found in PATH.\n" "${cmd}"
        exit 1
    fi
done

if [ "$(id -u)" -eq 0 ]; then
    print_err "Warning: Running as root is not recommended."
fi

OS=$(uname)

case "${OS}" in
Darwin)
    ADJUSTED_ID="darwin"
    ;;
*)
    print_err "Error: Unsupported or unrecognized OS distribution: ${OS}"
    exit 1
    ;;
esac

# Update the rc (.zshrc) file for `update`
update_rc
