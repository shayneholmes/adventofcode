date=$(shell date +"%d")

part2: | $(date)/part2.go

$(date)/part2.go: $(date)/part1.go | $(date)
	cp $< $@

$(date)/part1.go: | $(date)
	cp template.go $@

$(date):
	mkdir -p $(date)
