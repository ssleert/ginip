#!/bin/bash

function generateFile() {
	local -r file="$1"

	rm -v "${file}"
	for i in {0..1000}; do
		echo "[package${i}]" >> "${file}"
		
		if (( i > 24)); then
			echo "name= util$(((i * 124891)))" >> "${file}"
		elif (( i > 49 )); then
			echo "  name =util$(((i * 124891)))" >> "${file}"
		elif (( i > 74 )); then
			echo "   name=util$(((i * 124891)))" >> "${file}"
			echo ""
		else 
			echo " name = util$(((i * 124891)))" >> "${file}"
		fi

		echo "      version = $(((i * 12489)))" >> "${file}"

		echo " ;asd = basd" >> "${file}"
		echo " ## asd = bhq" >> "${file}"
		echo "" >> "${file}"

	done
}

function confirmFile() {
	local -r utility="$1"
	local -r file="$2"
	
	for i in {0..1000}; do
		if [[ "util$(((i * 124891)))" = "$(${utility} ${file} package${i} name)" ]]; then
			 echo -e "\033[0;31m test complete \033[0m"
		fi

		if [[ "$(((i * 12489)))" = "$(${utility} ${file} package${i} version)" ]]; then
			echo -e "\033[0;31m test complete \033[0m"
		fi
	done
}

function main() {
	#generateFile "generated.ini"
	confirmFile "./inip" "generated.ini"
}

main "$@"
