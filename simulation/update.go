package simulation

import {

  "Go-Amusement/Ride"
  "Go-Amusement/Appointment"
  "GO-Amusement/Person"

}

func update(k key, t int){
  r := Ride.get(k)
  if r.nextRideTime == t {
    a = Appointment.get(r.appointments[r.curAppointment])
    for i := 0; i < len(a.positions); i += 1 {
      Person.get(a.peopleServed[i]).release()
    }
    r.curAppointment += 1
    count := 0
    peep := make(0,[]Person.key)
    for i:= 0; i < len(a.present); i += 1 {
      if count > 0 && i % r.rowWidth == 0 {
        peep = append(peep,r.pop(count))
      }
      if a.present[i] == false {
        count += 1
      } else {
        peep = append(peep,a.positions[i]) 
      if count > 0 {
        peep = append(peep,r.pop(count))
      }
      a.peopleServed = peep
    }
    a.save()
    for i := r.curAppointment + 1; i < len(r.Appointments); i++ {
      a = Appointment.get(r.Appointments[i])
      a.time = r.getInterval(i, t)
      a.numUntil -= 1
      a.open = !(a.numUntil == 1)
      a.save()
    }
      
    //Determine Broadcast
    r.save()
    
  }
}
    
