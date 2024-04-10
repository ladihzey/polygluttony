def call_it(one, two)
  puts one unless one.nil?
  puts two unless two.nil?
end

call_it('one', 'two', 'three')
call_it('one', 'two')
call_it('one')
call_it()
