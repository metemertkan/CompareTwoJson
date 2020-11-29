# CompareTwoJson

It is an example for comparing 2 json files.
To run in windows: execute "go build" and "./pvh.exe"

For sample json https://pokeapi.co/ has been used.

1-Program creates example json file with api calls, then reverses the array and saves as follows : "pokemons1.json" "pokemons2.json"
2-readPokemonsAllIntoMemory -> this method reads all json file into memory
3-readPokemonsInMemoryOptimisedWay -> this method reads file as chunks. Every pokemon item is a chunks. So It does not loads whole file into memory
4-compareTwoJson -> This method makes a simple comparison. Worst case : O(n^2)
5-compareTwoJsonWithMap -> This one uses hashmap to acces items so it reduces complexity. Worst case : O(nlogn)
6-compareTwoJsonWithMapParallel -> This one uses go routines iterate files. 

Example output for 500 Pokemons:

Reading Pokemons took  2.156518s
Reading Pokemons Optimised took  2.1165867s
========================
Simple Compare Started
Result = true
And it took  388.0321ms
Simple Compare Finished
========================
========================
Map Compare Started
Result = true
And it took  261.0324ms
Map Compare Finished
========================
========================
Parallel Compare Started
Result = true
And it took  77.9666ms
Parallel Compare Finished
========================
