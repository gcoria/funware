class TreeNode
    attr_accessor :value, :left, :right

   def initialize(value)
    @value = value
    @left = nil
    @right = nil
   end

   def insert(value)
    if value < @value
        if @left.nil?
            @left = TreeNode.new(value)
        else
            @left.insert(value)
        end
    else
        if @right.nil?
            @right = TreeNode.new(value)
        else
            @right.insert(value)
        end
    end
   end
end
