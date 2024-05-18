class Tree(object):
  def __init__(self, x):
    self.value = x
    self.left = None
    self.right = None

def find_max(root):
    while root.right:
        root = root.right
    return root.value

def remove(root, val):
    if not root:
        return None
    
    if val < root.value:
        root.left = remove(root.left, val)
    elif val > root.value:
        root.right = remove(root.right, val)
    else:
        if root.left:
            root.value = find_max(root.left)
            root.left = remove(root.left, root.value)
        else:
            root = root.right
    return root

def solution(t, queries):
    for querie in queries:
        t = remove(t, querie)
    return t

def print_tree(root):
    if not root:
        return
    print_tree(root.left)
    print(root.value, ", ", end="")
    print_tree(root.right)

tree = Tree(7)
tree.left = Tree(2)
tree.left.left = Tree(1)
tree.left.right = Tree(6)
tree.left.right.left = Tree(4)
tree.left.right.left.left = Tree(3)
tree.left.right.left.right = Tree(5)
tree.right = Tree(9)
tree.right.left = Tree(8)
tree.right.right = Tree(10)

print_tree(tree)
print()
tree = solution(tree, [2])
print_tree(tree)
# 1, 2, 3, 7, 8

# tree = Tree(3)
# tree.left = Tree(2)
# tree = solution(tree, [1,2,3,5])
# print_tree(tree)