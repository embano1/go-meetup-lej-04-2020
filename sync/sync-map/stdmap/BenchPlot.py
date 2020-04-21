#!/usr/bin/env python

# You'll need:
# $ pip install pandas matplotlib
# $ make data.tsv
import datetime
import pandas as pd
import matplotlib.pyplot as plt

if __name__ == "__main__":
    df = pd.read_csv("data.tsv", sep="\t", header=None)
    # 0  BenchmarkStoreStandardMap-8  209
    # 1      BenchmarkStoreSyncMap-8  650
    # 2     BenchmarkDeleteRegular-8  158
    # 3        BenchmarkDeleteSync-8  217
    # 4  BenchmarkLoadRegularFound-8  107
    # 5     BenchmarkLoadSyncFound-8  196

    data = {
        "builtin": df[~df[0].str.contains("Sync")].iloc[:, 1].values,
        "sync": df[df[0].str.contains("Sync")].iloc[:, 1].values,
    }
    index = ["store", "delete", "load"]
    df = pd.DataFrame(data, index=index)
    #         builtin  sync
    # store       209   650
    # delete      158   217
    # load        107   196

    ax = df.plot(
        kind="bar",
        figsize=(10, 6),
        title="Go Benchmark builtin map and sync.Map (%s)" % datetime.date.today(),
    )
    ax.set_ylabel("ns/op")
    ax.set_xticklabels(index, rotation=45)
    fig = plt.gcf()
    fig.savefig('output.png')

