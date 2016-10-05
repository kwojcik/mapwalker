# mapwalker
What happens when you insert into a map that you're iterating over.

# Intro
If you are iterating over each element in a map, and you insert a new element into the map at each step, how long does it take you to reach the end?

I've never had to do this before, and it's not even supported in some languages, like Python:
```
>>> d = {'hello': 'world'}
>>> for e in d:
...     d['foo'] = 'bar'
...
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
RuntimeError: dictionary changed size during iteration
```

The position of an element in a hashtable is random assuming you have a good hash function, so each element that you insert has a chance of ending up ahead of or behind your iterator. The probability of either happening is a function of where the iterator is in the map. In the middle, there is a 50/50 chance of either happening. Near the beginning, it's much more likely that the element ends up ahead of you. Near the end, it's more likely that the element will end up behind you.

The other variable is how large the map is when you start iterating. If the map starts of with 4 elements when you start iterating, then the 5th element has a 80% chance of going ahead of you and 20% of going behind you. If you start iterating over a map of 999 elements, the 1000th has a 0.1% chance of going behind you and a 99.9% chance of going ahead of you.

But it's still not that simple. A map will typically dynamically resize itself when it starts to get full. The specifics about how a map growswill turn out to be the interesting part of this exploration.

# Experiments

To demonstrate the complexity, let's run three experiments to demonstrate that growing the hash table makes things a little complicated.

1. A real golang map
2. A simulated map that never needs to grow
3. A real golang map, with a large pre-allocated capacity

### To follow along at home
```
git clone https://github.com/kwojcik/mapwalker.git
cd mapwalker
make
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

# Thoughts so far
1. The real map grew about 2.5x as we iterated through it
2. The simulated map grew about 2.8x as we iterated through it.
3. The sparse map (preallocated large enough so as to not grow) behaved exactly like the simulated map.

The simulated and sparse map behaved pretty similarly. The averages and stddev were identical, but the distribution was a little different. I still think it's safe to say that our simulated map accurately models a map's behavior, minus growth.

The real map that had to grow actually ended up smaller than the simulated one 100% of the time, which is interesting. I expected the randomness of growth to give a much larger variance in final size. The randomness I'm referring to here is that during growth I expecte the element that were iterating on to end up in a random bucket in the hash table. Meaning sometimes after growing the iterator would jump to the end of the map and other times it would be at the very beginning.

I guess I was wrong, let's dig in further.

# Digging further into normal maps
Let's take some data on the iteration length vs the initial size of the map; maybe we guess the final size based on the initial size?

```
# Initial sizes 100, 200, 300, ..., 50000
for x in `seq 100 100 50000`; do
  echo -n "$x ";
  ./mapwalker --rawResults --onlyMap --iterations 1 --initial $x;
done | awk '{print $1 " " $2/$1}' > data

# Initial sizes 50000, 60000, 70000, ..., 500000
for x in `seq 50000 10000 500000`; do
  echo -n "$x ";
  ./mapwalker --rawResults --onlyMap --iterations 1 --initial $x;
done | awk '{print $1 " " $2/$1}' >> data

# Plot
echo 'set terminal png; plot "data" using 1:2 title "Final size ratio vs initial size"' | gnuplot > test.png
```

Here we see the ratio of the final size to the initial size, as a function of initial size. It's
not what I expected...
![normal map final size ratio](https://github.com/kwojcik/mapwalker/blob/master/images/1.png)
