C
Run the ball clock by entering "./ball_clock <clock size> <minutes>".  The program takes two parameters, the number of balls in the queue and the number of minutes to run.

GoLang
This implementation of the ball clock is state driven. A state driven design means the application will return to the main routine after each tick of the clock.  Although slower than a more conventional design, the state driven design provides a way to provide animation so you could draw the clock and update as it runs.

Ruby
Run the ball clock by entering 'ruby -r ./ball_clock.rb -e "BallClock.new.run_clock" <clock size> <minutes>'.

The Ball Clock Puzzle
A ball clock is a 12-hour mechanical device that uses a track, queue, rotator, marbles, and three chutes to keep track of time. Here is an explanation of these parts:
1. The queue holds all the marbles that are not currently in use.
2. The rotator picks up one marble every minute and deposits it on the track.
3. The track delivers each marble to one of the three chutes.
4. The chutes keep track of time. One chute tracks minutes, one tracks five minute increments,
and the last one tracks hours.
5. The 1 Minute chute holds up to 4 marbles and the 5 Minute and 1 Hour chutes both hold up
to 11 marbles.To read the clock, you simply calculate the number of marbles in each chute according to the chute’s value. For example, if there is one marble in the 1 Minute Chute, 3 marbles in the 5 Minute chute, and 6 marbles in the 1 Hour chute, the time would be 6:16.
How the Ball Clock Works
1. A timed rotator spins, picking up one marble every minute and depositing it on the track.
2a. The marble rolls down the track and stops at the 1 Minute chute. If this chute is not full, the marble is deposited into the 1 Minute chute and the clock waits for the next marble.
2b. If the 1 Minute chute is full, the marbles in the chute are dumped out, in reverse order, and returned to the queue at the bottom of the clock. Then the current marble continues to the 5
Minute chute.
3a. If the 5 Minute chute is not full, the current marble is deposited into the 5 Minute chute and the clock waits for the next marble.
3b. If the 5 Minute chute is full, the marbles in the chute are dumped out, in reverse order, and returned to the queue at the bottom of the clock. Then, the current marble continues to the 1 Hour chute.
4a. If the 1 Hour chute is not full, the current marble is deposited into the 1 Hour chute and the clock waits for the next marble.
4b. If the 1 Hour chute is full, the marbles in the chute are dumped out, in reverse order, and
returned to the queue at the bottom of the clock. The current marble is then also returned to the queue. This situation, where all the chutes are empty, represents 12:00.
About The Ball Clock Queue
Assume for a moment that the ball clock starts out with 30 balls in the queue. These marbles
are all numbered, 1 through 30, and are placed in the queue in numerical order. As the marbles
rotate through the clock, their order will be messed up. For example, during the first minute, ball 1 will be placed in the 1 Minute chute. Then, during the next three minutes, balls 2, 3, and 4 will be deposited into the 1 Minute chute. When ball 5 comes down the track, it will trigger the 1 Minute chute to empty, since it is full. When it empties, the balls will be returned to the queue in reverse order, so now the order of the queue is 6 through 30, in numerical order, then 4, 3, 2, 1.
