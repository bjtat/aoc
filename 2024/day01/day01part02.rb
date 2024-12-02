lhs = []
rhs = []
File.open('day01.txt', 'r') do |f|
  f.each_line do |line|
    split_line = line.split
    lhs << split_line[0].to_i
    rhs << split_line[1].to_i
  end
  f.close
end

lhs.sort!
rhs_ht = rhs.each_with_object(Hash.new(0)) { |num, counts| counts[num] += 1 }

similarity = 0
lhs.each do |left|
  similarity += left * rhs_ht[left]
end

puts similarity
