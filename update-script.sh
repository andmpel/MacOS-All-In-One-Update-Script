#!/bin/bash
#Text Color Variables
LG='\033[1;32m' #Light Green
NC='\033[0m' #No Color
echo -e "${LG}Updating Brew Formula's${NC}"
brew update
brew upgrade
brew cleanup -s

echo -e "\n${LG}Updating Brew Casks${NC}"
brew cask outdated
brew cask cleanup
brew cask upgrade

echo -e "\n${LG}Brew Diagnostics${NC}"
brew doctor
brew missing

echo -e "\n${LG}Updating App Store Applications${NC}"
mas outdated
mas upgrade

echo -e "\n${LG}Updating Atom${NC}"
apm-beta upgrade -c false

echo -e "\n${LG}Updating npm Packages${NC}"
npm update -g

gem update

echo -e "\n${LG}Updating Python 2.7.X pips${NC}"
pip2 freeze - local | grep -v ‘^\-e’ | cut -d = -f 1 | xargs pip2 install -U

echo -e "\n${LG}Updating Python 3.X pips${NC}"
pip3 freeze - local | grep -v ‘^\-e’ | cut -d = -f 1 | xargs pip3 install -U

echo -e "\n${LG}Updating Mac OS${NC}"
softwareupdate -i -a
