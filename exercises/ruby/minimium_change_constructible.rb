def minimum_change_constructible(coins)
 coins.sort!
 minimum_change = 0
 for coin in coins
    if coin > minimum_change + 1
        break
    end
    minimum_change += coin
 end
 return minimum_change + 1
end

# Alternative 1: More Ruby-idiomatic version of the original
def minimum_change_constructible_v2(coins)
  coins.sort!
  minimum_constructible = 0
  
  coins.each do |coin|
    break if coin > minimum_constructible + 1
    minimum_constructible += coin
  end
  
  minimum_constructible + 1
end

# Alternative 2: Functional approach with reduce
def minimum_change_constructible_v3(coins)
  sorted_coins = coins.sort
  
  min_constructible = sorted_coins.reduce(0) do |acc, coin|
    # If we hit a gap, we want to stop the reduction
    return acc + 1 if coin > acc + 1
    acc + coin
  end
  
  min_constructible + 1
end

# Alternative 3: Dynamic Programming approach (less efficient but educational)
def minimum_change_constructible_dp(coins)
  return 1 if coins.empty?
  
  max_possible = coins.sum
  # dp[i] represents whether amount i can be constructed
  dp = Array.new(max_possible + 1, false)
  dp[0] = true  # We can always make 0 with no coins
  
  coins.each do |coin|
    # Work backwards to avoid using the same coin multiple times
    (coin..max_possible).each do |amount|
      dp[amount] = true if dp[amount - coin]
    end
  end
  
  # Find the first amount that cannot be constructed
  (1..max_possible + 1).each do |amount|
    return amount unless dp[amount]
  end
  
  max_possible + 1
end

# Alternative 4: Set-based approach (more intuitive but less efficient)
def minimum_change_constructible_set(coins)
  return 1 if coins.empty?
  
  constructible = Set.new([0])  # We can always make 0
  
  coins.each do |coin|
    new_amounts = Set.new
    constructible.each do |amount|
      new_amounts.add(amount + coin)
    end
    constructible.merge(new_amounts)
  end
  
  # Find the first positive amount not in the set
  amount = 1
  while constructible.include?(amount)
    amount += 1
  end
  
  amount
end

# Alternative 5: Optimized version with early termination
def minimum_change_constructible_optimized(coins)
  return 1 if coins.empty?
  
  sorted_coins = coins.sort
  min_constructible = 0
  
  sorted_coins.each do |coin|
    # Early termination: if the smallest coin is > 1, we can't make 1
    return 1 if min_constructible == 0 && coin > 1
    
    # If there's a gap, we found our answer
    return min_constructible + 1 if coin > min_constructible + 1
    
    min_constructible += coin
  end
  
  min_constructible + 1
end

# Alternative 6: One-liner functional approach
def minimum_change_constructible_oneliner(coins)
  coins.sort.reduce(0) { |sum, coin| coin > sum + 1 ? (return sum + 1) : sum + coin } + 1
end

require 'set'

# Test all versions
test_cases = [
  [1, 2, 5],
  [1, 1, 1, 1, 1],
  [1, 5, 1, 1, 1, 10, 15, 20, 100],
  [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
  [5, 7, 1, 1, 2, 3, 22],
  [],
  [2, 3, 4, 5],
  [1]
]

methods = [
  :minimum_change_constructible,
  :minimum_change_constructible_v2,
  :minimum_change_constructible_v3,
  :minimum_change_constructible_dp,
  :minimum_change_constructible_set,
  :minimum_change_constructible_optimized,
  :minimum_change_constructible_oneliner
]

puts "Testing all approaches:"
puts "=" * 50

test_cases.each_with_index do |coins, i|
  puts "Test case #{i + 1}: #{coins.inspect}"
  
  methods.each do |method|
    result = send(method, coins.dup)  # dup to avoid modifying original
    puts "  #{method}: #{result}"
  end
  
  puts
end

# Performance comparison for larger inputs
puts "Performance comparison (larger input):"
puts "=" * 50

large_coins = (1..100).to_a + [150, 200, 300]

require 'benchmark'

Benchmark.bm(30) do |x|
  x.report("Original (greedy):") { 1000.times { minimum_change_constructible(large_coins.dup) } }
  x.report("Ruby-idiomatic:") { 1000.times { minimum_change_constructible_v2(large_coins.dup) } }
  x.report("Functional:") { 1000.times { minimum_change_constructible_v3(large_coins.dup) } }
  x.report("Optimized:") { 1000.times { minimum_change_constructible_optimized(large_coins.dup) } }
  x.report("One-liner:") { 1000.times { minimum_change_constructible_oneliner(large_coins.dup) } }
  # Skip DP and Set approaches for performance test as they're much slower
end