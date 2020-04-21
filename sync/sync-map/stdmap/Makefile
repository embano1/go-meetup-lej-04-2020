SRCS = *.go

output.png: data.tsv
	python BenchPlot.py

data.tsv: $(SRCS)
	go test -bench . | grep ^Benchmark | awk '{print $$1"\t"$$3}' > $@

.PHONY: clean
clean:
	rm -f output.png
	rm -f data.tsv

