# This is a shorthand for StandardError => e
def execute
  raise StandardError.new("This is a standard error")
end

def regular_raise
  begin
    execute
  rescue => e
    puts 'Regular raise:'
    puts e.backtrace
    raise
  end
end

def re_raise
  begin
    execute
  rescue => e
    puts 'Re-raise:'
    puts e.backtrace
    raise e
  end
end

begin
  regular_raise
rescue => e
  puts 'Regular raise outside:'
  puts e.backtrace
  puts
end

begin
  re_raise
rescue => e
  puts 'Re-raise outside:'
  puts e.backtrace
  puts
end

