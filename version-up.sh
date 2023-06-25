#!/bin/bash

echo $(git tag -l)

V_STR=$(git describe --tags --abbrev=0)
# V_STR_ARRAY=(`echo $V_STR | tr "." "\n"`)

V_STR_ARRAY=(`echo $(git tag -l) | tr " " "\n"`)

for (( i=0; i<${#V_STR_ARRAY[@]}; i++ )); do
    echo "${V_STR_ARRAY[i]}"
done

# VERSION=""

# for (( i=0; i<${#V_STR_ARRAY[@]}; i++ )); do
#     if [ $i == 2 ]; then
#         VERSION="${VERSION}$((${V_STR_ARRAY[i]} + 1))"
#     else
#         VERSION="${VERSION}${V_STR_ARRAY[i]}."
#     fi
# done

# git tag $VERSION