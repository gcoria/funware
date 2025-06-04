def river_sizes(matrix)
    sizes = []
    visited = {}
    
    matrix.each_with_index do |row, i|
        row.each_with_index do |val, j|
            if val == 1 && !visited[cords(i,j)]
                size = explore(matrix, i, j, visited)
                sizes << size
            end
        end
    end
    return sizes
end

def cords(i, j)
    "#{i},#{j}"
end

def explore(matrix, i, j, visited)
    size = 0
    stack = [[i,j]]

    while !stack.empty?
        current = stack.pop
        curr_i, curr_j = current[0], current[1]
        key = cords(curr_i, curr_j)

        if curr_i < 0 || curr_j < 0 || curr_i >= matrix.length || curr_j >= matrix[0].length || matrix[curr_i][curr_j] == 0 || visited[key]
            next
        end

        visited[key] = true
        size += 1

        stack << [curr_i+1, curr_j]
        stack << [curr_i-1, curr_j]
        stack << [curr_i, curr_j+1]
        stack << [curr_i, curr_j-1]
    end    

    return size
end

def print_matrix(matrix)
    puts  "  " + "-" * matrix.length * 2    
    matrix.each do |row|
        puts "| " + row.join(" ") + " |"
    end
    puts "  " + "-" * matrix.length * 2 
end

matrix = [
  [1, 0, 1, 0, 1],
  [1, 0, 0, 1, 1],
  [1, 0, 1, 1, 1],
  [1, 1, 0, 0, 1]
]

print_matrix(matrix)
puts "River sizes: #{river_sizes(matrix).inspect}"