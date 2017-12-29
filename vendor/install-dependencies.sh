#!/bin/sh
export PATH=/usr/bin:$PATH
export PATH=/bin:$PATH
export PATH=/mingw64/bin:$PATH

pacman-key --init
pacman-key --populate
pacman-key --refresh-keys

pacman -S mingw-w64-x86_64-gcc --noconfirm
pacman -S make --noconfirm
pacman -S perl --noconfirm
pacman -S diffutils --noconfirm