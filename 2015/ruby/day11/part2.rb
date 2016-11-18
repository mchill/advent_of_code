#!/usr/bin/ruby

password = 'hxbxxyzz'
valid = false

until valid
    password.next!

    if password.include?('iol') or password.scan(/(\w)\1/).count < 2
        next
    end

    for i in 2..password.length - 1
        if (password[i-2].next == password[i-1] and
            password[i-1].next == password[i])
            valid = true
        end
    end
end

puts password
