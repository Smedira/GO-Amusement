package Ride
 
import(
  //"GO-Amusement/Appointment"
  "GO-Amusement/Person"
)
 
type ride struct{

  appointments []key
  queue []key //of people, FIFO order
  GuarunteedQueueFlow int
  Broadcast Range
  timeDistribution key
  nextRideTime int
  lastBroadcast int //time
  exit
  entrance
  numTrains int
  curAppointment int

}

func (r *ride) pop(int count) []Person.key{

  p = make([]Person.person,0)

  for i := 0; i < len(r.partySize); i++{
    if r.partySize[i] <= count{
      p = append(p,r.queue(i: r.partySize[i] + i)
      count -= r.partySize[i]
      r.queue = append(r.queue[:i], r.queue[r.partySize[i] + i:]...)
      if count == 0 { break }

  return p

