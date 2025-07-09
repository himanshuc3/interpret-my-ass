# Parking lot example

## Alternate approach

1. All vehicles are instances of BaseVehicle (implementing Vehicle interface) and therefore it is going to be consumed as an interface.

- Alternatively, we could use embedding of structures which inherits the properties from the struct.

```go
type Car struct {
    BaseVehicle
}
```

- Embedding implement composition and helps with code reuse
- Any functions/interface defined on BaseVehicle, directly embeds into Car
- Property names can have conflicting names and lead to ambiguous behavior
- Property access can be directly made on the Car.propertyNameInBaseVehicle or Car.BaseVehicle.propertyName

2. Using go-routines and waitgroups for implementing parallel parking.

- WaitGroups, as their name suggests, helps in waiting for all go routines to finish.

```go
var (
    wg sync.WaitGroup
)

wg.Add(1)
go calculateImagePos()
//...
//...
wg.Wait()
```
