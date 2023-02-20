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

# DEVNOTES
#
# In this program, we first define the hamiltonian_circuit function that takes a graph represented as an adjacency matrix and returns a 
# Hamiltonian circuit if one exists. The function initializes a list path with all elements set to -1, and sets the first element to 0. 
# We then call the hamiltonian_circuit_helper function with the graph, the path list, and the position of the second vertex (i.e., 1).
#
# The hamiltonian_circuit_helper function is a recursive function that tries to extend the path list by considering all possible vertices that can 
# be visited next. If a valid vertex is found, the function updates the path list and calls itself recursively with the updated path. If a Hamiltonian 
# circuit is found, the function returns True, otherwise it backtracks by resetting the vertex at the current position in the path list and tries the 
# next vertex.
#
# The is_valid function is a helper function that checks whether a given vertex can be added to the path list at the current position. The function 
# checks if the vertex is adjacent to the last vertex in the path, and whether the vertex has already been visited.
#
# In the example above, the program is tested with a sample graph represented as an adjacency matrix. The output is either the Hamiltonian circuit 
# found or a message indicating that no Hamiltonian circuit was found.
#
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
