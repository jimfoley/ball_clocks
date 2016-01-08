#include "ball_clock.h"

int main(int argc,char *argv[])
{
	InitBallClock(argc,argv);
	RunClock();
	CleanUp();
}

void InitBallClock(int argc,char *argv[])
{
	if (argc != 3) {
		printf("Usage: %s <clock size> <minutes>\n",argv[0]);
		exit (0);
	}
	nBallsInQueue = atoi(argv[1]);
	nMinsToRun = atoi(argv[2]);
	
	aQueue=(int*)malloc(nBallsInQueue*sizeof(int));
	for (i=0;i<nBallsInQueue;i++) {
		aQueue[i]=i;
	}
}

void RunClock() {
	while (1==1) {
		nMinsRan++;
		if (nQueueFront == nBallsInQueue)
			nQueueFront = 0;
		if (nMinTrackPos < 4) {
			aMinTrack[nMinTrackPos] = aQueue[nQueueFront];
			nQueueFront++;
			nMinTrackPos++;
		}
		else {
			nQueueBack+4 >= nBallsInQueue ? EmptyMinuteReset() : EmptyMinute();
			nMinTrackPos = 0;
			if (nFiveTrackPos < 11) {
				aFiveMinTrack[nFiveTrackPos]=aQueue[nQueueFront];
				nQueueFront++;
				nFiveTrackPos++;
			}
			else {
				nQueueBack+11 >= nBallsInQueue ? EmptyFiveMinuteReset() : EmptyFiveMinute();
				nFiveTrackPos=0;
				if (nHourTrackPos < 11) {
					aHourTrack[nHourTrackPos] = aQueue[nQueueFront];
					nQueueFront++;
					nHourTrackPos++;
				}
				else {
					nQueueBack+11 >= nBallsInQueue ? EmptyHourReset() : EmptyHour();
					nHourTrackPos = 0;
					aQueue[nQueueBack] = aQueue[nQueueFront];
					nQueueFront++;
					nQueueBack++;
					if (nQueueBack == nBallsInQueue)
						nQueueBack=0;
					nCycles++;
					PrintQueue();
					break;
				}
			}
		}
	}
}

void EmptyMinuteReset() {
	for (i=MINUTE_TRACK-1; i>0; i--) {
		aQueue[nQueueBack] = aMinTrack[i];
		nQueueBack++;
		if (nQueueBack == nBallsInQueue)
			nQueueBack = 0;
	}
}

void EmptyMinute() {
	for (i=MINUTE_TRACK-1; i>0; i--) {
		aQueue[nQueueBack] = aMinTrack[i];
		nQueueBack++;
	}
}

void EmptyFiveMinuteReset() {
	for (i=FIVE_MIN_TRACK-1; i>0; i--) {
		aQueue[nQueueBack] = aFiveMinTrack[i];
		nQueueBack++;
		if (nQueueBack == nBallsInQueue)
			nQueueBack = 0;
	}
}

void EmptyFiveMinute() {
	for (i=FIVE_MIN_TRACK-1; i>0; i--) {
		aQueue[nQueueBack] = aFiveMinTrack[i];
		nQueueBack++;
	}
}

void EmptyHourReset() {
	for (i=HOUR_TRACK-1; i>0; i--) {
		aQueue[nQueueBack] = aHourTrack[i];
		nQueueBack++;
		if (nQueueBack == nBallsInQueue)
			nQueueBack = 0;
	}
}

void EmptyHour() {
	for (i=HOUR_TRACK-1; i>0; i--) {
		aQueue[nQueueBack] = aHourTrack[i];
		nQueueBack++;
	}
}

void PrintQueue() {
	for (i=0; i<nBallsInQueue; i++)
		printf("%d ",aQueue[i]);
	printf("\n");
}

void CleanUp()
{
	free (aQueue);
}