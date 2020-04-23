SHELL := /bin/bash
SRCS = *.go

output.png: data.tsv requirements.txt
	# Temporary install dependencies in an env.
	python -m venv .tmpenv && source .tmpenv/bin/activate && \
		pip install -r requirements.txt && \
		python BenchPlot.py && \
		deactivate && rm -rf .tmpenv

data.tsv: $(SRCS)
	go test -bench . | grep ^Benchmark | awk '{print $$1"\t"$$3}' > $@

.PHONY: clean
clean:
	rm -f output.png
	rm -f data.tsv
	rm -rf .tmpenv
