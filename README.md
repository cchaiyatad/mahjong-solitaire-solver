# Mahjong Solitaire Solver

The purpose of this work is to generate a solvable mahjong solitaire board and solve it by implementing from T. Stam's paper [Solving Mahjong Solitaire Positions](http://iivq.net/scriptie/scriptie-bsc.pdf) paper. However, we only **Random** and **MaxBlock** heuristic and **Random** and **MultipleFirst** strategies.

## What can we do?

This system is a REST API system that you can

- Get a list of layouts to create a board.
- When you provide a **layout**, a **strategy**, and a **heuristic** function, the system can generate and try to solve them (sometimes it fails; \*surprised pikachu face\*) and return a board and the steps.

A new layout can be added too!, but unfortunately, it can not be done via REST API. See [Adding more layout](#adding-more-layout).

## API

### GET layout

Returns a list of layouts with 200 respond that can be use in `/solve` endpoint.

#### Endpoint

`GET /layout`

#### Example

##### Example Request

`GET /layout`

##### Example Respond

`["simple","small-pyramid","turtle"]`

### GET solve

When specific a `strategy`, `heuristic`, and `layout`, it will return a board and steps to solve a board with 200 respond.

#### Endpoint

`GET /solve`

#### Parameters

- strategy
- heuristic
- layout

All parameters are needed to be given, otherwise it will return 400 with `bad request` respond.

#### Example

##### Example Request

`GET /solve?strategy=random&heuristic=random&layout=small-pyramid`

##### Example Respond

``` JSON
{
  "board": {
    "tiles": [
      { "id": 0, "face": 37, "position": {"x": 0, "y": 0, "z": 0} },
      { "id": 2, "face": 12, "position": {"x": 2, "y": 0, "z": 0} },
      { "id": 12, "face": 37, "position": {"x": 0, "y": 2, "z": 0} },
      { "id": 14, "face": 12, "position": {"x": 2, "y": 2, "z": 0} },
      { "id": 42, "face": 25, "position": {"x": 0, "y": 1, "z": 1} },
      { "id": 44, "face": 25, "position": {"x": 2, "y": 1, "z": 1} }
    ],
    "size": { "x_size": 6, "y_size": 6, "z_size": 2 },
    "layout": "small-pyramid"
  },
  "order": [ [42, 44], [2, 14], [0, 12] ],
  "params": { "strategy": "random", "heuristic": "random", "layout": "small-pyramid" }
}
```

The respond has 3 parts. 

1. `params` will be the same with the request.
2. `board` has `layout` which will be the same as `layout` in `params` respond, `size` of board, and `tiles`. A `face` value in tiles object is represent a face of a tiles which is gurantee to be between 1-38 (inclusive).
3. `order` is a list of list of `id` of `tile`, it represent an order to pick in order to solve a board. Note that even if it can not solve the board, it will still return this field with incomplete step.


## How to run

### On local machine

- Execute `./start_app.sh` to run and now you can access at `localhost` at port `8080` unless you edited `PORT` vairable to another port.
- For testing, execute `./start_test.sh` script.

### With Docker

- Execute `start_app_docker.sh` script

## Adding more layout

If you want to add more layouts, you will need to create a .xml file that has a layout specification and place it in `assets/layout` folder.

The below xml file represent a board name **sample**, its size is  **4 \* 4 \* 2**, and has **6 tiles**.

``` xml
<layout>
    <name>sample</name>
    <size x="4" y="4" z="2"/>
    <tiles>
        <tile x="0" y="0" z="0"/>
        <tile x="2" y="0" z="0"/>
        <tile x="0" y="2" z="0"/>
        <tile x="2" y="2" z="0"/>
        <tile x="0" y="1" z="1"/>
        <tile x="2" y="1" z="1"/>
    </tiles>
</layout>
```

In the xml file, the root element, **layout**, is consists of 3 important elements, namely **name**, **size**, and **tiles**.

- The **name** element represents the name of the layout which will be appeared on `/layout` endpoint and use as a layout parameter in `/solve` endpoint.
- The **size** element represents a 3-dimension layout size. Note that if there is any tile in the **tiles** attribute that is outside the board size, it will be ignored and not created when calling `/solve` endpoint.
- The **tiles** element can have any number of **tile** elements. The **tile** element, **x, y, and z** attributes are needed to represent a position on a board. Note that the number of **tile** elements should be an even number and can not exceed 144.
