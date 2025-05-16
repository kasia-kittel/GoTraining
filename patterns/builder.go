package patterns

type VehicleType int

const (
	BikeType VehicleType = iota
	CarType
)

// a product
type Vehicle struct {
	color    string
	motor    string
	maxSpeed int
}

type Builder interface {
	setColor()
	setMotor()
	setMaxSpeed()
	getVehicle() Vehicle
}

type CarBuilder struct {
	color    string
	motor    string
	maxSpeed int
}

// implementation details hidden in
// concrete builders
func (c *CarBuilder) setColor() {
	c.color = "white"
}

func (c *CarBuilder) setMotor() {
	c.motor = "114 KM"
}

func (c *CarBuilder) setMaxSpeed() {
	c.maxSpeed = 180
}

func (c *CarBuilder) getVehicle() Vehicle {
	return Vehicle{
		color:    c.color,
		motor:    c.motor,
		maxSpeed: c.maxSpeed,
	}
}

type BikeBuilder struct {
	color    string
	motor    string
	maxSpeed int
}

func (b *BikeBuilder) setColor() {
	b.color = "blue"
}

func (b *BikeBuilder) setMotor() {
	b.motor = "electric"
}

func (b *BikeBuilder) setMaxSpeed() {
	b.maxSpeed = 50
}


func (b *BikeBuilder) getVehicle() Vehicle {
	return Vehicle{
		color:    b.color,
		motor:    b.motor,
		maxSpeed: b.maxSpeed,
	}
}


func getBuilder(vehicleType VehicleType) Builder {

	switch vehicleType {
	case BikeType:
		return &BikeBuilder{}
	case CarType:
		return &CarBuilder{}
	}

	return nil
}

// director
type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildVehicle() Vehicle {
	d.builder.setColor()
	d.builder.setMotor()
	d.builder.setMaxSpeed()
	return d.builder.getVehicle()
}
