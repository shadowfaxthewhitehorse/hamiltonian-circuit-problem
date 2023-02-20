from typing import List


# PROBLEM:
#
# A Python program that uses backtracking to solve the Hamiltonian circuit problem
 
def hamiltonian_circuit(graph: List[List[int]]) -> List[int]:
    n = len(graph)
    path = [-1] * n
    path[0] = 0  # Start at vertex 0
    if hamiltonian_circuit_helper(graph, path, 1):
        return path
    else:
        return []

def hamiltonian_circuit_helper(graph: List[List[int]], path: List[int], pos: int) -> bool:
    if pos == len(graph):
        # Check if the last vertex in the path is adjacent to the first vertex
        if graph[path[-1]][path[0]] == 1:
            return True
        else:
            return False

    for v in range(1, len(graph)):
        if is_valid(graph, path, pos, v):
            path[pos] = v
            if hamiltonian_circuit_helper(graph, path, pos + 1):
                return True
            path[pos] = -1

    return False

def is_valid(graph: List[List[int]], path: List[int], pos: int, v: int) -> bool:
    # Check if vertex v is adjacent to the last vertex in the path
    if graph[path[pos - 1]][v] == 0:
        return False

    # Check if vertex v has already been visited
    if v in path[:pos]:
        return False

    return True

# Test the program with a sample graph
graph = [[0, 1, 1, 1, 0],
         [1, 0, 1, 0, 1],
         [1, 1, 0, 1, 0],
         [1, 0, 1, 0, 1],
         [0, 1, 0, 1, 0]]
path = hamiltonian_circuit(graph)
if path:
    print("Hamiltonian circuit found:", path)
else:
    print("No Hamiltonian circuit found.")
