class LogFile
  # Multiline replaced the event checks so they match multiline events
  class Multiline < LogFile
    def log(num = 1)
      num.times do |i|
        i += @count + @next
        # This pattern is chosen specifically as we can match it with all the different match types
        @file.puts 'BEGIN ' + @path + " test event #{i}"
        @file.puts " line 2 of test event #{i}"
        @file.puts " line 3 of test event #{i}"
        @file.puts ' END of test event'
      end
      @file.flush
      @count += num
      self
    end

    def skip_one
      @count -= 1
    end

    def logged?(event: event, check_file: true, check_order: true)
      return false if event['host'] != @host
      return false if check_file && event['file'] != @path
      return false if event['message'] != 'BEGIN ' + @path + " test event #{@next}" + $/ + " line 2 of test event #{@next}" + $/ + " line 3 of test event #{@next}" + $/ + ' END of test event'
      @count -= 1
      @next += 1
      true
    end
  end
end