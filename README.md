# Ant Colony Simulation

A Go-based ant colony simulation that demonstrates emergent behavior and swarm intelligence through various scenarios. The simulation showcases how simple rules followed by individual ants can lead to complex, colony-level behaviors. This is a fun project to practice Golang and experiment with Raylib bindings.

## Results

- Multiple simulation scenarios:
  - `home`: Basic home navigation scenario
  - `wandering`: Ants wandering behavior demonstration
  - `food`: Food foraging simulation
  - `pheromone`: Pheromone-based path finding
  - `one_nest`: Single nest colony simulation
  - `two_nests`: Two competing colonies
  - `four_nests`: Four colony interaction
  - `surrounding`: Environmental interaction scenario

## Examples

 <video src="docs/ant_sim_four_nests.mp4" width="300" controls></video><video src="docs/ant_sim_surrounding_nest.mp4" width="300" controls></video> <video src="docs/ant_sim_pheromone_only.mp4" width="300" controls></video>

## Prerequisites

- Go 1.20 or higher
- [raylib-go](https://pkg.go.dev/github.com/gen2brain/raylib-go/raylib) - Golang bindings for raylib, see how to properly install it depending on your os [here](https://github.com/gen2brain/raylib-go)

## Try it locally

1. Clone the repository:

```bash
git clone https://github.com/firasjaber/ant-sim.git
cd ant-sim
```

2. Install dependencies:

```bash
go mod download
```

3. Run a simulation scenario:

```bash
go run main.go [scenario_name]
```

For example:

```bash
go run main.go four_nests
```

## Simulation Scenarios

- `home`: Demonstrates ants returning to their nest
- `wandering`: Shows random exploration behavior
- `food`: Simulates food foraging with pheromone trails
- `pheromone`: Focuses on pheromone-based communication
- `one_nest`: Single colony behavior
- `two_nests`: Competition between two colonies
- `four_nests`: Complex interactions between four colonies
- `surrounding`: Environmental factors affecting ant behavior

## Built With

- [Go](https://golang.org/)
- [raylib-go](https://github.com/gen2brain/raylib-go)
