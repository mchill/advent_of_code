#!/usr/bin/ruby
require 'digest'

input = 'ckczppom'
number = 1
hash = Digest::MD5.hexdigest(input + number.to_s)

until hash.start_with?('00000')
    number += 1
    hash = Digest::MD5.hexdigest(input + number.to_s)
end

puts number
