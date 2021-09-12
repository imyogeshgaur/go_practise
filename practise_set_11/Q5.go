package main;

type TVRemote interface {
	off();
	chanelChange();
}

type SmartTVRemote interface{
	watchYoutube();
	connectYoutube();
}