#!/bin/sh

###################################################################################################
# File: easy_install.sh
# Description: This script installs the update-all.sh script and sets up the necessary environment.
###################################################################################################

###################################################################################################
# Global Variables & Constants
###################################################################################################
# Exit the script immediately if any command fails or if an unset variable is used
set -eu

readonly FILE_NAME="update-all.sh"
readonly UPDATE_SCRIPT_SOURCE_URL="https://raw.githubusercontent.com/andmpel/MacOS-All-In-One-Update-Script/HEAD/${FILE_NAME}"

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
    _local_bin=""
    case "${ADJUSTED_ID}" in
    darwin)
        _rc="${HOME}/.zshrc"
        _local_bin="${HOME}/.local/bin"
        ;;
    *)
        print_err "Error: Unsupported or unrecognized distribution ${ADJUSTED_ID}"
        exit 1
        ;;
    esac

    # Ensure HOME is set
    if [ -z "${HOME:-}" ]; then
        print_err "Error: HOME environment variable is not set."
        exit 1
    fi

    # Create local bin directory if it doesn't exist
    if ! mkdir -p "${_local_bin}"; then
        print_err "Error: Failed to create directory ${_local_bin}."
        exit 1
    fi

    # Download the update script
    if ! curl -fsSLo "${_local_bin}/update" "${UPDATE_SCRIPT_SOURCE_URL}"; then
        print_err "Error: Failed to download update script from ${UPDATE_SCRIPT_SOURCE_URL}."
        exit 1
    fi

    # Make the script executable
    if ! chmod +x "${_local_bin}/update"; then
        print_err "Error: Failed to make ${_local_bin}/update executable."
        exit 1
    fi

    # Prepare the PATH update block
    UPDATE_SOURCE_STR="
# Local Bin
if [ -d \"${_local_bin}\" ]; then
    export PATH=\"\${PATH}:${_local_bin}\"
fi
"

    # Update or create the rc file
    if [ -f "${_rc}" ]; then
        case ":${PATH}:" in
        *":${_local_bin}:"*) ;; # Already in PATH, do nothing
        *)
            println "==> Updating ${_rc} for ${ADJUSTED_ID}..."
            println "${UPDATE_SOURCE_STR}" >>"${_rc}"
            ;;
        esac
    else
        println "==> Profile not found. ${_rc} does not exist."
        println "==> Creating the file ${_rc}... Please note that this may not work as expected."
        if ! touch "${_rc}"; then
            print_err "Error: Failed to create ${_rc}."
            exit 1
        fi
        println "${UPDATE_SOURCE_STR}" >>"${_rc}"
    fi

    println ""
    println "==> Close and reopen your terminal to start using 'update'"
    println "    OR"
    println "==> Run the following to use it now:"
    println ">>> source ${_rc} # This loads update"
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
