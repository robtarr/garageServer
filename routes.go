package main

func (s *server) routes() {
	s.mux.HandleFunc("/doorStatus", s.DoorStatus)
	s.mux.HandleFunc("/closeDoor", s.CloseDoor)
	s.mux.HandleFunc("/openDoor", s.OpenDoor)
	s.mux.HandleFunc("/openClose", s.OpenClose)
}
