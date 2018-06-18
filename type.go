package main 

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
    gas_pedal uint16  // min 0 max 65535
    brake_pedal uint16
    steering_wheel int16  // -32k - +32k
    top_seep_kmh float64
}


// value receiver
// point receiver
func (c car) kmh() float64{
    // c.top_seep_kmh = 500
    return float64(c.gas_pedal) * (c.top_seep_kmh/usixteenbitmax)
}

func (c car) mph() float64{
    return float64(c.gas_pedal) * (c.top_seep_kmh/usixteenbitmax/kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) {
    c.top_seep_kmh = newspeed
}

func newer_top_speed(c car, speed float64) car {
    c.top_seep_kmh = speed
    return c
}

func main() {
    a_car := car{gas_pedal: 22341,
                brake_pedal: 0,
                steering_wheel: 12561,
                top_seep_kmh: 225.0}
    // a_car := car{22341, 0, 12561, 225.0}

    fmt.Println(a_car.top_seep_kmh)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
    a_car.new_top_speed(500)
    fmt.Println(a_car.top_seep_kmh)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
    a_car = newer_top_speed(a_car, 1000)
    fmt.Println(a_car.top_seep_kmh)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
}