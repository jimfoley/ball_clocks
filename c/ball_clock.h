#include <stdio.h>
#include <stdlib.h>

#define MINUTE_TRACK 4
#define FIVE_MIN_TRACK 11
#define HOUR_TRACK 11

int nBallsInQueue,nMinsToRun;
int* aQueue;
int aMinTrack[4];
int aFiveMinTrack[11];
int aHourTrack[11];
int nMinTrackPos = 0;
int nFiveTrackPos = 0;
int nHourTrackPos = 0;
int nQueueFront = 0;
int nQueueBack = 0;
int nMinsRan = 0;
int nCycles = 0;

int i;

void InitBallClock(int argc,char *argv[]);
void RunClock();
void EmptyMinuteReset();
void EmptyMinute();
void EmptyFiveMinuteReset();
void EmptyFiveMinute();
void EmptyHourReset();
void EmptyHour();
void PrintQueue();
void CleanUp();





