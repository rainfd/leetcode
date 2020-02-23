#!/bin/bash

if [[ $# -ne 4 ]]; then
  echo "Usage: new.sh [language] [level] [number] [link]
               level:  1 easy 2 medium 3 hard
               number: order number
       Example: ./new.sh Go 1 1 https://leetcode-cn.com/problems/two-sum/"
  exit 1
fi

language=$1
level=$2
number=$3
link=$4

case $language in
  "c") ext=".c";;
  "go") ext=".go";;
  "python") ext=".py";;
  "bash") ext=".sh";;
  "java") ext=".java";;
esac

case $level in 
  1) level="Easy";;
  2) level="Meidum";;
  3) level="Hard";;
esac

title="$(echo ${link} | grep -o -E '/problems/[a-z-]*/' | awk -F/ '{ print $3 }')" # two-sum
# title case
proname="$(echo ${title} | awk -F- '{for(j=1;j<=NF;j++){ $j=toupper(substr($j,1,1)) substr($j,2) }} 1')"
dname="${number}.${proname// /_}"
filename="${proname// /}$ext"

# echo "$title" two-sum
# echo "$proname" Two Sum
# echo "$filename" TwoSum.go
# echo "$dname" 1.Two_Sum

echo "Deirectory ${language}/${dname} created."
mkdir -p "${language}/${dname}"
echo "File ${language}/${dname}/${filename} created."
touch "${language}/${dname}/${filename}"

# Add item to README.md
grep "${link}" README.md && echo "Problem ${number} already exists." && exit
cp README.md _README.md
# head -n 8 README.md > heads
tail +9 README.md > problems
new="| ${number} | [${proname}]($link) | ${language} | ${level} |"
echo "${new}" >> problems
cat heads > README.md
cat problems | sort -n -k 2 > README.md
rm problems

# # echo "${new}"
# echo ${problems}
# sorted_problems=$(cat <(echo "${probelms}") <(echo "${new}") | sort -n -k 2)
# echo ${sorted_problems}
# cat <(echo "${heads}") <(echo "${problems}")
