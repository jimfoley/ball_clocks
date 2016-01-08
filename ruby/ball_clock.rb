#!/usr/bin/ruby

class BallClock

  def initialize
    if (ARGV.length != 1) && (ARGV.length != 2)
      puts 'Usage: ruby -r ./ball_clock.rb -e "BallClock.new.run_clock" <clock size> <minutes>'
    end
    @balls_in_queue = ARGV[0].to_i
    @mins_to_run = ARGV[1].to_i
    @queue = (0..@balls_in_queue-1).to_a
    @min_track = (0..3).to_a
    @five_min_track = (0..10).to_a
    @hour_track = (0..10).to_a
    @min_track_pos = 0
    @five_track_pos = 0
    @hour_track_pos = 0
    @queue_front = 0
    @queue_back = 0
    @mins_ran = 0
    @cycles = 0
  end

  def run_clock
    loop do
      @mins_ran+=1
      if @queue_front==@balls_in_queue
        @queue_front=0
      end
      if @min_track_pos<4
        @min_track[@min_track_pos] = @queue[@queue_front]
        @queue_front+=1
        @min_track_pos+=1
      else
        if @queue_back+4>=@balls_in_queue
          empty_minute_reset
        else
          empty_minute
        end
        @min_track_pos=0
        if @five_track_pos<11
          @five_min_track[@five_track_pos]=@queue[@queue_front]
          @queue_front+=1
          @five_track_pos+=1
        else
          if @queue_back+11>=@balls_in_queue
            empty_five_minute_reset
          else
            empty_five_minute
          end
          @five_track_pos=0
          if @hour_track_pos<11
            @hour_track[@hour_track_pos]=@queue[@queue_front]
            @queue_front+=1
            @hour_track_pos+=1
          else
            if @queue_back+11>=@balls_in_queue
              empty_hour_reset
            else
              empty_hour
            end
            @hour_track_pos=0
            @queue[@queue_back]=@queue[@queue_front]
            @queue_front+=1
            @queue_back+=1
            if @queue_back==@balls_in_queue
              @queue_back=0
            end
            @cycles+=1
            ARGV.length == 1 ? check_one : check_two
          end
        end
      end
    end
  end

  def check_one
    0.upto @balls_in_queue-1 do |i|
      if @queue[i]!=i
        break
      end
      puts "#{@queue}"
      exit
    end
  end

  def check_two
    print '{"Min":['
      (0..@min_track.length-2).step(1).each do |a|
        print @min_track[a].to_s + ','
      end
      print @min_track[@min_track.length-1]
      print '],"FiveMin":['
      (0..@five_min_track.length-2).step(1).each do |a|
        print @five_min_track[a].to_s + ','
      end
      print @five_min_track[@five_min_track.length-1]
      print '],"Hour":['
      (0..@hour_track.length-2).step(1).each do |a|
      print @hour_track[a].to_s + ','
    end
    print @hour_track[@hour_track.length-1]
    print '],"Queue":['
    (0..@queue.length-2).step(1).each do |a|
      print @queue[a].to_s + ','
    end
    print @queue[@queue.length-1]
    print '"}'
    exit
  end

  def empty_minute_reset
    (@min_track.length-1).step(0,-1).each { |i|
      @queue[@queue_back]=@min_track[i]
      @queue_back+=1
      if @queue_back==@balls_in_queue
        @queue_back=0
      end
    }
  end

  def empty_minute
    (@min_track.length-1).step(0,-1) { |i| @queue[@queue_back]=@min_track[i]; @queue_back+=1 }
  end

  def empty_five_minute_reset
    (@five_min_track.length-1).step(0,-1) { |i|
      @queue[@queue_back]=@five_min_track[i]
      @queue_back+=1
      if @queue_back==@balls_in_queue
        @queue_back=0
      end
    }
  end

  def empty_five_minute
    (@five_min_track.length-1).step(0,-1) { |i| @queue[@queue_back]=@five_min_track[i]; @queue_back+=1 }
  end

  def empty_hour_reset
    (@hour_track.length-1).step(0,-1) { |i|
      @queue[@queue_back]=@hour_track[i]
      @queue_back+=1
      if @queue_back==@balls_in_queue
        @queue_back=0
      end
    }
  end

  def empty_hour
    (@hour_track.length-1).step(0,-1) { |i| @queue[@queue_back]=@hour_track[i]; @queue_back+=1 }
  end

end