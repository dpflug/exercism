##
# dpflug's Exercism Exercises
#
# @file
# @version 0.1

GODIRS = $(wildcard go/*/.)

all: $(GODIRS)

.PHONY: all $(GODIRS)

$(GODIRS):
	cd $@ ; go fmt . ; go test -v --bench . --benchmem

# end
