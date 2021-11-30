# MultiThreadedPacman

This is an implementation of the original PacMan game using a multithreaded approach.
Each enemy is an independent thread and the number of enemies is configurable.
to undestand better the terminal 
**P-represents the player and its move with Up/Down/Left/Right Arrows**
**G-represents the enemies(Touch one of this and you are dead**
**c-represents the enemies when theyc an be eaten +200 for each one**
**#-represents wall**
**.-represents food +10 points**
**X-represents pallets +50 points and avility to eat enemies**




## Implementation

The project was built using Golang.
a simple implementation! its run on terminal

## Functional Requirements

* The game's maze layout can be static.
* The pacman gamer must be controlled by the user.
* Enemies are autonomous entities that will move a random way.
* Enemies and pacman should respect the layout limits and walls.
* Enemies number can be configured on game's start.
* Each enemy's behaviour will be implemented as a separated thread.
* Enemies and pacman threads must use the same map or game layout data structure resource.
* Display obtained pacman's scores.
* Pacman loses when an enemy touches it.
* Pacman wins the game when it has taken all coins in the map.

## System Requirements

This project uses go modules, so Ebiten will be installed automatically as
you build the project. For Ebiten to work, you'll need to have installed:

* [Golang](https://golang.org/) version 1.15 or above
* Depending on which platform you are using, you might need to install some extra dependencies.
> The project was Tested in MacOS and Windows.

## Build/Run

### First build, then run

The project its run Like this:

```bash
go run main.go N
```

N representes de number of enemies and the spected n needs to be >0 and less than 13:

```bash
go run main.go 5
```

you can also run it with the following comand

```bash
make test
```
## Architecture

Visit the architecture document [here](./ARCHITECTURE.md).

## Video Presentation

You can check out the video.