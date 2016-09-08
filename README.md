# mapwalker
An exploration into what happens when you insert into a map that you're iterating over.

Currently, two experiments are run: one with a real golang map and one with a simulation of
a map that never needs to grow.

# install
```
make
```

# run
```
./mapwalker
```

# results
```
kevin$ ./mapwalker

Results for map
Initial size: 1024
Iterations: 1000
Final size:
        Average 2520.300000
        Stddev  12.861104
Distribution of final size of map
2482-2487  0.4%   ▌                      4
2487-2493  0.8%   █                      8
2493-2498  3.2%   ███▋                   32
2498-2504  5.9%   ██████▋                59
2504-2509  10.2%  ███████████▋           102
2509-2515  11.6%  █████████████▏         116
2515-2520  17.7%  ████████████████████▏  177
2520-2526  16.1%  ██████████████████▏    161
2526-2531  15.2%  █████████████████▏     152
2531-2537  8.6%   █████████▊             86
2537-2542  6.2%   ███████▏               62
2542-2548  2.5%   ██▉                    25
2548-2553  0.9%   █▏                     9
2553-2559  0.3%   ▍                      3
2559-2564  0.4%   ▌                      4

Results for NoGrowMap
Initial size: 1024
Iterations: 1000
Final size:
        Average 2781.450000
        Stddev  28.846135
Distribution of final size of map
2675-2689  0.1%   ▏                      1
2689-2704  0.3%   ▎                      3
2704-2718  0.5%   ▌                      5
2718-2732  2.9%   ██▊                    29
2732-2746  8.2%   ███████▉               82
2746-2761  11.7%  ███████████▏           117
2761-2775  16.1%  ███████████████▌       161
2775-2789  20.9%  ████████████████████▏  209
2789-2803  17.4%  ████████████████▋      174
2803-2818  11.9%  ███████████▍           119
2818-2832  5.8%   █████▋                 58
2832-2846  3%     ██▉                    30
2846-2860  0.9%   ▉                      9
2860-2875  0.2%   ▏                      2
2875-2889  0.1%   ▏                      1
```
