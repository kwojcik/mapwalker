# mapwalker
An exploration into what happens when you insert into a map that you're iterating over.

Currently, three experiments are run:
1. A real golang map
2. A simulated map that never needs to grow
3. A real golang map with a large pre-allocated capacity

# Install
```
make
```

# Run
```
./mapwalker
```

# Distributions of types of map
```
kevin$ ./mapwalker

Results for map
Initial size: 1024
Iterations: 1000
Final size:
        Average 2520.973000
        Stddev  13.164812
Distribution of final size of map
2478-2487  0.4%   ▍                      4
2487-2496  1.9%   █▋                     19
2496-2504  8.3%   ██████▋                83
2504-2513  18.8%  ███████████████        188
2513-2522  22%    █████████████████▌     220
2522-2531  25.2%  ████████████████████▏  252
2531-2540  14.9%  ███████████▉           149
2540-2548  7%     █████▋                 70
2548-2557  1.2%   █                      12
2557-2566  0.3%   ▎                      3

Results for NoGrowMap
Initial size: 1024
Iterations: 1000
Final size:
        Average 2781.450000
        Stddev  28.846135
Distribution of final size of map
2675-2696  0.3%   ▎                      3
2696-2718  0.6%   ▌                      6
2718-2739  6.7%   ████▊                  67
2739-2761  16.1%  ███████████▍           161
2761-2782  26%    ██████████████████▍    260
2782-2803  28.4%  ████████████████████▏  284
2803-2825  15.6%  ███████████            156
2825-2846  5.1%   ███▋                   51
2846-2868  1%     ▊                      10
2868-2889  0.2%   ▏                      2

Results for sparse map
Initial size: 1024
Iterations: 1000
Final size:
        Average 2782.760000
        Stddev  27.400190
Distribution of final size of map
2693-2710  0.3%   ▎                      3
2710-2727  1.8%   █▋                     18
2727-2744  5.4%   ████▋                  54
2744-2761  12.8%  ███████████▏           128
2761-2778  22.9%  ███████████████████▉   229
2778-2794  23.1%  ████████████████████▏  231
2794-2811  17.6%  ███████████████▎       176
2811-2828  11.7%  ██████████▏            117
2828-2845  3.4%   ███                    34
2845-2862  1%     ▉                      10
```

# Digging further into normal maps
Can we guess the final size based on the initial size? Let's take some data

```
for x in `seq 100 100 50000`; do echo -n "$x "; ./mapwalker --rawResults --onlyMap --iterations 1 --initial $x; done | awk '{print $1 " " $2/$1}' > data
for x in `seq 50000 10000 500000`; do echo -n "$x "; ./mapwalker --rawResults --onlyMap --iterations 1 --initial $x; done | awk '{print $1 " " $2/$1}' >> data
echo 'set terminal png; plot "data" using 1:2 title "Final size ratio vs initial size"' | gnuplot > test.png
```

Here we see the ratio of the final size to the initial size, as a function of initial size. It's
not what I expected...
![normal map final size ratio](https://github.com/kwojcik/mapwalker/blob/master/images/1.png)
