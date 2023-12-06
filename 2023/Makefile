# 4h is time difference, plus an hour so I don't have to wait for it to drop
date=$(shell date -v +4H +"%d")
days := $(shell seq -f '%02g' 01 25)

.DEFAULT_GOAL := today

.PHONY: today
today: $(date)/part2.go $(date)/input.txt

%/part2.go: %/part1.go | %
	cp $< $@

# Specify these targets explicitly so they don't get picked up as intermediate ones.
$(patsubst %,%/part1.go,$(days)):

%/part1.go: | %
	cp template.go $@

%/input.txt: | %
	aocd $| > $@

$(days):
	mkdir -p $@
