#!/bin/bash
find . -type f -name "*.sh" | sed 's/\.sh$//' | sed 's/^\.///' | awk -f'/' '{print $NF}' sort -r | sed 's/\$$//'|