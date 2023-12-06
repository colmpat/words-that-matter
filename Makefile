SERVICES := $(wildcard services/*)

all: $(SERVICES)

$(SERVICES):
	@echo "Building: $@"
	$(MAKE) -C $@

clean:
	@echo "Cleaning..."
	@for dir in $(SERVICES); do \
		$(MAKE) -C $$dir clean; \
	done

.PHONY: all $(SERVICES)
