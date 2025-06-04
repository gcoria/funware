# Max Profit with K Transactions - Ruby Solution

def max_profit(k, prices)
  n = prices.length
  return 0 if n <= 1 || k == 0

  # If k >= n/2, we can make as many transactions as we want
  # This becomes the unlimited transactions problem
  if k >= n / 2
    return max_profit_unlimited(prices)
  end

  # DP approach for limited transactions
  # buy[i] = max profit after at most i transactions, currently holding stock
  # sell[i] = max profit after at most i transactions, not holding stock
  buy = Array.new(k + 1, -prices[0])
  sell = Array.new(k + 1, 0)

  # Process each day
  (1...n).each do |i|
    # Process transactions in reverse order to avoid using updated values
    k.downto(1) do |j|
      # sell[j] = max(keep previous sell[j], sell current stock)
      sell[j] = [sell[j], buy[j] + prices[i]].max
      # buy[j] = max(keep previous buy[j], buy stock using prev transaction)
      buy[j] = [buy[j], sell[j - 1] - prices[i]].max
    end
  end

  sell[k]
end

# Handle the case where we can make unlimited transactions
def max_profit_unlimited(prices)
  profit = 0
  (1...prices.length).each do |i|
    profit += prices[i] - prices[i - 1] if prices[i] > prices[i - 1]
  end
  profit
end

# Test with the sample input
prices = [5, 11, 3, 50, 60, 90]
k = 2

result = max_profit(k, prices)
puts "Input: prices = #{prices}, k = #{k}"
puts "Output: #{result}"
puts "Explanation: Buy at 5, sell at 11 (profit: 6); Buy at 3, sell at 90 (profit: 87); Total: #{result}"

puts "\nAdditional test cases:"

# Test case 1
prices1 = [2, 4, 1]
k1 = 2
result1 = max_profit(k1, prices1)
puts "prices = #{prices1}, k = #{k1} => #{result1}"

# Test case 2
prices2 = [3, 2, 6, 5, 0, 3]
k2 = 2
result2 = max_profit(k2, prices2)
puts "prices = #{prices2}, k = #{k2} => #{result2}"

# Test case 3 - Edge case with k larger than needed
prices3 = [1, 2, 3, 4, 5]
k3 = 10
result3 = max_profit(k3, prices3)
puts "prices = #{prices3}, k = #{k3} => #{result3}" 