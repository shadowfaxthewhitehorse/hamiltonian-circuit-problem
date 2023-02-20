package main

import "fmt"

//
// PROBLEM
//
// A program in Golang to solve the Hamiltonian circuit problem using backtracking.

func hamiltonianCircuit(graph [][]int) []int {
    n := len(graph)
    path := make([]int, n)
    for i := 0; i < n; i++ {
        path[i] = -1
    }
    path[0] = 0 // Start at vertex 0
    if hamiltonianCircuitHelper(graph, path, 1) {
        return path
    } else {
        return []int{}
    }
}

func hamiltonianCircuitHelper(graph [][]int, path []int, pos int) bool {
    if pos == len(graph) {
        // Check if the last vertex in the path is adjacent to the first vertex
        if graph[path[pos-1]][path[0]] == 1 {
            return true
        } else {
            return false
        }
    }

    for v := 1; v < len(graph); v++ {
        if isValid(graph, path, pos, v) {
            path[pos] = v
            if hamiltonianCircuitHelper(graph, path, pos+1) {
                return true
            }
            path[pos] = -1
        }
    }

    return false
}

func isValid(graph [][]int, path []int, pos int, v int) bool {
    // Check if vertex v is adjacent to the last vertex in the path
    if graph[path[pos-1]][v] == 0 {
        return false
    }

    // Check if vertex v has already been visited
    for i := 0; i < pos; i++ {
        if path[i] == v {
            return false
        }
    }

    return true
}


// DEVNOTES
//
// In this program, we first define the hamiltonianCircuit function that takes a graph represented as an adjacency matrix and returns a 
// Hamiltonian circuit if one exists. The function initializes a slice path with all elements set to -1, and sets the first element to 0. 
// We then call the hamiltonianCircuitHelper function with the graph, the path slice, and the position of the second vertex (i.e., 1).
//
// The hamiltonianCircuitHelper function is a recursive function that tries to extend the path slice by considering all possible vertices that can 
// be visited next. If a valid vertex is found, the function updates the path slice and calls itself recursively with the updated path. If a Hamiltonian 
// circuit is found, the function returns true, otherwise it backtracks by resetting the vertex at the current position in the path slice and tries the 
// next vertex.
// 
// The isValid function is a helper function that checks whether a given vertex can be added to the path slice at the current position. The function checks 
// if the vertex is adjacent to the last vertex in the path, and whether the vertex has already been visited.
//
// In the code below, the program is tested with a sample graph represented as an adjacency matrix. 
//
func main() {
    // Test the program with a sample graph
    graph := [][]int{{0, 1, 1, 1, 0},
                     {1, 0, 1, 0, 1},
                     {1, 1, 0, 1, 0},
                     {1, 0, 1, 0, 1},
                     {0, 1, 0, 1, 0}}
    path := hamiltonianCircuit(graph)
    if len(path) > 0 {
        fmt.Println("Hamiltonian circuit found:", path)
    } else {
        fmt.Println("No Hamiltonian circuit found.")
    }
}
