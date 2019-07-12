function defaultItemA() {
  return [
    {
      title: 'Navigate with arrow keys',
      labelA: [1],
      childA: [
        {
          title: 'Up/down to move within layers',
          labelA: [1],
          childA: [],
        },
        {
          title: 'Left/right to move between layers',
          labelA: [1],
          childA: [],
        },
        {
          title: 'Enter, then left/right to edit',
          labelA: [1],
          childA: [],
        },
        {
          title: 'Escape to cancel',
          labelA: [1],
          childA: [],
        },
      ],
    },
    {
      title: 'Filter by labels',
      labelA: [2],
      childA: [
        {
          title: 'Shift-1 to add label 1',
          labelA: [2],
          childA: [],
        },
        {
          title: '1 to only show label 1',
          labelA: [2],
          childA: [],
        },
        {
          title: '1 again to show all items',
          labelA: [2],
          childA: [],
        },
      ],
    },
    {
      title: 'Group by labels',
      labelA: [2],
      childA: [
        {
          title: 'G to group by labels',
          labelA: [2],
          childA: [],
        },
        {
          title: 'G again to show structure',
          labelA: [2],
          childA: [],
        },
      ],
    },
    {
      title: 'Plan multiple frenzies',
      labelA: [3],
      childA: [],
    },
    {
      title: 'Get started for free',
      labelA: [3],
      childLabelA: [3],
      childA: [],
    },
  ];
}

function rebalanceItemA(pathItemA, pathDispA, itemA, viewLabelA) {
  // rebalances pathItemA, pathDispA, labelChildA
  // childLabelA has all children's labelA as well as it's own labelA
  let i;
  let r;
  const out = { dispA: [], itemA: [], labelA: [] };
  for (i = 0; i < itemA.length; i += 1) {
    const item = JSON.parse(JSON.stringify(itemA[i]));
    const disp = JSON.parse(JSON.stringify(itemA[i]));
    item.pathItemA = pathItemA.concat([i]);
    disp.pathDispA = pathDispA.concat([out.dispA.length]);
    disp.pathItemA = item.pathItemA;

    r = rebalanceItemA(item.pathItemA, disp.pathDispA, item.childA, viewLabelA);
    item.childA = r.itemA;
    item.childLabelA = item.labelA.concat(r.labelA);
    out.itemA.push(item);

    disp.childA = r.dispA;
    disp.childLabelA = disp.labelA.concat(r.labelA);
    if (
      viewLabelA.length === 0
      || viewLabelA.filter(l => disp.childLabelA.indexOf(l) !== -1).length > 0
    ) {
      out.dispA.push(disp);
    }

    out.labelA = out.labelA.concat(item.childLabelA);
  }
  return out;
}

function viewLabelHelper(viewLabelA0, label) {
  let viewLabelA = viewLabelA0;
  if (viewLabelA.indexOf(label) === -1) {
    viewLabelA.push(label);
  } else {
    viewLabelA = viewLabelA.filter(l => l !== label);
  }
  return viewLabelA;
}

function labelItemHelper(label, itemA0, pathItemA0, step) {
  // recursively goes to item and toggles label
  const itemA = itemA0;
  if (step === pathItemA0.length - 1) {
    if (itemA[pathItemA0[step]].labelA.indexOf(label) === -1) {
      itemA[pathItemA0[step]].labelA.push(label);
    } else {
      itemA[pathItemA0[step]].labelA = itemA[pathItemA0[step]].labelA.filter(l => l !== label);
    }
  } else {
    itemA[pathItemA0[step]].childA = labelItemHelper(
      label, itemA[pathItemA0[step]].childA, pathItemA0, step + 1,
    );
  }
  return itemA;
}

function moveItemHelper(dir, itemA0, pathItemA0, step) {
  let temp;
  const itemA = itemA0;
  const row = pathItemA0[step];
  if (step === pathItemA0.length - 1) {
    if (dir === 'up' && row > 0) {
      temp = itemA[row];
      itemA.splice(row, 1);
      itemA.splice(row - 1, 0, temp);
      itemA[row].pathItemA[itemA[row].pathItemA.length - 1] = row;
      itemA[row - 1].pathItemA[itemA[row - 1].pathItemA.length - 1] = row - 1;
    } else if (dir === 'down' && row <= itemA.length - 2) {
      temp = itemA[row];
      itemA.splice(row, 1);
      itemA.splice(row + 1, 0, temp);
      itemA[row].pathItemA[itemA[row].pathItemA.length - 1] = row;
      itemA[row + 1].pathItemA[itemA[row + 1].pathItemA.length - 1] = row + 1;
    }
  } else {
    itemA[row].childA = moveItemHelper(dir, itemA[row].childA, pathItemA0, step + 1);
  }

  return itemA;
}

function deleteItemHelper(itemA0, pathItemA0, step) {
  let i;
  const itemA = itemA0;
  if (step === pathItemA0.length - 1) {
    itemA.splice(pathItemA0[step], 1);
    for (i = 0; i < itemA.length; i += 1) {
      itemA[i].pathItemA[step] = i;
    }
  } else {
    itemA[pathItemA0[step]].childA = deleteItemHelper(
      itemA[pathItemA0[step]].childA, pathItemA0, step + 1,
    );
  }
  return itemA;
}

function addItemHelper(title, itemA0, viewLabelA, parentPathItemA, step) {
  // recursively go through itemA and add new item
  // labelA is automatically set to viewLabelA
  const itemA = itemA0;
  if (step === parentPathItemA.length) {
    // parentPathItemA[parentPathItemA.length - 1] doesn't exist yet, so must push
    itemA.push({
      title,
      labelA: viewLabelA,
      childA: [],
    });
  } else {
    itemA[parentPathItemA[step]].childA = addItemHelper(
      title, itemA[parentPathItemA[step]].childA, viewLabelA, parentPathItemA, step + 1,
    );
  }

  return itemA;
}

function updateItemHelper(title, itemA0, pathItemA, step) {
  // recursively go through itemA and update title to new title
  const itemA = itemA0;
  if (step === pathItemA.length - 1) {
    itemA[pathItemA[step]].title = title;
  } else {
    itemA[pathItemA[step]].childA = updateItemHelper(
      title, itemA[pathItemA[step]].childA, pathItemA, step + 1,
    );
  }
  return itemA;
}

module.exports = {
  defaultItemA() {
    return defaultItemA();
  },
  rebalanceItemA(pathItemA, pathDispA, itemA0, viewLabelA) {
    return rebalanceItemA(pathItemA, pathDispA, itemA0, viewLabelA);
  },
  viewLabelHelper(viewLabelA0, label) {
    return viewLabelHelper(viewLabelA0, label);
  },
  labelItemHelper(label, itemA0, pathItemA0, step) {
    return labelItemHelper(label, itemA0, pathItemA0, step);
  },
  moveItemHelper(dir, itemA0, pathItemA0, step) {
    return moveItemHelper(dir, itemA0, pathItemA0, step);
  },
  deleteItemHelper(itemA0, pathItemA0, step) {
    return deleteItemHelper(itemA0, pathItemA0, step);
  },
  addItemHelper(title, itemA0, viewLabelA, pathItemA, step) {
    return addItemHelper(title, itemA0, viewLabelA, pathItemA, step);
  },
  updateItemHelper(title, itemA0, pathItemA, step) {
    return updateItemHelper(title, itemA0, pathItemA, step);
  },
};
