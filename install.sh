#!/bin/zsh

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
readonly FILE_PATH="${HOME}/${FILE_NAME}"
readonly UPDATE_SCRIPT_SOURCE_URL="https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/${FILE_NAME}"

readonly UPDATE_ALIAS_SEARCH_STR="alias update='zsh ${FILE_PATH}'"

UPDATE_ALIAS_SOURCE_STR=$(
    cat <<EOF

# Alias for Update
${UPDATE_ALIAS_SEARCH_STR}
EOF
)

###################################################################################################
# Functions
###################################################################################################

# Function: println
# Description: Prints each argument on a new line, suppressing any error messages.
println() {
    command printf %s\\n "$*" 2>/dev/null
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
        println >&2 "Error: Unsupported or unrecognized distribution ${ADJUSTED_ID}"
        exit 1
        ;;
    esac

    # Check if `alias update='sudo sh ${HOME}/.update.sh'` is already defined, if not then append it
    if [ -f "${_rc}" ]; then
        if ! grep -qxF "${UPDATE_ALIAS_SEARCH_STR}" "${_rc}"; then
            println "=> Updating ${_rc} for ${ADJUSTED_ID}..."
            println "${UPDATE_ALIAS_SOURCE_STR}" >>"${_rc}"
        fi
    else
        # Notify if the rc file does not exist
        println "=> Profile not found. ${_rc} does not exist."
        println "=> Creating the file ${_rc}... Please note that this may not work as expected."
        # Create the rc file
        touch "${_rc}"
        # Append the sourcing block to the newly created rc file
        println "${UPDATE_ALIAS_SOURCE_STR}" >>"${_rc}"
    fi

    println ""
    println "=> Close and reopen your terminal to start using 'update' alias"
    println "   OR"
    println "=> Run the following to use it now:"
    println ">>> source ${_rc} # This loads update alias"
}

# Function: dw_file
# Description: Download file using wget or curl if available
dw_file() {
    # Check if curl is available
    if command -v curl >/dev/null 2>&1; then
        curl -fsSL -o "${FILE_PATH}" ${UPDATE_SCRIPT_SOURCE_URL}
    # Check if wget is available
    elif command -v wget >/dev/null 2>&1; then
        wget -O "${FILE_PATH}" ${UPDATE_SCRIPT_SOURCE_URL}
    else
        println >&2 "Error: Either install wget or curl"
        exit 1
    fi
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
    println >&2 "Error: Unsupported or unrecognized OS distribution ${ADJUSTED_ID}"
    exit 1
    ;;
esac

# Default behavior
_action="y"

# Check if the script is running in interactive mode, for non-interactive mode `_action` defaults to 'y'
if [ -t 0 ]; then
    # Interactive mode
    if [ -f "${FILE_PATH}" ]; then
        println "=> File already exists: ${FILE_PATH}"
        println "=> Do you want to replace it (default: y)? [y/n]: "
        # Read input, use default value if no input is given
        read -r _rp_conf
        _rp_conf="${_rp_conf:-${_action}}"
        _action="${_rp_conf}"
    fi
fi

if [ "${_action}" = "y" ]; then
    println "=> Updating the file: ${FILE_PATH}"
    # Download the necessary file from the specified source
    dw_file
    # Update the configuration file with the latest changes
    update_rc
elif [ "${_action}" = "n" ]; then
    println "=> Keeping existing file: ${FILE_PATH}"
else
    println >&2 "Error: Invalid input. Please check your entry and try again."
    exit 1
fi