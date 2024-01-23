#!/bin/bash
find . -type f -name "*.sh" | sed 's/\.sh$//' | sed 's/^\.///' | awk -F'/' '{print $NF}' sort -r | sed 's/\$$//'