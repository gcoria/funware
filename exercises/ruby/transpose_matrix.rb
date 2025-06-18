def transpose_matrix(matrix)
    transposed = Array.new(matrix[0].length) { Array.new(matrix.length) }
    
    matrix.each_with_index do |row, i|
        row.each_with_index do |col, j|
            transposed[j][i] = col
        end
    end
    transposed
end

matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

puts transpose_matrix(matrix).inspect