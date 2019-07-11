# Plan Your Next Frenzy

## Modes

- normal: navigate structure of items
- move: move items


## Persistence

- database is updated when itemA changes
- otherwise all on the client side


## Functions

- rebalance: calculates
- move: when entering move mode, all items are automatically shown (viewLabelA = [])


## Variables

- itemA: structure of all items
- dispA: display structure of all items

- pathItemA: position of item in itemA, used for reorganization
- pathDispA: position of item in dispA, used for navigation
- posA: position of current navigation, matches pathDispA, not pathA
- (when viewLabelA !== [], pathDispA !== pathA)

- viewLabelA: list of labels to be shown, empty means show all items, [1, 2, 3] means show all items with 1, 2, 3
- labelA: list of labels of item
- labelChildA: list of labels of item and it's descendents, calculated in rebalance
