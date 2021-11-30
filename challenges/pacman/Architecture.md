# ARCHITECTURE

## Package Index

The project is  made in a single packages, This one contains diferent stucturs and operations for each element

## Architecture
see Diagram Folder for a better Undestanding of this 

### Game Objects
the project contains mainly have three different main Structurs that are :"enemies","player",and "pallet"
which serves as an abstraction of every object in the maze.
Since the project requirements confirm that the maze can be static its not necesary to implement to this one, the dots or pallets 

Since Every member within structure is assigned a unique memory location and  Changing the value of one data member will not affect other data members in structure, this help us to make autonomous entities that will move a random way, and making sure that Each enemy's behaviour will be implemented as a separated thread.

### Collision Detection

An important step of the behavior of a movable game object is the ability to detect collisions.
player and enemies are avaible to decet collision, enemies with players and walls, and players with dots, enemies, pallets, and walls

### Goroutines

Communication between the other goroutines and the level goroutine is achieved through channels.
this chanels construct diferent threads for each Game obejct that is created, making it posible to have executing independently but concurrently sharing process resources. 


### Win or loose 
THe programs count de number of food that are in the maze, if you can finish all then the program stops and show you a little message
saying you win, and your score

## PacMan Behavior

PacMan can move in four different directions: Up, Down, Left and Right.
It is controlled by using the arrow keys.

## Enemie Behavior

We decided to adopt the original PacMan's enemie AI and made some tweaks to it.
The state diagram for every enemie can be 

Just like PacMan, a enemies state machine is represented by the state pattern.

The enemies have two types od states that are 'Running' and 'killing' the first one is when the power of the pallet is active
And they can be kill, the secong one is the normal state when they are able to kill and if the player touch it they are dead.
see Folder with diagram for a better understanding


Whenever PacMan eats a power pellet, every enemie will transition into the `Running`
state. Whenever a enemie is in this state, it will be eaten when it comes in contact
with PacMan. Then, it goes back to home and spawns in the `Killing` state.

Even if PacMan is still under the influence of the power pellet, any enemie that is
not in `running` state can kill him.

