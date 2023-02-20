package main

import "fmt"

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
